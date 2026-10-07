package service

import (
	"context"
	"fmt"
	"time"

	"projek-backend/app/model"
	"projek-backend/app/repository"
)

type PaymentService struct {
	paymentRepository *repository.PaymentRepository
	rentalRepository  *repository.RentalRepository
}

func NewPaymentService(
	paymentRepository *repository.PaymentRepository,
	rentalRepository *repository.RentalRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
		rentalRepository:  rentalRepository,
	}
}

func (s *PaymentService) Create(
	ctx context.Context,
	userID int,
	role string,
	request *model.CreatePaymentRequest,
) (*model.Payment, error) {

	if userID <= 0 {
		return nil, fmt.Errorf("user tidak valid")
	}

	if role != "customer" && role != "staff" && role != "admin" {
		return nil, fmt.Errorf("role tidak valid")
	}

	if request.RentalID <= 0 {
		return nil, fmt.Errorf("rental_id tidak valid")
	}

	if request.Amount <= 0 {
		return nil, fmt.Errorf("amount harus lebih dari 0")
	}

	if request.PaymentMethod != "cash" &&
		request.PaymentMethod != "transfer" &&
		request.PaymentMethod != "qris" {
		return nil, fmt.Errorf(
			"payment_method harus cash, transfer, atau qris",
		)
	}

	rental, err := s.rentalRepository.FindByID(
		ctx,
		request.RentalID,
	)
	if err != nil {
		return nil, err
	}

	if role == "customer" && rental.UserID != userID {
		return nil, fmt.Errorf("akses rental ditolak")
	}

	if request.Amount != rental.TotalPrice {
		return nil, fmt.Errorf(
			"jumlah pembayaran harus sama dengan total harga rental",
		)
	}

	_, err = s.paymentRepository.FindByRentalID(
		ctx,
		request.RentalID,
	)

	if err == nil {
		return nil, fmt.Errorf(
			"rental sudah memiliki pembayaran",
		)
	}

	now := time.Now()

	payment := &model.Payment{
		RentalID:      request.RentalID,
		Amount:        request.Amount,
		PaymentMethod: request.PaymentMethod,
		PaymentStatus: "paid",
		PaidAt:        &now,
	}

	if err := s.paymentRepository.Create(
		ctx,
		payment,
	); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) GetByID(
	ctx context.Context,
	id int,
	userID int,
	role string,
) (*model.Payment, error) {

	if id <= 0 {
		return nil, fmt.Errorf("id payment tidak valid")
	}

	if userID <= 0 {
		return nil, fmt.Errorf("user tidak valid")
	}

	if role != "customer" && role != "staff" && role != "admin" {
		return nil, fmt.Errorf("role tidak valid")
	}

	payment, err := s.paymentRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	rental, err := s.rentalRepository.FindByID(
		ctx,
		payment.RentalID,
	)
	if err != nil {
		return nil, err
	}

	if role == "customer" && rental.UserID != userID {
		return nil, fmt.Errorf("akses payment ditolak")
	}

	return payment, nil
}
