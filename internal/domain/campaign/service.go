package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {
	campaign, err := NewCampaign(newCampaing.Nome, newCampaing.Content, newCampaing.Emails)
	s.Repository.Save(campaign)
	return campaign.ID, err
}
