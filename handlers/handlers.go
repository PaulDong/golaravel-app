package handlers

import (
	"bnlogic/data"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/PaulDong/golaravel"
)

type Handlers struct {
	App    *golaravel.Golaravel
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("Error rendering:", err)
	}
}
