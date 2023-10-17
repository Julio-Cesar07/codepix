package repositories

import (
	"github/Julio-Cesar07/codepix/domain/model"
)	

type PixKeyRepository interface {
	RegisterKey(pixKey *model.PixKey) (*model.PixKey, error)
	FindKeyByKind(key string, kind string) (*model.PixKey, error)
	AddBank(bank *model.Bank) error
	AddAccount(account *model.Account) error
	FindAccount(account_id string) (*model.Account, error)
}