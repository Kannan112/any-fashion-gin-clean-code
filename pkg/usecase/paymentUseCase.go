package usecase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/razorpay/razorpay-go"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type PaymentUsecase struct {
	paymentrepo interfaces.PaymentRepo
	cartrepo    interfaces.CartRepository
	userrepo    interfaces.UserRepository
	config      config.Config
}

func NewPaymentUsecase(paymentrepo interfaces.PaymentRepo, cartRepo interfaces.CartRepository, userRepo interfaces.UserRepository, config config.Config) services.PaymentUsecases {
	return &PaymentUsecase{
		paymentrepo: paymentrepo,
		cartrepo:    cartRepo,
		userrepo:    userRepo,
		config:      config,
	}
}

func (c *PaymentUsecase) ListPaymentMethod(ctx context.Context) ([]domain.PaymentMethod, error) {
	result, err := c.paymentrepo.ListPaymentMethod(ctx)
	if err != nil {
		return nil, fmt.Errorf("found an error checking: %v", err)
	}
	return result, nil
}

func (c *PaymentUsecase) UpdatePaymentMethod(id int, Payment req.PaymentReq) error {
	err := c.paymentrepo.UpdatePaymentMethod(id, Payment)
	return err
}

func (c *PaymentUsecase) RazorPayCheckout(ctx context.Context, userId int) (res.RazorpayOrder, error) {
	var razorPay res.RazorpayOrder
	cart, err := c.cartrepo.FindCart(ctx, userId)
	if err != nil {
		return razorPay, errors.New("failed to find cart")
	}
	if cart.Sub_total == 0 {
		return razorPay, errors.New("there are no products in your list")
	}

	// Check the addresses, move to user repo as FindAddress
	checkBool, err := c.userrepo.FindAddress(ctx, userId)
	if err != nil {
		return razorPay, errors.New("error found at finding address in checkout usecase")
	}
	if !checkBool {
		return razorPay, errors.New("add addresses")
	}

	razorpayKey := config.GetConfig().RazorKey
	razorpaySecret := config.GetConfig().RazorSec

	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	razorPayAmount := cart.Sub_total * 100

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  "receipt_id",
	}
	// Create an order on Razorpay
	order, err := client.Order.Create(data, nil)
	if err != nil {
		return razorPay, fmt.Errorf("failed to create Razorpay order: %v", err)
	}

	razorpayOrderID := order["id"]

	razorPayOrder := res.RazorpayOrder{
		ShopOrderID:     cart.Id,
		AmountToPay:     uint(cart.Sub_total),
		RazorpayAmount:  uint(razorPayAmount),
		RazorpayKey:     razorpayKey,
		RazorpayOrderID: razorpayOrderID,
		UserID:          uint(userId),
		Email:           "example@gmail.com",
		Phone:           "99999999",
	}
	return razorPayOrder, nil

}

func (c *PaymentUsecase) VerifyRazorPay(ctx context.Context, body req.RazorPayRequest) error {
	razorpayKey := config.GetConfig().RazorKey
	razorPaySecret := config.GetConfig().RazorSec
	fmt.Println("test")

	//varify signature
	data := body.RazorPayOrderId + "|" + body.RazorPayPaymentId
	h := hmac.New(sha256.New, []byte(razorPaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		fmt.Println("test1")
		return errors.New("faild to veify signature")
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(body.Razorpay_signature)) != 1 {
		return errors.New("razorpay signature not match")
	}

	fmt.Println("test2")

	// then vefiy payment
	client := razorpay.NewClient(razorpayKey, razorPaySecret)
	// fetch payment and vefify
	payment, err := client.Payment.Fetch(body.RazorPayPaymentId, nil, nil)

	if err != nil {
		return err
	}

	// check payment status
	if payment["status"] != "captured" {
		fmt.Println("test3")
		return errors.New("faild to verify razorpay payment")
	}

	return nil
}

func (c *PaymentUsecase) MakeStripeOrder(ctx context.Context, userID uint) (res.StripeOrder, error) {
	cart, err := c.cartrepo.FindCart(ctx, int(userID))
	if err != nil {
		return res.StripeOrder{}, err
	}
	userDetails, err := c.userrepo.GetUserDetailsFromUserID(userID)
	if err != nil {
		return res.StripeOrder{}, err
	}

	stripe.Key = config.GetConfig().StripeKey

	params := &stripe.PaymentIntentParams{

		Amount:       stripe.Int64(int64(cart.Total)),
		ReceiptEmail: stripe.String(userDetails.Email),
		Currency:     stripe.String(string(stripe.CurrencyINR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		return res.StripeOrder{}, err
	}
	clientSecret := paymentIntent.ClientSecret
	stripePublishKey := c.config.PublishableKey

	stripeOrder := res.StripeOrder{
		CartID:         cart.Id,
		AmountToPay:    uint(cart.Total),
		ClientSecret:   clientSecret,
		PublishableKey: stripePublishKey,
	}

	fmt.Println("->>>>", stripeOrder)
	return stripeOrder, nil

}

func (c *PaymentUsecase) VerifyStripeOrder(ctx context.Context, stripePaymentID string) error {
	stripe.Key = c.config.StripeKey

	// get payment by payment_id
	paymentIntent, err := paymentintent.Get(stripePaymentID, nil)

	if err != nil {
		return fmt.Errorf("failed to get payment intent from stripe")
	}

	// verify the payment intent
	if paymentIntent.Status != stripe.PaymentIntentStatusSucceeded && paymentIntent.Status != stripe.PaymentIntentStatusRequiresCapture {
		return ErrPaymentNotApproved
	}
	return nil
}
