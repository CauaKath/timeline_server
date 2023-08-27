package repo

import (
	"errors"

	"github.com/cauakath/timeline-server/domain"
	"github.com/cauakath/timeline-server/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type timelineRepo struct {
	db  *gorm.DB
	rdb *redis.Client
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

func NewTimelineRepo(db *gorm.DB, rdb *redis.Client) domain.TimelineRepo {
	return &timelineRepo{
		db:  db,
		rdb: rdb,
	}
}
