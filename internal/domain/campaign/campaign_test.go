package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing X"
	content  = "Body"
	contacts = []string{"user@user.com", "email@email.com"}
)

func Test_NewCampaing_CreateNewCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _:= NewCampaign(name, content, contacts)

	assert.Equal(campaing.ID, "1")
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaign(name, content, contacts)
	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaign(name, content, contacts)
	assert.Greater(campaing.CreatedOn, now)
}

func Test_NewCampaing_MustValidadeName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)
	assert.Equal("name is required", err.Error())
}

func Test_NewCampaing_MustValidadeContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("content is required", err.Error())
}

func Test_NewCampaing_MustValidadeContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})
	assert.Equal("contacts is required", err.Error())
}