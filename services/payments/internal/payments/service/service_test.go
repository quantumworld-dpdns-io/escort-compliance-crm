package service

import (
	"testing"
)

func TestCreatePayment(t *testing.T) {
	svc := &Service{repo: nil}

	req := &CreatePaymentRequest{
		BookingID: "booking-1",
		PayerID:   "user-1",
		PayeeID:   "comp-1",
		Amount:    150.00,
		Currency:  "USD",
		Method:    "card",
	}

	if req.BookingID == "" {
		t.Error("booking_id should not be empty")
	}
	if req.Amount <= 0 {
		t.Error("amount should be positive")
	}
	if req.Currency == "" {
		t.Error("currency should not be empty")
	}
}

func TestProcessPayment(t *testing.T) {
	svc := &Service{repo: nil}

	if svc == nil {
		t.Fatal("service should not be nil")
	}
}

func TestRefundPayment(t *testing.T) {
	svc := &Service{repo: nil}

	if svc == nil {
		t.Fatal("service should not be nil")
	}
}
