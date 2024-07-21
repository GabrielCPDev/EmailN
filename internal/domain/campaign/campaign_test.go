package campaign

import "testing"

func TestNewCampaing(t *testing.T) {
	name := "Campaing X"
	content := "Body"
	contacts := []string{"user@user.com", "email@email.com"}

	campaing := NewCampaign(name, content, contacts)

	if campaing.ID != "1" {
		t.Errorf("expected 1")
	} else if campaing.Content != content {
		t.Errorf("expected correct name")
	} else if len(campaing.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}
}
