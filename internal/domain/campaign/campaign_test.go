package campaign_test

import (
	"emailn/internal/domain/campaign"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	fake = faker.New()

	name      = "Campaing X"
	content   = fake.Lorem().Text(1000)
	createdBy = "teste@teste.com.br"
	contacts  = []string{"user@user.com", "email@email.com"}
)

func Test_NewCampaing_CreateNewCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := campaign.NewCampaign(name, content, contacts, createdBy)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := campaign.NewCampaign(name, content, contacts, createdBy)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaing, _ := campaign.NewCampaign(name, content, contacts, createdBy)
	assert.Greater(campaing.CreatedOn, now)
}

func Test_NewCampaing_MustStatusStartWithPending(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := campaign.NewCampaign(name, content, contacts, createdBy)
	assert.Equal(campaign.Pending, campaing.Status)
}

func Test_NewCampaing_MustValidadeNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign("", content, contacts, createdBy)
	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaing_MustValidadeNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := campaign.NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)
	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaing_MustValidadeContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, "", contacts, createdBy)
	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaing_MustValidadeContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, fake.Lorem().Text(1050), contacts, createdBy)
	assert.Equal("content is required with max 1024", err.Error())
}
func Test_NewCampaing_MustValidadeContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, []string{"email_invalid"}, createdBy)
	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaing_MustValidadeContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, nil, createdBy)
	assert.Equal("contacts is required with min 1", err.Error())
}
