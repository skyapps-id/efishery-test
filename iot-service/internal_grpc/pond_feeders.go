package internal_grpc

import (
	"context"
	"errors"
	"iot-service/dto"
	"iot-service/pkg"
	service "iot-service/proto/pond_feeders"
	"os"
)

type (
	PondFeedersInternalGRPC interface {
		FetchPondFeeders(ctx context.Context, pondUuid string) (*dto.PondFeederGRPCResponse, error)
	}

	pondFeedersInternalGRPCImpl struct {
	}
)

func NewPondFeedersInternalGRPC() PondFeedersInternalGRPC {
	return &pondFeedersInternalGRPCImpl{}
}

func (i *pondFeedersInternalGRPCImpl) FetchPondFeeders(ctx context.Context, pondUuid string) (*dto.PondFeederGRPCResponse, error) {
	svrGRPC := os.Getenv("GRPC_FEEDER_SERVICE")
	if svrGRPC == "" {
		svrGRPC = "localhost:9001"
	}
	gRPC := pkg.NewGrpcDial(svrGRPC)

	pondFeeders := service.NewPondFeedersServiceClient(gRPC.ClientConn)

	req := &service.PondFeedersRequest{PondUuid: pondUuid}
	response, err := pondFeeders.FetchPondFeeders(ctx, req)
	if err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, errors.New(response.Message)
	}

	if len(response.Data) == 0 {
		return nil, errors.New("not found record")
	}

	var results = []dto.PondFeeder{}
	for _, row := range response.Data {
		results = append(results, dto.PondFeeder{
			PondUUID:   row.PondUuid,
			FeederUUID: row.FeederUuid,
			Name:       row.Name,
			Barcode:    row.Barcode,
		})
	}

	return &dto.PondFeederGRPCResponse{
		Status:  response.Status,
		Message: response.Message,
		Data:    results,
	}, nil
}
