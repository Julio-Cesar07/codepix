package repositories

import "github.com/Julio-Cesar07/codepix/domain/model"

type TransactionRepository interface {
	Register(transaction *model.Transaction) error
	Save(transaction *model.Transaction) error
	Find(transaction_id string) (*model.Transaction, error)
}