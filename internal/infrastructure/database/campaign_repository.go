package database

import (
	"gomail/internal/domain/campaign"
)

type CampaignRepository struct {
	campaign []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaign = append(c.campaign, *campaign)

	return nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
	return c.campaign
}
