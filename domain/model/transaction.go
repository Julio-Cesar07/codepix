package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base `valid:"required"`
	AccountFrom *Account `valid:"-"`
	Amount float64	`json:"amount" valid:"notnull"`
	PikKeyTo *PixKey `valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"cancel_description" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0{
		return errors.New("The amount must be greater than 0.")
	}

	if transaction.Status != TransactionPending && 
	transaction.Status != TransactionCompleted && 
	transaction.Status != TransactionConfirmed && 
	transaction.Status != TransactionError {
		return errors.New("Invalid status for the transaction.")
	}

	if transaction.PikKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("The source and destination account cannot be the same.")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount: amount,
		PikKeyTo: pixKeyTo,
		Description: description,
		Status: TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transaction *Transaction) touch() {
	transaction.UpdatedAt = time.Now()
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.touch()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.touch()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Cancel(cancel_description string) error {
	transaction.Status = TransactionError
	transaction.CancelDescription = cancel_description
	transaction.touch()

	err := transaction.isValid()

	return err
}