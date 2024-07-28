package endpoints

import "gomail/internal/domain/campaign"

type Handler struct {
	CampaingService campaign.Service
}
