package domain

type RecentTrackData struct {
	Recenttrack struct {
		Track []struct {
			Artist struct {
				Mbid string `json:"mbid"`
				Text string `json:"#text"`
			} `json:"artist"`
			Streamable string `json:"streamable"`
			Image      []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Mbid  string `json:"mbid"`
			Album struct {
				Mbid string `json:"mbid"`
				Text string `json:"#text"`
			} `json:"album"`
			Name string `json:"name"`
			URL  string `json:"url"`
			Date struct {
				Uts  string `json:"uts"`
				Text string `json:"#text"`
			} `json:"date"`
		} `json:"track"`
		Attr struct {
			User       string `json:"user"`
			TotalPages string `json:"totalPages"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"recenttracks"`
}

type RecentTrackResponse struct {
	Name   string `json:"name"`
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Image  string `json:"image"`
	URL    string `json:"url"`
}
