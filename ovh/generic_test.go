package ovh

import (
	"errors"
	"testing"

	"github.com/emmanuelCarre/ovh-cli/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFetchResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockClient(ctrl)
	var ressource Resource
	client := Client(m)
	m.EXPECT().Get("/hello/2", ressource).Times(1)
	err := FetchResource(&client, "/hello/", 2, ressource)
	assert.Nil(t, err)
}

func TestFetchResourceWithIsse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockClient(ctrl)
	var ressource Resource
	client := Client(m)
	err := errors.New("Error message")
	m.EXPECT().Get("/hello/2", ressource).Times(1).Return(err)
	result := FetchResource(&client, "/hello/", 2, ressource)
	assert.Equal(t, err, result)
}

func TestFetchResourcesID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockClient(ctrl)
	var ressource Resource
	client := Client(m)
	m.EXPECT().Get("/hello/2", ressource).Times(1)
	FetchResourcesID(&client, "/hello/")
	err := FetchResource(&client, "/hello/", 2, ressource)
	assert.Nil(t, err)
}
