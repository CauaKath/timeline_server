package repo

import (
	"errors"
	"strconv"

	"github.com/cauakath/timeline-server/domain"
	"github.com/cauakath/timeline-server/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type timelineRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (n *timelineRepo) GetTimeline(timelineId int) (model.Timeline, error) {
	var timeline model.Timeline

	err := n.db.Model(model.Timeline{}).Select("id", "title", "type", "location", "start", "end").Where("id = ?", timelineId).Find(&timeline).Error

	if err != nil {
		return timeline, errors.New("failed to get timeline with id: " + strconv.Itoa(timelineId))
	}

	return timeline, nil
}

func (n *timelineRepo) ListTimelines() ([]model.Timeline, error) {
	var timelines []model.Timeline

	err := n.db.Model(model.Timeline{}).Select("id", "title", "type", "location", "start", "end").Find(&timelines).Error

	if err != nil {
		return timelines, errors.New("failed to get timelines from database")
	}

	return timelines, nil
}

func (n *timelineRepo) CreateTimeline(createTimeline model.Timeline) error {
	if err := n.db.Create(&createTimeline).Error; err != nil {
		return errors.New("failed to create timeline")
	}

	return nil
}

func (n *timelineRepo) UpdateTimeline(updateTimeline model.Timeline, timelineId int) error {
	if err := n.db.Where("id = ?", timelineId).Updates(&updateTimeline).Error; err != nil {
		return errors.New("failed to update timeline")
	}

	return nil
}

func NewTimelineRepo(db *gorm.DB, rdb *redis.Client) domain.TimelineRepo {
	return &timelineRepo{
		db:  db,
		rdb: rdb,
	}
}
