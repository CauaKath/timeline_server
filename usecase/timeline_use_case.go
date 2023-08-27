package usecase

import (
	"errors"

	"github.com/cauakath/timeline-server/domain"
	"github.com/cauakath/timeline-server/model"
)

type timelineUseCase struct {
	timelineRepo domain.TimelineRepo
}

func (n *timelineUseCase) ListTimelines() ([]model.Timeline, error) {
	timelines, err := n.timelineRepo.ListTimelines()

	if err != nil {
		return []model.Timeline{}, errors.New("failed to get timelines: " + err.Error())
	}

	return timelines, nil
}

func (n *timelineUseCase) CreateTimeline(createTimeline model.Timeline) error {
	err := n.timelineRepo.CreateTimeline(createTimeline)

	if err != nil {
		return errors.New("failed to create timeline: " + err.Error())
	}

	return nil
}

func NewTimelineUseCase(timelineRepo domain.TimelineRepo) domain.TimelineUseCase {
	return &timelineUseCase{
		timelineRepo: timelineRepo,
	}
}
