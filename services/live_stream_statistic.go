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
	UpsertBitrateIn(stream_key uuid.UUID, bitrate float64) error
	UpsertBitrateOut(stream_key uuid.UUID, bitrate float64) error
	UpsertFPSIn(stream_key uuid.UUID, fps int16) error
	UpsertFPSOut(stream_key uuid.UUID, fps int16) error
	UpsertNumberOfRequests(stream_key uuid.UUID, number_of_requests int) error
	UpsertDataTransferred(stream_key uuid.UUID, data_transferred float64) error
	UpsertDevice(stream_key uuid.UUID, device string) error
	UpsertOS(stream_key uuid.UUID, os string) error
	UpsertLocation(stream_key uuid.UUID, location string) error
}

func (s *StatisticService) UpsertBitrateOut(stream_key uuid.UUID, bitrate float64) error {
	return s.livestreamStatisticRepo.UpsertBitrateOut(stream_key, bitrate)
}

func (s *StatisticService) UpsertDataTransferred(stream_key uuid.UUID, dataTransfer float64) error {
	return s.livestreamStatisticRepo.UpsertDataTransferred(stream_key, dataTransfer)
}
func (s *StatisticService) UpsertFPSOut(stream_key uuid.UUID, fps int16) error {
	return s.livestreamStatisticRepo.UpsertFPSOut(stream_key, fps)
}

func (s *StatisticService) UpsertFPSIn(stream_key uuid.UUID, fps int16) error {
	return s.livestreamStatisticRepo.UpsertFPSIn(stream_key, fps)
}
