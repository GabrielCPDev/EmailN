package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/internalErrors"
	internalmock "emailn/internal/test/internalMock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Test Y",
		Content:   "Body Hi!",
		Emails:    []string{"teste1@test.com"},
		CreatedBy: "teste@teste.com.br",
	}
	campaignPendenting *campaign.Campaign
	campaignStarted    *campaign.Campaign
	repositoryMock     *internalmock.CampaignRepositoryMock
	service            = campaign.ServiceImp{}
)

func SetUp() {
	repositoryMock = new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
	campaignPendenting, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	campaignStarted = &campaign.Campaign{ID: "1", Status: campaign.Started}
}

func Test_Create_Campaign(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalErrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	SetUp()
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalErrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaing(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaignPendenting.ID
	})).Return(campaignPendenting, nil)

	campaignReturned, _ := service.GetBy(campaignPendenting.ID)

	assert.Equal(campaignPendenting.ID, campaignReturned.ID)
	assert.Equal(campaignPendenting.Name, campaignReturned.Name)
	assert.Equal(campaignPendenting.Content, campaignReturned.Content)
	assert.Equal(campaignPendenting.Status, campaignReturned.Status)
	assert.Equal(campaignPendenting.CreatedBy, campaignReturned.CreatedBy)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something wrong'"))

	_, err := service.GetBy(campaignPendenting.ID)

	assert.Equal(internalErrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete("campaignIdInvalid")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignStarted, nil)

	err := service.Delete(campaignStarted.ID)

	assert.Equal("Campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignPendenting, nil)
	repositoryMock.On("Delete", mock.Anything).Return(errors.New("error to delete campaign"))

	err := service.Delete(campaignPendenting.ID)

	assert.Equal(internalErrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_when_delete_has_success(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignPendenting, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignPendenting == campaign
	})).Return(nil)

	err := service.Delete(campaignPendenting.ID)

	assert.Nil(err)
}

func Test_Start_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start("campaignIdInvalid")

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())

}

func Test_Start_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignStarted, nil)

	err := service.Start(campaignStarted.ID)

	assert.Equal("Campaign status invalid", err.Error())
}

func Test_Start_ReturnError_when_func_SendMail_fail(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignPendenting, nil)
	sendMail := func(campaign *campaign.Campaign) error {
		return errors.New("error to send mail")
	}

	service.SendMail = sendMail

	err := service.Start(campaignPendenting.ID)
	assert.Equal(internalErrors.ErrInternal.Error(), err.Error())
}

func Test_Start_should_send_mail(t *testing.T) {
	SetUp()
	assert := assert.New(t)

	repositoryMock.On("GetBy", mock.Anything).Return(campaignPendenting, nil)
	emailWasSent := false
	sendMail := func(campaign *campaign.Campaign) error {
		if campaign.ID == campaignPendenting.ID {
			emailWasSent = true
		}
		return nil
	}
	repositoryMock.On("Update", mock.Anything).Return(nil)

	service.SendMail = sendMail

	service.Start(campaignPendenting.ID)
	assert.True(emailWasSent)
}

func Test_Start_ReturnNil_when_updated_to_done(t *testing.T) {
	SetUp()
	assert := assert.New(t)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignPendenting, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignPendenting.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)

	sendMail := func(campaign *campaign.Campaign) error {
		return nil
	}

	service.SendMail = sendMail

	service.Start(campaignPendenting.ID)
	assert.Equal(campaign.Done, campaignPendenting.Status)
}
