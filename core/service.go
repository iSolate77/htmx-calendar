package core

import (
	"context"
)

type Service struct {
	repo Repository
}

type Repository interface {
	GetSubject(ctx context.Context, id int) (Subject, error)
	GetSubjects(ctx context.Context) ([]Subject, error)
	CreateSubject(ctx context.Context, s string) (Subject, error)
	UpdateSubject(ctx context.Context, id int, name string) error
	DeleteSubject(ctx context.Context, id int) error
	GetCurriculum(ctx context.Context, id int) (Curriculum, error)
	GetCurriculums(ctx context.Context) ([]Curriculum, error)
	CreateCurriculum(ctx context.Context, c Curriculum) (int, error)
	UpdateCurriculum(ctx context.Context, c Curriculum) error
	DeleteCurriculum(ctx context.Context, id int) error
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetSubject(ctx context.Context, id int) (subject Subject, err error) {
	sub, err := s.repo.GetSubject(ctx, id)
	if err != nil {
		return subject, err
	}
	return sub, nil
}

func (s *Service) GetSubjects(ctx context.Context) ([]Subject, error) {
	return s.repo.GetSubjects(ctx)
}

func (s *Service) CreateSubject(ctx context.Context, name string) (Subject, error) {
	return s.repo.CreateSubject(ctx, name)
}

func (s *Service) UpdateSubject(ctx context.Context, id int, name string) error {
	subject, err := s.repo.GetSubject(ctx, id)
	if err != nil {
		return err
	}
	subject.Name = name
	return s.repo.UpdateSubject(ctx, id, name)
}

func (s *Service) DeleteSubject(ctx context.Context, id int) (err error) {
	if _, err = s.repo.GetSubject(ctx, id); err != nil {
		return err
	}
	return s.repo.DeleteSubject(ctx, id)
}
