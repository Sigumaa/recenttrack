package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sigumaa/recenttrack/repository"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type RecentTrackController struct {
	r repository.RecentTrackRepository
}

func NewRecentTrackController(i *do.Injector) (RecentTrackController, error) {
	r := do.MustInvoke[repository.RecentTrackRepository](i)

	return RecentTrackController{r}, nil
}

func (c RecentTrackController) GetRecentTrack(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	data, err := c.r.GetRecentTrack(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (c RecentTrackController) RegisterEndpoints(r *chi.Mux) {
	r.Get("/recenttrack", c.GetRecentTrack)
}
