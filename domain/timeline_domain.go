package domain

import "github.com/cauakath/timeline-server/model"

type TimelineRepo interface {
	CreateTimeline(createTimeline model.Timeline) error
	UpdateTimeline(updateTimeline model.Timeline, timelineId int) error
	GetTimeline(timelineId int) (model.Timeline, error)
	ListTimelines() ([]model.Timeline, error)
	DeleteTimeline(timelineId int) error
}

type TimelineUseCase interface {
	CreateTimeline(createTimeline model.Timeline) error
	UpdateTimeline(updateTimeline model.Timeline, timelineId int) error
	GetTimeline(timelineId int) (model.Timeline, error)
	ListTimelines() ([]model.Timeline, error)
	DeleteTimeline(timelineId int) error
}
