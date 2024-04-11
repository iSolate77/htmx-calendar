package core

import "context"

type ExamRepository interface {
	GetOneExam(ctx context.Context, id int) (Exam, error)
	GetAllExams(ctx context.Context) ([]Exam, error)
	CreateExam(ctx context.Context, e Exam) (Exam, error)
	UpdateExam(ctx context.Context, id int, e Exam) error
	DeleteExam(ctx context.Context, id int) error
}

type ExamService struct {
	repo ExamRepository
}

func NewExamService(repo ExamRepository) *ExamService {
	return &ExamService{repo: repo}
}

func (s *ExamService) GetOneExam(ctx context.Context, id int) (Exam, error) {
	return s.repo.GetOneExam(ctx, id)
}

func (s *ExamService) GetExams(ctx context.Context) ([]Exam, error) {
	return s.repo.GetAllExams(ctx)
}

func (s *ExamService) CreateExam(ctx context.Context, e Exam) (Exam, error) {
	return s.repo.CreateExam(ctx, e)
}

func (s *ExamService) UpdateExam(ctx context.Context, id int, e Exam) error {
	return s.repo.UpdateExam(ctx, id, e)
}

func (s *ExamService) DeleteExam(ctx context.Context, id int) (err error) {
	return s.repo.DeleteExam(ctx, id)
}
