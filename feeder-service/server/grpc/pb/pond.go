package pb

import (
	"context"
	"feeder-service/server/grpc/pb/pond"
	"feeder-service/service"
)

type PondGRPC struct {
	pond.UnimplementedPondServiceServer
	Services service.PondService
}

func (g *PondGRPC) Pond(ctx context.Context, req *pond.PondRequest) (*pond.PondResponse, error) {
	ID := req.GetId()
	if ID == "" {
		return &pond.PondResponse{Status: false, Message: "id is empty", Data: nil}, nil
	}
	result, err := g.Services.FetchByUUID(ctx, ID)
	if err != nil {
		return &pond.PondResponse{Status: false, Message: err.Error(), Data: nil}, nil
	}

	return &pond.PondResponse{Status: true, Message: "success", Data: &pond.Pond{
		Uuid: result.UUID,
		Name: result.Name,
		// CreatedAt: result.CreatedAt,
		// UpdatedAt: result.UpdatedAt,
	}}, nil
}
