package repo

import (
	"context"

	"github.com/iSolate77/htmx-calendar/core"
	"github.com/iSolate77/htmx-calendar/core/repo/db"
)

type subjectRepo struct {
	qr *db.Queries
}

func NewSubjectRepo(dbt db.DBTX) *subjectRepo {
	return &subjectRepo{
		qr: db.New(dbt),
	}
}

func (r *subjectRepo) CreateSubject(ctx context.Context, name string) (subject core.Subject, err error) {
	if name == "" {
		return subject, core.ErrEmpty
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

func (r *subjectRepo) UpdateSubject(ctx context.Context, subjectid int, name string) error {
	args := db.UpdateSubjectParams{
		Subjectid: int32(subjectid),
		Name:      name,
	}
	err := r.qr.UpdateSubject(ctx, args)
	return err
}

func (r *subjectRepo) DeleteSubject(ctx context.Context, id int) error {
	subject, err := r.qr.GetSubject(ctx, int32(id))
	if err != nil {
		return err
	}
	if subject.Subjectid == 0 {
		return core.ErrNotFound
	}
	err = r.qr.DeleteSubject(ctx, int32(id))
	return err
}

func (r *subjectRepo) GetOneSubject(ctx context.Context, id int) (subject core.Subject, err error) {
	if id == 0 {
		return subject, core.ErrNotFound
	}
	res, err := r.qr.GetSubject(ctx, int32(id))
	if err != nil {
		return subject, err
	}
	subject = core.Subject{
		ID:   int(res.Subjectid),
		Name: res.Name,
	}
	return subject, nil
}

func (r *subjectRepo) GetAllSubjects(ctx context.Context) ([]core.Subject, error) {
	res, err := r.qr.ListSubjects(ctx)
	if err != nil {
		return nil, err
	}
	var subjects []core.Subject
	for _, s := range res {
		subjects = append(subjects, core.Subject{
			Name: s.Name,
		})
	}
	return subjects, nil
}
