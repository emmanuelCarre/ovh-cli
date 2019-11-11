package ovh

import (
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
	m.EXPECT().Get("/hello/2", ressource)
	err := FetchResource(&client, "/hello/", 2, ressource)
	assert.Nil(t, err)
}
