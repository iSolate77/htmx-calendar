package repo

import (
	"context"

	"github.com/iSolate77/htmx-calendar/core"
	"github.com/iSolate77/htmx-calendar/core/repo/db"
)

type curriculumRepo struct {
	qr *db.Queries
}

func NewCurriculumRepo(dbt db.DBTX) *subjectRepo {
	return &subjectRepo{
		qr: db.New(dbt),
	}
}

func (r *curriculumRepo) CreateCurriculum(ctx context.Context, args core.Curriculum) (curriculum core.Curriculum, err error) {
	res, err := r.qr.CreateCurriculum(ctx, db.CreateCurriculumParams{
		Subjectid:   int32(args.SubjectID),
		Description: args.Description,
	})
	if err != nil {
		return curriculum, err
	}
	return core.Curriculum{
		SubjectID:   int(res.Subjectid),
		Description: res.Description,
	}, nil
}

func (r *curriculumRepo) DeleteCurriculum(ctx context.Context, curriculumID int) error {
	return r.qr.DeleteCurriculum(ctx, int32(curriculumID))
}

func (r *curriculumRepo) GetCurriculum(ctx context.Context, curriculumID int) (curriculum core.Curriculum, err error) {
	res, err := r.qr.GetCurriculum(ctx, int32(curriculumID))
	if err != nil {
		return curriculum, err
	}
	return core.Curriculum{
		ID:          int(res.Curriculumid),
		SubjectID:   int(res.Subjectid),
		Description: res.Description,
	}, nil
}

func (r *curriculumRepo) UpdateCurriculum(ctx context.Context, args core.Curriculum) error {
	err := r.qr.UpdateCurriculum(ctx, db.UpdateCurriculumParams{
		Curriculumid: int32(args.ID),
		Subjectid:    int32(args.SubjectID),
		Description:  args.Description,
	})
	return err
}

func (r *curriculumRepo) ListCurriculums(ctx context.Context) (curriculums []core.Curriculum, err error) {
	res, err := r.qr.ListCurriculum(ctx)
	if err != nil {
		return curriculums, err
	}
	for _, v := range res {
		curriculums = append(curriculums, core.Curriculum{
			ID:          int(v.Curriculumid),
			SubjectID:   int(v.Subjectid),
			Description: v.Description,
		})
	}
	return curriculums, nil
}
