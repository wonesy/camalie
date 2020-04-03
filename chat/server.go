package chat

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SendMessage(ctx context.Context, msg *pbchat.Message) (*empty.Empty, error) {
	fmt.Println(msg.GetSender())
	return &empty.Empty{}, nil
}
