package repositories

import (
	"fmt"

	"github.com/Julio-Cesar07/codepix/domain/model"

	"github.com/jinzhu/gorm"
)

type GormTransactionRepository struct {
	Db *gorm.DB
}

func (repository *GormTransactionRepository) Register(transaction *model.Transaction) error {
	err := repository.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *GormTransactionRepository) Save(transaction *model.Transaction) error {
	err := repository.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *GormTransactionRepository) Find(transaction_id string) (*model.Transaction, error) {
	var transaction model.Transaction

	repository.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", transaction_id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("Transaction not found.")
	}

	return &transaction, nil
}

