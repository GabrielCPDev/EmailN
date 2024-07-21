package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID, Name  string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index] = Contact{Email: email}
	}
	return &Campaign{ID: "1", Name: name, Content: content, Contacts: contacts, CreatedOn: time.Now()}
}
