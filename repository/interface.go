package repository

import "github.com/Sigumaa/recenttrack/domain"

type RecentTrackRepository interface {
	GetRecentTrack(user string) (domain.RecentTrackResponse, error)
}
