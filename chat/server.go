package chat

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/wonesy/camalie/db"
	pbchat "github.com/wonesy/camalie/proto/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server manages all endpoints for a chat server
type Server struct {
	conn *db.Connection
}

// NewServer constructor for server
func NewServer(conn *db.Connection) *Server {
	return &Server{
		conn: conn,
	}
}

// CreateHub creates a new chat hub and immediately allows the client to join
func (s *Server) CreateHub(ctx context.Context, r *pbchat.CreateHubRequest) (*empty.Empty, error) {
	// create a new hub in the database and save it in the
	return &empty.Empty{}, nil
}

// JoinHub allows a client to join an existing hub
func (s *Server) JoinHub(ctx context.Context, r *pbchat.JoinHubRequest) (*empty.Empty, error) {
	query := "SELECT * FROM hub WHERE id=$1"

	rows, err := s.conn.Query(query, r.GetHub())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan()
	}

	return &empty.Empty{}, nil
}

// LeaveHub allows a client to leave a hub they've joined
func (s *Server) LeaveHub(ctx context.Context, r *pbchat.LeaveHubRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
