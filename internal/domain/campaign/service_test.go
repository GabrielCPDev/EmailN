package campaign

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaign) error {
	args := r.Called(campaing)
	return args.Error(0)
}

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@test.com"},
	}
	service = Service{}
)

func Test_Create_Save_Campaing(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.Equal("error to save on database", err.Error())
}
