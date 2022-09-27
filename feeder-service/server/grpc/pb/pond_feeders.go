package pb

import (
	"context"
	"feeder-service/server/grpc/pb/pond_feeders"
	"feeder-service/service"
)

type PondFeedersGRPC struct {
	pond_feeders.UnimplementedPondFeedersServiceServer
	Services service.PondFeederService
}

func (g *PondFeedersGRPC) FetchPondFeeders(ctx context.Context, req *pond_feeders.PondFeedersRequest) (*pond_feeders.PondFeedersResponse, error) {
	var (
		pondUUID = req.GetPondUuid()
		results  = []*pond_feeders.PondFeeder{}
	)

	if pondUUID == "" {
		return &pond_feeders.PondFeedersResponse{Status: false, Message: "field pond_uuid is empty", Data: nil}, nil
	}
	feeders, err := g.Services.FetchPondFeeders(ctx, pondUUID)
	if err != nil {
		return &pond_feeders.PondFeedersResponse{Status: false, Message: err.Error(), Data: results}, nil
	}

	for _, row := range *feeders {
		results = append(results, &pond_feeders.PondFeeder{
			PondUuid:   row.PondUUID,
			FeederUuid: row.FeederUUID,
			Name:       row.Name,
			Barcode:    row.Barcode,
		})
	}

	return &pond_feeders.PondFeedersResponse{Status: true, Message: "success", Data: results}, nil
}
