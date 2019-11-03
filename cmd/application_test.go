package cmd

import (
	"strconv"
	"testing"

	mock_ovh "github.com/emmanuelCarre/ovh-cli/mocks"
	"github.com/emmanuelCarre/ovh-cli/ovh"
	"github.com/emmanuelCarre/ovh-cli/ovh/model"

	"github.com/golang/mock/gomock"
)

func Test_fetchApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_ovh.NewMockClient(ctrl)
	var application model.Application
	appllicationID := 10
	m.EXPECT().Get(applicationURI+strconv.Itoa(appllicationID), application).Return(nil)
	tmp := *m
	tmp1 := ovh.Client(tmp)
	fetchApplication(&tmp1, appllicationID, &application)
}
