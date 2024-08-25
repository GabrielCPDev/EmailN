package campaign

import (
	"emailn/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started        = "Started"
	Done           = "Done"
)

type Contact struct {
	ID         string
	Email      string `validate:"email"`
	CampaignId string
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}
	campaing := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
		Status:    Pending,
	}
	err := internalErrors.ValidateStruct(campaing)

	if err == nil {
		return campaing, nil
	}
	return nil, err
}
