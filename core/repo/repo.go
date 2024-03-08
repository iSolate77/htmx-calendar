package repo

import (
	"context"

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

func (r *repo) CreateSubject(ctx context.Context, name string) (db.Subject, error) {
	return r.qr.CreateSubject(ctx, name)
}

func (r *repo) UpdateSubject(ctx context.Context, arg db.UpdateSubjectParams) error {
	return r.qr.UpdateSubject(ctx, arg)
}

func (r *repo) DeleteSubject(ctx context.Context, subjectid int32) error {
	return r.qr.DeleteSubject(ctx, subjectid)
}

func (r *repo) GetSubject(ctx context.Context, subjectid int32) (db.Subject, error) {
	return r.qr.GetSubject(ctx, subjectid)
}

func (r *repo) ListSubjects(ctx context.Context) ([]db.Subject, error) {
	return r.qr.ListSubjects(ctx)
}
