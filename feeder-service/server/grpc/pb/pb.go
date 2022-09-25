package pb

import (
	"context"
	"feeder-service/server/grpc/pb/chat"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *chat.MessageTitle) (*chat.Message, error) {
	keyword := in.GetTitle()

	return &chat.Message{Title: keyword, Content: "Hello From the Server!"}, nil
}
