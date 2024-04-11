package repo

import (
	"context"
	"github.com/iSolate77/htmx-calendar/core"
	"github.com/iSolate77/htmx-calendar/core/repo/db"
)

type holidayRepo struct {
	qr *db.Queries
}

func NewHolidayRepo(dbt db.DBTX) *holidayRepo {
	return &holidayRepo{
		qr: db.New(dbt),
	}
}

func (r *holidayRepo) GetOneHoliday(ctx context.Context, id int) (holiday core.Holiday, err error) {
	if id == 0 {
		return holiday, core.ErrNotFound
	}
	res, err := r.qr.GetHoliday(ctx, int32(id))
	if err != nil {
		return holiday, err
	}
	return core.Holiday{
		ID:          int(res.ID),
		Name:        res.Name,
		Description: res.Description,
		Date:        res.Date,
	}, nil
}

func (r *holidayRepo) GetAllHolidays(ctx context.Context) (holidays []core.Holiday, err error) {
	res, err := r.qr.ListHolidays(ctx)
	if err != nil {
		return nil, err
	}
	for _, h := range res {
		holidays = append(holidays, core.Holiday{
			ID:          int(h.ID),
			Name:        h.Name,
			Description: h.Description,
			Date:        h.Date,
		})
	}
	return holidays, nil
}

func (r *holidayRepo) CreateHoliday(ctx context.Context, h core.Holiday) (holiday core.Holiday, err error) {
	if h.Name == "" {
		return holiday, core.ErrEmpty
	}
	res, err := r.qr.CreateHoliday(ctx, db.CreateHolidayParams{
		Name:        h.Name,
		Description: h.Description,
		Date:        h.Date,
	})
	if err != nil {
		return holiday, err
	}
	return core.Holiday{
		ID:          int(res.ID),
		Name:        res.Name,
		Description: res.Description,
		Date:        res.Date,
	}, nil
}

func (r *holidayRepo) UpdateHoliday(ctx context.Context, h core.Holiday) (holiday core.Holiday, err error) {
	if h.Name == "" {
		return holiday, core.ErrEmpty
	}
	res, err := r.qr.UpdateHoliday(ctx, db.UpdateHolidayParams{
		ID:          int32(h.ID),
		Name:        h.Name,
		Description: h.Description,
		Date:        h.Date,
	})
	if err != nil {
		return holiday, err
	}
	return core.Holiday{
		ID:          int(res.ID),
		Name:        res.Name,
		Description: res.Description,
		Date:        res.Date,
	}, nil
}

func (r *holidayRepo) DeleteHoliday(ctx context.Context, id int) error {
	if id == 0 {
		return core.ErrNotFound
	}
	_, err := r.qr.DeleteHoliday(ctx, int32(id))
	return err
}
