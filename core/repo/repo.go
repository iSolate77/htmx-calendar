package repo

import (
	"context"
	"errors"

	"github.com/iSolate77/htmx-calendar/core"
	"github.com/iSolate77/htmx-calendar/core/repo/db"
)

type repo struct {
	qr *db.Queries
}

func New(dbt db.DBTX) *repo {
	return &repo{
		qr: db.New(dbt),
	}
}

func (r *repo) CreateSubject(ctx context.Context, name string) (subject core.Subject, err error) {
	if name == "" {
		return subject, errors.New("name cannot be empty")
	}
	res, err := r.qr.CreateSubject(ctx, name)
	if err != nil {
		return subject, err
	}
	return core.Subject{
		ID:   int(res.Subjectid),
		Name: res.Name,
	}, nil

}

func (r *repo) UpdateSubject(ctx context.Context, subjectid int, name string) error {
	args := db.UpdateSubjectParams{
		Subjectid: int32(subjectid),
		Name:      name,
	}
	err := r.qr.UpdateSubject(ctx, args)
	return err
}

func (r *repo) DeleteSubject(ctx context.Context, subjectid int32) error {
	subject, err := r.qr.GetSubject(ctx, subjectid)
	if err != nil {
		return err
	}
	if subject.Subjectid == 0 {
		return errors.New("subject not found")
	}
	err = r.qr.DeleteSubject(ctx, subjectid)
	return err
}

func (r *repo) GetSubject(ctx context.Context, subjectid int) (db.Subject, error) {
	if subjectid == 0 {
		return db.Subject{}, errors.New("subject not found")
	}
	return r.qr.GetSubject(ctx, int32(subjectid))
}

func (r *repo) GetSubjects(ctx context.Context) ([]db.Subject, error) {
	return r.qr.ListSubjects(ctx)
}
