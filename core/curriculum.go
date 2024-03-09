package core

import "context"

type CurriculumService struct {
	repo CurriculumRepository
}

type CurriculumRepository interface {
	GetOneCurriculum(ctx context.Context, id int) (Curriculum, error)
	GetAllCurriculums(ctx context.Context) ([]Curriculum, error)
	CreateCurriculum(ctx context.Context, cr Curriculum) (Curriculum, error)
	UpdateCurriculum(ctx context.Context, cr Curriculum) error
	DeleteCurriculum(ctx context.Context, id int) error
}

func NewCurriculumService(repo CurriculumRepository) *CurriculumService {
	return &CurriculumService{repo: repo}
}

func (s *CurriculumService) GetOneCurriculum(ctx context.Context, id int) (Curriculum, error) {
	return s.repo.GetOneCurriculum(ctx, id)
}

func (s *CurriculumService) GetCurriculums(ctx context.Context) ([]Curriculum, error) {
	return s.repo.GetAllCurriculums(ctx)
}

func (s *CurriculumService) CreateCurriculum(ctx context.Context, curriculum Curriculum) (Curriculum, error) {
	return s.repo.CreateCurriculum(ctx, curriculum)
}

func (s *CurriculumService) UpdateCurriculum(ctx context.Context, id int, name string) error {
	curr, err := s.repo.GetOneCurriculum(ctx, id)
	if err != nil {
		return err
	}
	curr.Description = name
	return s.repo.UpdateCurriculum(ctx, curr)
}

func (s *CurriculumService) DeleteCurriculum(ctx context.Context, id int) (err error) {
	if _, err = s.repo.GetOneCurriculum(ctx, id); err != nil {
		return err
	}
	return s.repo.DeleteCurriculum(ctx, id)
}
