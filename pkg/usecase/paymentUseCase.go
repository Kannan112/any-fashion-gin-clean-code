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

func (c *PaymentUsecase) RazorPayCheckout(ctx context.Context, userId int) (res.RazorPayResponse, error) {
	var RazorPay res.RazorPayResponse
	cart, err := c.cartrepo.FindCart(ctx, userId)
	if err != nil {
		return RazorPay, errors.New("failed to find cart")
	}
	if cart.Sub_total == 0 {
		return RazorPay, errors.New("there is no products in your list")
	}

	//	check the addresses move to user repo as FIND addres
	checkbool, err := c.userrepo.FindAddress(ctx, userId)
	if err != nil {
		return RazorPay, errors.New("error found at find address checkout usecase")
	}
	if !checkbool {
		return RazorPay, errors.New("add addresses")
	}
	razorpayKey := config.GetConfig().RazorKey
	razorpaySecret := config.GetConfig().RazorSec

	client := razorpay.NewClient(razorpayKey, razorpaySecret)

	razorPayAmount := cart.Sub_total * 100

	data := map[string]interface{}{
		"amount":   razorPayAmount,
		"currency": "INR",
		"receipt":  "reciept_id",
	}
	// create an order on razor pay
	order, err := client.Order.Create(data, nil)

	if err != nil {
		return RazorPay, fmt.Errorf("faild to create razorpay order")
	}

	return res.RazorPayResponse{
		Email:       "abhinandarun11@gmail.com",
		PhoneNumber: "9846789099",
		RazorpayKey: razorpayKey,
		PaymentId:   uint(1),
		OrderId:     order["id"],
		Total:       uint(razorPayAmount),
		AmountToPay: uint(cart.Sub_total),
	}, nil
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
	return stripeOrder, nil

}

func (c *PaymentUsecase) VerifyStripeOrder(ctx context.Context, stripePaymentID string) error {
	stripe.Key = stripePaymentID
	paymentIntent, err := paymentintent.Get(stripePaymentID, nil)

	if err != nil {
		return errors.New("failed to get payment intent from stripe")
	}
	if paymentIntent.Status == "succeeded" {
		return nil
	} else {
		return errors.New("payment not approved")
	}
}
