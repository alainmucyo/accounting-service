package channel

import (
	"accounting-service/store/entities/channel"
	"accounting-service/store/postgres"
	"errors"
)

type Service struct {
	db *postgres.Database
}

func New(db *postgres.Database) *Service {
	return &Service{db: db}
}

func (s *Service) FindAll() ([]channel.Channel, error) {
	channels := make([]channel.Channel, 0)
	if s.db.DB.Find(&channels).Error != nil {
		return nil, errors.New("error while getting all channels")
	}
	return channels, nil
}

func (s *Service) FindByName(channelName string) (channel.Channel, error) {
	var foundChannel channel.Channel
	if s.db.DB.Where(&channel.Channel{Name: channelName}).First(&foundChannel).Error != nil {
		return channel.Channel{}, errors.New("channel not found")
	}
	return foundChannel, nil
}

func (s *Service) Create(channel channel.Channel) (channel.Channel, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return channel, errors.New("error start")
	}
	if s.db.DB.Save(&channel).Error != nil {
		ctx.Rollback()
		return channel, errors.New("error save")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return channel, errors.New("error commit")
	}
	return channel, nil
}
