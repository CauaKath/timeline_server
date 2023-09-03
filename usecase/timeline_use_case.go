package usecase

import (
	"errors"

	"github.com/cauakath/timeline-server/domain"
	"github.com/cauakath/timeline-server/model"
)

type timelineUseCase struct {
	timelineRepo domain.TimelineRepo
}

func (n *timelineUseCase) GetTimeline(timelineId int) (model.Timeline, error) {
	timeline, err := n.timelineRepo.GetTimeline(timelineId)

	if err != nil {
		return model.Timeline{}, errors.New("failed to get timeline: " + err.Error())
	}

	return timeline, nil
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

func (n *timelineUseCase) UpdateTimeline(updateTimeline model.Timeline, timelineId int) error {
	err := n.timelineRepo.UpdateTimeline(updateTimeline, timelineId)

	if err != nil {
		return errors.New("failed to update timeline: " + err.Error())
	}

	return nil
}

func (n *timelineUseCase) DeleteTimeline(timelineId int) error {
	err := n.timelineRepo.DeleteTimeline(timelineId)

	if err != nil {
		return errors.New("failed to delete timeline: " + err.Error())
	}

	return nil
}

func NewTimelineUseCase(timelineRepo domain.TimelineRepo) domain.TimelineUseCase {
	return &timelineUseCase{
		timelineRepo: timelineRepo,
	}
}
