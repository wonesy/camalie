package chat

import (
	"github.com/stretchr/testify/suite"
	"github.com/wonesy/camalie/chat/mocks"
)

type ClientTestSuite struct {
	suite.Suite

	ws *mocks.ClientWebSocket
}

func (s *ClientTestSuite) SetupTest() {
	s.ws = new(mocks.ClientWebSocket)
}

func (s *ClientTestSuite) TestNewClient() {
	c := NewClient(s.ws)

	s.NotNil(c)
}
