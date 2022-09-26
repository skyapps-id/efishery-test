package pb

import (
	"context"
	"feeder-service/server/grpc/pb/feeder"
	"feeder-service/service"
)

type FeederGRPC struct {
	feeder.UnimplementedFeederServiceServer
	Services service.FeederService
}

func (g *FeederGRPC) FatchFeeders(ctx context.Context, req *feeder.FeederRequestAll) (*feeder.FeederResponses, error) {
	var (
		results = []*feeder.Feeder{}
	)

	feeders, err := g.Services.Fatch(ctx)
	if err != nil {
		return &feeder.FeederResponses{Status: false, Message: err.Error(), Data: results}, nil
	}

	for _, row := range *feeders {
		results = append(results, &feeder.Feeder{
			Uuid:    row.UUID,
			Barcode: row.Barcode,
			Name:    row.Name,
		})
	}

	return &feeder.FeederResponses{Status: true, Message: "success", Data: results}, nil
}

func (g *FeederGRPC) FindFeederById(ctx context.Context, req *feeder.FeederRequestById) (*feeder.FeederResponse, error) {
	ID := req.GetId()
	if ID == "" {
		return &feeder.FeederResponse{Status: false, Message: "field id is empty", Data: nil}, nil
	}
	result, err := g.Services.FatchByUUID(ctx, ID)
	if err != nil {
		return &feeder.FeederResponse{Status: false, Message: err.Error(), Data: nil}, nil
	}

	return &feeder.FeederResponse{Status: true, Message: "success", Data: &feeder.Feeder{
		Uuid:    result.UUID,
		Barcode: result.Barcode,
		Name:    result.Name,
		// CreatedAt: result.CreatedAt,
		// UpdatedAt: result.UpdatedAt,
	}}, nil
}

func (g *FeederGRPC) FindFeederByBarcode(ctx context.Context, req *feeder.FeederRequestByBarcode) (*feeder.FeederResponses, error) {
	var (
		barcode = req.GetBarcode()
		results = []*feeder.Feeder{}
	)
	if len(barcode) == 0 {
		return &feeder.FeederResponses{Status: false, Message: "field barcode is empty", Data: results}, nil
	}
	feeders, err := g.Services.FatchByBarcode(ctx, barcode)
	if err != nil {
		return &feeder.FeederResponses{Status: false, Message: err.Error(), Data: results}, nil
	}

	for _, row := range *feeders {
		results = append(results, &feeder.Feeder{
			Uuid:    row.UUID,
			Barcode: row.Barcode,
			Name:    row.Name,
		})
	}

	return &feeder.FeederResponses{Status: true, Message: "success", Data: results}, nil
}
