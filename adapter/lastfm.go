package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sigumaa/recenttrack/domain"
	"github.com/Sigumaa/recenttrack/repository"
	"github.com/samber/do"
)

type Lastfm struct {
	APIKey string
}

func NewRecentTracks(i *do.Injector) (repository.RecentTrackRepository, error) {

	apiKey := os.Getenv("LASTFM_API_KEY")

	return &Lastfm{
		APIKey: apiKey,
	}, nil
}

func (l *Lastfm) GetRecentTrack(user string) (domain.RecentTrackResponse, error) {

	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getRecentTracks&user=%s&api_key=%s&format=json", user, l.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return domain.RecentTrackResponse{}, err
	}
	defer resp.Body.Close()

	var data domain.RecentTrackData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println(err)
		return domain.RecentTrackResponse{}, err
	}

	log.Println("ok")

	return domain.RecentTrackResponse{
		Name:   data.Recenttrack.Track[0].Name,
		Album:  data.Recenttrack.Track[0].Album.Text,
		Artist: data.Recenttrack.Track[0].Artist.Text,
		Image:  data.Recenttrack.Track[0].Image[3].Text,
		URL:    data.Recenttrack.Track[0].URL,
	}, nil
}
