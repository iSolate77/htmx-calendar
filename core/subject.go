package core

import "context"

type SubjectRepository interface {
	GetOneSubject(ctx context.Context, id int) (Subject, error)
	GetAllSubjects(ctx context.Context) ([]Subject, error)
	CreateSubject(ctx context.Context, s string) (Subject, error)
	UpdateSubject(ctx context.Context, id int, name string) error
	DeleteSubject(ctx context.Context, id int) error
}

type SubjectService struct {
	repo SubjectRepository
}


func NewSubjectService(repo SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (s *SubjectService) GetOneSubject(ctx context.Context, id int) (Subject, error) {
	return s.repo.GetOneSubject(ctx, id)
}

func (s *SubjectService) GetSubjects(ctx context.Context) ([]Subject, error) {
	return s.repo.GetAllSubjects(ctx)
}

func (s *SubjectService) CreateSubject(ctx context.Context, name string) (Subject, error) {
	return s.repo.CreateSubject(ctx, name)
}

func (s *SubjectService) UpdateSubject(ctx context.Context, id int, name string) error {
	subject, err := s.repo.GetOneSubject(ctx, id)
	if err != nil {
		return err
	}
	subject.Name = name
	return s.repo.UpdateSubject(ctx, id, name)
}

func (s *SubjectService) DeleteSubject(ctx context.Context, id int) (err error) {
	if _, err = s.repo.GetOneSubject(ctx, id); err != nil {
		return err
	}
	return s.repo.DeleteSubject(ctx, id)
}
