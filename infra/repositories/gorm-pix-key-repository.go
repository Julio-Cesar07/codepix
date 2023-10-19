package repositories

import (
	"fmt"

	"github.com/Julio-Cesar07/codepix/domain/model"

	"github.com/jinzhu/gorm"
)

type GormPixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (repository GormPixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := repository.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository GormPixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := repository.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository GormPixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error){
	err := repository.Db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (repository GormPixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	repository.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("No key was found.")
	}

	return &pixKey, nil
}

func (repository GormPixKeyRepositoryDb) FindAccount(account_id string) (*model.Account, error) {
	var account model.Account

	repository.Db.Preload("Bank").First(&account, "id = ?", account_id)

	if account.ID == "" {
		return nil, fmt.Errorf("No account was found")
	}

	return &account, nil
}

func (repository GormPixKeyRepositoryDb) FindBank(bank_id string) (*model.Bank, error) {
	var bank model.Bank

	repository.Db.First(&bank, "id = ?", bank_id)

	if bank.ID == "" {
		return nil, fmt.Errorf("No bank found")
	}

	return &bank, nil
}