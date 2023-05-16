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
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
	cartRepo  interfaces.CartRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, cartRepo interfaces.CartRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}
func (c *OrderUseCase) OrderAll(id int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(id)
	return order, err
}
func (c *OrderUseCase) UserCancelOrder(orderId, userId int) error {
	err := c.orderRepo.UserCancelOrder(orderId, userId)
	return err
}
func (c *OrderUseCase) ListAllOrders(userId int) ([]domain.Order, error) {
	order, err := c.orderRepo.ListAllOrders(userId)
	return order, err
}

func (c *OrderUseCase) RazorPayCheckout(ctx context.Context, userId int, paymentId int) (res.RazorPayResponse, error) {

	cart, err := c.cartRepo.FindCart(ctx, userId)

	if cart.Sub_total == 0 {
		return res.RazorPayResponse{}, fmt.Errorf("there is no products in your list")
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
		return res.RazorPayResponse{}, fmt.Errorf("faild to create razorpay order")
	}

	return res.RazorPayResponse{
		Email:       "",
		PhoneNumber: "",
		RazorpayKey: razorpayKey,
		PaymentId:   uint(paymentId),
		OrderId:     order["id"],
		Total:       uint(razorPayAmount),
		AmountToPay: uint(cart.Sub_total),
	}, nil
}
func (c *OrderUseCase) VerifyRazorPay(ctx context.Context, body req.RazorPayRequest) error {
	razorpayKey := config.GetConfig().RazorKey
	razorPaySecret := config.GetConfig().RazorSec

	//varify signature
	data := body.RazorPayOrderId + "|" + body.RazorPayPaymentId
	h := hmac.New(sha256.New, []byte(razorPaySecret))
	_, err := h.Write([]byte(data))
	if err != nil {
		return errors.New("faild to veify signature")
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(body.Razorpay_signature)) != 1 {
		return errors.New("razorpay signature not match")
	}

	// then vefiy payment
	client := razorpay.NewClient(razorpayKey, razorPaySecret)
	// fetch payment and vefify
	payment, err := client.Payment.Fetch(body.RazorPayPaymentId, nil, nil)

	if err != nil {
		return err
	}

	// check payment status
	if payment["status"] != "captured" {
		return errors.New("faild to verify razorpay payment")
	}

	return nil
}
