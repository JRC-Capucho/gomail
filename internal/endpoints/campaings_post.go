package endpoints

import (
	"errors"
	"gomail/internal/contract"
	internalerrors "gomail/internal/internalErrors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign

	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaingService.Create(request)
	if err != nil {
		code := 400
		if errors.Is(err, internalerrors.ErrInternal) {
			code = 500
		}
		render.Status(r, code)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
