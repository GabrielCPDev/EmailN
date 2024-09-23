package endpoints

import (
	internalmock "emailn/internal/test/internalMock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsStart_should_return_200(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATH", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)

	assert.Nil(err)
	assert.Equal(200, status)
}


func Test_CampaignsStart_should_return_err(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something is whrong")
	service.On("Start", mock.Anything).Return(errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATH", "/", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignStart(rr, req)

	assert.Equal(errExpected, err)
}
