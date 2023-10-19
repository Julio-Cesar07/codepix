package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/Julio-Cesar07/codepix/application/grpc/pb"
	usecases "github.com/Julio-Cesar07/codepix/application/use-cases"
	"github.com/Julio-Cesar07/codepix/infra/repositories"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// registrando serviço
	pixRepository := repositories.GormPixKeyRepositoryDb{Db: database}
	pixUseCase := usecases.PixKeyUseCase{
		PixKeyRepository: pixRepository,
	}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)


	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server.", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server.", err)
	}	
}