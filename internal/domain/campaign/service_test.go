package campaign

import (
	"emailn/internal/contract"
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
	service := Service{}

	newCampaing := contract.NewCampaing{
		Nome:    "Test X",
		Content: "Body",
		Emails:  []string{"test1", "test2"}}

	id, err := service.Create(newCampaing)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_Create_Save_Campaing(t *testing.T) {

	newCampaing := contract.NewCampaing{
		Nome:    "Test X",
		Content: "Body",
		Emails:  []string{"test1", "test2"}}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.Name != newCampaing.Nome ||
			campaign.Content != newCampaing.Content ||
			len(campaign.Contacts) != len(newCampaing.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)
}
