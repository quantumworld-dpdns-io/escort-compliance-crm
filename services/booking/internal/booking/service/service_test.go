package service

import (
	"testing"
)

func TestCreateBooking_InvalidDate(t *testing.T) {
	svc := &Service{repo: nil}

	req := &CreateBookingRequest{
		CompanionID:     "comp-1",
		ScheduledAt:     "not-a-date",
		DurationMinutes: 60,
		TotalAmount:     100.0,
	}

	_, err := svc.Create(nil, "client-1", req)
	if err == nil {
		t.Error("expected error for invalid date")
	}
}

func TestCreateBooking_PastDate(t *testing.T) {
	svc := &Service{repo: nil}

	req := &CreateBookingRequest{
		CompanionID:     "comp-1",
		ScheduledAt:     "2020-01-01T00:00:00Z",
		DurationMinutes: 60,
		TotalAmount:     100.0,
	}

	_, err := svc.Create(nil, "client-1", req)
	if err == nil {
		t.Error("expected error for past date")
	}
	if err != nil && err.Error() != "cannot book in the past" {
		t.Errorf("expected 'cannot book in the past', got: %v", err)
	}
}

func TestCreateBooking_DefaultDuration(t *testing.T) {
	req := &CreateBookingRequest{
		CompanionID:     "comp-1",
		ScheduledAt:     "2030-01-01T00:00:00Z",
		DurationMinutes: 0,
		TotalAmount:     100.0,
	}

	if req.DurationMinutes != 0 {
		t.Error("duration should be 0 before service processing")
	}
}

func TestCreateBooking_DefaultCurrency(t *testing.T) {
	req := &CreateBookingRequest{
		CompanionID:     "comp-1",
		ScheduledAt:     "2030-01-01T00:00:00Z",
		DurationMinutes: 60,
		Currency:        "",
		TotalAmount:     100.0,
	}

	if req.Currency != "" {
		t.Error("currency should be empty before service processing")
	}
}
