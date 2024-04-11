package repo

import (
	"context"
	"github.com/iSolate77/htmx-calendar/core"
	"github.com/iSolate77/htmx-calendar/core/repo/db"
)

type examRepo struct {
	qr *db.Queries
}

func NewExamRepo(dbt db.DBTX) *examRepo {
	return &examRepo{
		qr: db.New(dbt),
	}
}

func (r *examRepo) GetOneExam(ctx context.Context, id int) (exam core.Exam, err error) {
	if id == 0 {
		return exam, core.ErrNotFound
	}
	res, err := r.qr.GetExam(ctx, int32(id))
	if err != nil {
		return exam, err
	}
	return core.Exam{
		ID:          int(res.ID),
		SubjectID:   int(res.SubjectID),
		CategoryID:  int(res.CategoryID),
		Description: res.Description,
		Date:        res.Date,
	}, nil
}

func (r *examRepo) GetAllExams(ctx context.Context) (exams []core.Exam, err error) {
	res, err := r.qr.ListExams(ctx)
	if err != nil {
		return nil, err
	}
	for _, e := range res {
		exams = append(exams, core.Exam{
			ID:          int(e.ID),
			SubjectID:   int(e.SubjectID),
			CategoryID:  int(e.CategoryID),
			Description: e.Description,
			Date:        e.Date,
		})
	}
	return exams, nil
}

func (r *examRepo) CreateExam(ctx context.Context, e core.Exam) (exam core.Exam, err error) {
	if e.Description == "" {
		return exam, core.ErrEmpty
	}
	res, err := r.qr.CreateExam(ctx, db.CreateExamParams{
		SubjectID:   int32(e.SubjectID),
		CategoryID:  int32(e.CategoryID),
		Description: e.Description,
		Date:        e.Date,
	})
	if err != nil {
		return exam, err
	}
	return core.Exam{
		ID:          int(res.ID),
		SubjectID:   int(res.SubjectID),
		CategoryID:  int(res.CategoryID),
		Description: res.Description,
		Date:        res.Date,
	}, nil
}

func (r *examRepo) UpdateExam(ctx context.Context, id int, e core.Exam) error {
	args := db.UpdateExamParams{
		ID:          int32(id),
		SubjectID:   int32(e.SubjectID),
		CategoryID:  int32(e.CategoryID),
		Description: e.Description,
		Date:        e.Date,
	}
	err := r.qr.UpdateExam(ctx, args)
	return err
}

func (r *examRepo) DeleteExam(ctx context.Context, id int) error {
	exam, err := r.qr.GetExam(ctx, int32(id))
	if err != nil {
		return err
	}
	if exam.ID == 0 {
		return core.ErrNotFound
	}
	err = r.qr.DeleteExam(ctx, int32(id))
	return err
}
