package services

import "github.com/google/uuid"

type StatisticService struct {
	livestreamStatisticRepo LiveStreamStatisticRepository
}

var LiveStreamStatisticService *StatisticService

func NewStatisticService(repo LiveStreamStatisticRepository) {
	LiveStreamStatisticService = &StatisticService{
		livestreamStatisticRepo: repo,
	}
}

type LiveStreamStatisticRepository interface {
	UpsertBitrateIn(streamKey uuid.UUID, bitrate float64) error
	UpsertBitrateOut(streamKey uuid.UUID, bitrate float64) error
	UpsertFPSIn(streamKey uuid.UUID, fps int16) error
	UpsertFPSOut(streamKey uuid.UUID, fps int16) error
	UpsertNumberOfRequests(streamKey uuid.UUID, numberOfRequests int) error
	UpsertDataTransferred(streamKey uuid.UUID, dataTransferred float64) error
}

func (s *StatisticService) UpsertBitrateOut(streamKey uuid.UUID, bitrate float64) error {
	return s.livestreamStatisticRepo.UpsertBitrateOut(streamKey, bitrate)
}

func (s *StatisticService) UpsertDataTransferred(streamKey uuid.UUID, dataTransfer float64) error {
	return s.livestreamStatisticRepo.UpsertDataTransferred(streamKey, dataTransfer)
}
func (s *StatisticService) UpsertFPSOut(streamKey uuid.UUID, fps int16) error {
	return s.livestreamStatisticRepo.UpsertFPSOut(streamKey, fps)
}

func (s *StatisticService) UpsertFPSIn(streamKey uuid.UUID, fps int16) error {
	return s.livestreamStatisticRepo.UpsertFPSIn(streamKey, fps)
}
