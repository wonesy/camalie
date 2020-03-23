package yenta

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	mocks "github.com/wonesy/camalie/yenta/mocks"
)

type HubTestSuite struct {
	suite.Suite
}

func TestHubSuite(t *testing.T) {
	suite.Run(t, new(HubTestSuite))
}

func (s *HubTestSuite) TestNewHub() {
	hub := NewHub()
	s.NotNil(hub)
}

func (s *HubTestSuite) TestRegisterSpoke_OK() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	spk := mocks.NewMockSpoke(ctrl)

	hub := NewHub()
	hub.Start()

	hub.Register(spk)

	// give time so that the registration is taken into account
	<-time.After(1 * time.Second)

	s.Equal(1, len(hub.spokes))
}

func (s *HubTestSuite) TestRegisterSpoke_AlreadyRegistered() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	uuid, _ := uuid.NewUUID()

	spk := mocks.NewMockSpoke(ctrl)

	spk.EXPECT().ID().Return(uuid).Times(2)

	hub := NewHub()
	hub.Start()

	hub.Register(spk)
	hub.Register(spk) // disregard, already registered

	// give time so that the registration is taken into account
	<-time.After(1 * time.Second)

	s.Equal(1, len(hub.spokes))
}

func (s *HubTestSuite) TestUnregisterSpoke_OK() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	id, _ := uuid.NewUUID()
	id2, _ := uuid.NewUUID()

	spk := mocks.NewMockSpoke(ctrl)
	spk2 := mocks.NewMockSpoke(ctrl)

	spk.EXPECT().ID().Return(id).Times(1)
	spk2.EXPECT().ID().Return(id2).Times(1)

	hub := NewHub()
	hub.spokes = append(hub.spokes, spk)
	hub.spokes = append(hub.spokes, spk2)
	s.Equal(2, len(hub.spokes))

	hub.Start()

	hub.Unregister(id2)

	// give time so that the registration is taken into account
	<-time.After(1 * time.Second)

	s.Equal(1, len(hub.spokes))
}

func (s *HubTestSuite) TestStop() {
	hub := NewHub()
	hub.Start()
	err := hub.Stop()
	s.NoError(err)
}

func (s *HubTestSuite) TestBroadcast() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	spk := mocks.NewMockSpoke(ctrl)
	spk2 := mocks.NewMockSpoke(ctrl)

	spk.EXPECT().Send("test").Return(nil).Times(1)
	spk2.EXPECT().Send("test").Return(nil).Times(1)

	hub := NewHub()
	hub.spokes = []Spoke{spk, spk2}
	hub.Start()

	hub.Broadcast("test")

	<-time.After(1 * time.Second)
}
