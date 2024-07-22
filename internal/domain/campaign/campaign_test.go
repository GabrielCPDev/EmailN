package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaing(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing X"
	content := "Body"
	contacts := []string{"user@user.com", "email@email.com"}

	campaing := NewCampaign(name, content, contacts)

	assert.Equal(campaing.ID, "1")
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}
