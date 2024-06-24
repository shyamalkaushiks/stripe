package services

import (
	"encoding/json"
	"fmt"
	"hello/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

var payamount float64

func (hs *HandlerService) loadform(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

func (hs *HandlerService) acceptdata(c *gin.Context) {
	var payment models.Paymentdatails
	err := c.ShouldBind(&payment)
	if err != nil {
		panic(err)
	}
	fmt.Println(payment.Amount, payment.Email, payment.Name)
	c.HTML(http.StatusOK, "userdetails.html", gin.H{
		"Payment": payment,
	})
	// hs.PaymentProcessing(c, payment)
}

func (hs *HandlerService) createCheckoutSession(c *gin.Context) {
	stripe.Key = "your api key"
	var details models.Paymentdatails

	if err := c.ShouldBind(&details); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payamount = details.Amount

	fmt.Println(payamount)
	fmt.Println("details on chcokt", details.Name+"d", details.Amount)
	domain := "http://localhost:8080"
	amountInPaise := int64(details.Amount)
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("inr"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Donation"),
					},
					UnitAmount: stripe.Int64(int64(amountInPaise * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/successpage"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}

func (hs *HandlerService) handleStripeWebhook(c *gin.Context) {
	event := stripe.Event{}

	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	//req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		// w.WriteHeader(http.StatusServiceUnavailable)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})

		return
	}

	switch event.Type {
	case "payment_intent.succeeded":
		fmt.Println("Payment intent succeeded")
		// session := sessions.Default(c)
		// paymentAmount := session.Get("payment_amount")
		fmt.Println("Payment amount:", payamount)
		details := models.TransactionDetails{
			CoustumerName: "shyamal",
			Amount:        payamount,
			Status:        "success",
		}
		if err := DB.Create(&details).Error; err != nil {
			fmt.Fprintf(os.Stderr, "Failed to save transaction details: %v\n", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		} else {
			fmt.Println("saved")

		}

		c.HTML(http.StatusOK, "success.html", nil)
	case "payment_intent.payment_failed":
		fmt.Println("Payment intent failed")

	default:
		fmt.Printf("Unhandled event type: %s\n", event.Type)
	}

}
