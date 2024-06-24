package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerService struct{}

func (hs *HandlerService) Bootstrap(r *gin.Engine) {
	r.GET("/", hs.loadform)
	r.POST("/donate", hs.acceptdata)
	// r.GET("/buynow", hs.PaymentProcessing)
	r.POST("/create-checkout-session", hs.createCheckoutSession)
	r.POST("/webhook", hs.handleStripeWebhook)

	r.GET("/successpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "success.html", nil)
	})

}
