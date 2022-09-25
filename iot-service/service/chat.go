package service

import (
	"context"
	"iot-service/pkg"
	service "iot-service/proto"
)

type (
	ChatService interface {
		Fatch() (*string, error)
	}

	chatServiceImpl struct {
		GRPC pkg.GRPC
	}
)

func NewChatService(GRPC pkg.GRPC) ChatService {
	return &chatServiceImpl{GRPC: GRPC}
}

func (i *chatServiceImpl) Fatch() (*string, error) {
	chat := service.NewChatServiceClient(i.GRPC.ClientConn)

	req := &service.MessageTitle{Title: "Service IoT"}
	response, err := chat.SayHello(context.Background(), req)
	if err != nil {

		return nil, err
	}
	return &response.Title, nil
}
