package domain

import "github.com/cauakath/timeline-server/model"

type TimelineRepo interface {
	CreateTimeline(createTimeline model.Timeline) error
	ListTimelines() ([]model.Timeline, error)
}

type TimelineUseCase interface {
	CreateTimeline(createTimeline model.Timeline) error
	ListTimelines() ([]model.Timeline, error)
}
