package core

import "context"

type HolidayRepository interface {
	GetOneHoliday(ctx context.Context, id int) (Holiday, error)
	GetAllHolidays(ctx context.Context) ([]Holiday, error)
	CreateHoliday(ctx context.Context, h Holiday) (Holiday, error)
	UpdateHoliday(ctx context.Context, id int, h Holiday) error
	DeleteHoliday(ctx context.Context, id int) error
}

type HolidayService struct {
	repo HolidayRepository
}

func NewHolidayService(repo HolidayRepository) *HolidayService {
	return &HolidayService{repo: repo}
}

func (s *HolidayService) GetOneHoliday(ctx context.Context, id int) (Holiday, error) {
	return s.repo.GetOneHoliday(ctx, id)
}

func (s *HolidayService) GetHolidays(ctx context.Context) ([]Holiday, error) {
	return s.repo.GetAllHolidays(ctx)
}

func (s *HolidayService) CreateHoliday(ctx context.Context, h Holiday) (Holiday, error) {
	return s.repo.CreateHoliday(ctx, h)
}

func (s *HolidayService) UpdateHoliday(ctx context.Context, id int, h Holiday) error {
	return s.repo.UpdateHoliday(ctx, id, h)
}

func (s *HolidayService) DeleteHoliday(ctx context.Context, id int) (err error) {
	return s.repo.DeleteHoliday(ctx, id)
}
