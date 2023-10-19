package usecases

import (
	"github.com/Julio-Cesar07/codepix/domain/model"
	"github.com/Julio-Cesar07/codepix/domain/repositories"
)

type TransactionUseCase struct {
	TransacationRepository repositories.TransactionRepository
	PixKeyRepository repositories.PixKeyRepository
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := t.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKeyToStruct, err := t.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKeyToStruct, description)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransacationRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Confirm()
	err = t.TransacationRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransacationRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Complete()
	err = t.TransacationRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransacationRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Cancel(reason)
	err = t.TransacationRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}