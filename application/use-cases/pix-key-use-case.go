package usecases

import (
	"errors"

	"github.com/Julio-Cesar07/codepix/domain/model"
	"github.com/Julio-Cesar07/codepix/domain/repositories"
)

type PixKeyUseCase struct {
	PixKeyRepository repositories.PixKeyRepository
}

func (useCase *PixKeyUseCase) Register(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := useCase.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	useCase.PixKeyRepository.RegisterKey(pixKey)

	if pixKey.ID == ""{
		return nil, errors.New("unable to create new key at the moment.")
	}

	return pixKey, nil
}

func (useCase *PixKeyUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := useCase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}