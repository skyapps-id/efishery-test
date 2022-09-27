package internal_grpc

import (
	"context"
	"errors"
	"iot-service/dto"
	"iot-service/pkg"
	service "iot-service/proto/feeder"
	"os"
)

type (
	FeederInternalGRPC interface {
		FetchFeeder(ctx context.Context) (*dto.FeederGRPCResponse, error)
	}

	feederInternalGRPCImpl struct {
	}
)

func NewFeederInternalGRPC() FeederInternalGRPC {
	return &feederInternalGRPCImpl{}
}

func (i *feederInternalGRPCImpl) FetchFeeder(ctx context.Context) (*dto.FeederGRPCResponse, error) {
	svrGRPC := os.Getenv("GRPC_FEEDER_SERVICE")
	if svrGRPC == "" {
		svrGRPC = "localhost:9001"
	}
	gRPC := pkg.NewGrpcDial(svrGRPC)
	fedders := service.NewFeederServiceClient(gRPC.ClientConn)

	req := &service.FeederRequestAll{}
	response, err := fedders.FetchFeeders(ctx, req)
	if err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, errors.New(response.Message)
	}

	if len(response.Data) == 0 {
		return nil, errors.New("not found record")
	}

	var results = []dto.FeederGRPC{}
	for _, row := range response.Data {
		results = append(results, dto.FeederGRPC{
			UUID:    row.Uuid,
			Name:    row.Name,
			Barcode: row.Barcode,
		})
	}

	return &dto.FeederGRPCResponse{
		Status:  response.Status,
		Message: response.Message,
		Data:    results,
	}, nil
}
