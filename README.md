# stripe

To use this application run  **go run main.go**

and in another tab in vs code  do 
stripe login
stripe listen --forward-to localhost:8080/webhook
it will monitor the webhook

then open browser and hit this url 
http://localhost:8080/
now you will see a form like this
![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/9f42ee5a-e15d-471e-87f5-7a419a484960)


now fill the details and click submit you will redirect on this url http://localhost:8080/donate

and this page will open
![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/fbc968b1-fd9e-4e42-915e-7accc04b7208)

now click on checkout session will create and you will send to stripe payment processing page
fill the demo details like this
![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/d9151796-ed9f-4d9c-aeca-a72002c14ed4)
click on pay
since i am using a demo account this page will open 
![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/1d0ef3fa-c3ea-45b0-8732-c2afc64e14bc) click on complete if transaction success

this page will open  ![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/af90ef91-ebd8-478f-94fd-c0155c397b4b)


****now in my code short description what will happening ****
i am taking details from user and on checkout a session is created and user is redirected to stripe and i am monitoring payment intent and on payment_intent.succeeded  i am saving details in db with status success
on failure i means when payment_intent.payment_failed

printing failed on ui we can also save this in db




screenshot of terminal

![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/de762597-5519-48a9-92cc-46125f07fb10)


![image](https://github.com/shyamalkaushiks/stripe/assets/47667670/2c5bf25b-d6da-4056-8ecc-d03f0f090f04)

