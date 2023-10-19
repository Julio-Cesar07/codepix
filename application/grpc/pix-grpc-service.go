package grpc

import (
	"context"

	"github.com/Julio-Cesar07/codepix/application/grpc/pb"
	usecases "github.com/Julio-Cesar07/codepix/application/use-cases"
)

type PixGrpcService struct {
	PixUseCase usecases.PixKeyUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {

	key, err := p.PixUseCase.Register(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error: err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id: key.ID,
		Status: "created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id: pixKey.ID,
		Kind: pixKey.Kind,
		Key: pixKey.Key,
		Account: &pb.Account{
			AccountId: pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId: pixKey.Account.BankID,
			BankName: pixKey.Account.Bank.Name,
			OwnerName: pixKey.Account.OwnerName,
			CreatedAt: pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixGrpcService(usecase usecases.PixKeyUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}