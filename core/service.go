package core

import (
	"context"

	"github.com/iSolate77/htmx-calendar/core/repo/db"
	"github.com/jackc/pgtype"
)

type service struct {
	repo repo
}

type repo interface {
	CreateSubject(ctx context.Context, name string) (db.Subject, error)
	UpdateSubject(ctx context.Context, arg db.UpdateSubjectParams) error
	DeleteSubject(ctx context.Context, subjectid int32) error
	GetSubject(ctx context.Context, subjectid int32) (db.Subject, error)
	ListSubjects(ctx context.Context) ([]db.Subject, error)
	CreateAcademicYear(ctx context.Context, startdate pgtype.Date, enddate pgtype.Date) (db.Academicyear, error)
	// UpdateAcademicYear(ctx context.Context, arg db.UpdateAcademicYearParams) error
	DeleteAcademicYear(ctx context.Context, yearid int32) error
	GetAcademicYear(ctx context.Context, yearid int32) (db.Academicyear, error)
	ListAcademicYears(ctx context.Context) ([]db.Academicyear, error)
	CreateCategory(ctx context.Context, name string) (db.Category, error)
	// UpdateCategory(ctx context.Context, arg db.UpdateCategoryParams) error
	DeleteCategory(ctx context.Context, categoryid int32) error
	GetCategory(ctx context.Context, categoryid int32) (db.Category, error)
	ListCategories(ctx context.Context) ([]db.Category, error)
	CreateCurriculum(ctx context.Context, subjectid int32, description string) (db.Curriculum, error)
	// UpdateCurriculum(ctx context.Context, arg db.UpdateCurriculumParams) error
	DeleteCurriculum(ctx context.Context, curriculumid int32) error
	GetCurriculum(ctx context.Context, curriculumid int32) (db.Curriculum, error)
	ListCurriculums(ctx context.Context) ([]db.Curriculum, error)
	CreateExam(ctx context.Context, subjectid int32, categoryid int32, description string, date string) (db.Exam, error)
	// UpdateExam(ctx context.Context, arg db.UpdateExamParams) error
	DeleteExam(ctx context.Context, examid int32) error
	GetExam(ctx context.Context, examid int32) (db.Exam, error)
	ListExams(ctx context.Context) ([]db.Exam, error)
	CreateHoliday(ctx context.Context, categoryid int32, date string) (db.Holiday, error)
	// UpdateHoliday(ctx context.Context, arg db.UpdateHolidayParams) error
	DeleteHoliday(ctx context.Context, holidayid int32) error
	GetHoliday(ctx context.Context, holidayid int32) (db.Holiday, error)
	ListHolidays(ctx context.Context) ([]db.Holiday, error)
	CreateSemester(ctx context.Context, yearid int32, name string, startdate pgtype.Date, weeksduration int32) (db.Semester, error)
	// UpdateSemester(ctx context.Context, arg db.UpdateSemesterParams) error
	DeleteSemester(ctx context.Context, semesterid int32) error
	GetSemester(ctx context.Context, semesterid int32) (db.Semester, error)
	ListSemesters(ctx context.Context) ([]db.Semester, error)
	CreateAcademicCalendar(ctx context.Context, semesterid int32, daynumber int32, subjectid pgtype.Int4, curriculumid pgtype.Int4, examid pgtype.Int4, holidayid pgtype.Int4, description pgtype.Text) (db.Academiccalendar, error)
	// UpdateAcademicCalendar(ctx context.Context, arg db.UpdateAcademicCalendarParams) error
	DeleteAcademicCalendar(ctx context.Context, calendarid int32) error
	GetAcademicCalendar(ctx context.Context, calendarid int32) (db.Academiccalendar, error)
	ListAcademicCalendar(ctx context.Context) ([]db.Academiccalendar, error)
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}

func (s *service) CreateSubject(ctx context.Context, name string) (db.Subject, error) {
	return s.repo.CreateSubject(ctx, name)
}

func (s *service) UpdateSubject(ctx context.Context, arg db.UpdateSubjectParams) error {
	return s.repo.UpdateSubject(ctx, arg)
}

func (s *service) DeleteSubject(ctx context.Context, subjectid int32) error {
	return s.repo.DeleteSubject(ctx, subjectid)
}

func (s *service) GetSubject(ctx context.Context, subjectid int32) (db.Subject, error) {
	return s.repo.GetSubject(ctx, subjectid)
}

func (s *service) ListSubjects(ctx context.Context) ([]db.Subject, error) {
	return s.repo.ListSubjects(ctx)
}

func (s *service) CreateAcademicYear(ctx context.Context, startdate pgtype.Date, enddate pgtype.Date) (db.Academicyear, error) {
	return s.repo.CreateAcademicYear(ctx, startdate, enddate)
}

func (s *service) DeleteAcademicYear(ctx context.Context, yearid int32) error {
	return s.repo.DeleteAcademicYear(ctx, yearid)
}

func (s *service) GetAcademicYear(ctx context.Context, yearid int32) (db.Academicyear, error) {
	return s.repo.GetAcademicYear(ctx, yearid)
}

func (s *service) ListAcademicYears(ctx context.Context) ([]db.Academicyear, error) {
	return s.repo.ListAcademicYears(ctx)
}

func (s *service) CreateCategory(ctx context.Context, name string) (db.Category, error) {
	return s.repo.CreateCategory(ctx, name)
}

func (s *service) DeleteCategory(ctx context.Context, categoryid int32) error {
	return s.repo.DeleteCategory(ctx, categoryid)
}

func (s *service) GetCategory(ctx context.Context, categoryid int32) (db.Category, error) {
	return s.repo.GetCategory(ctx, categoryid)
}

func (s *service) ListCategories(ctx context.Context) ([]db.Category, error) {
	return s.repo.ListCategories(ctx)
}

func (s *service) CreateCurriculum(ctx context.Context, subjectid int32, description string) (db.Curriculum, error) {
	return s.repo.CreateCurriculum(ctx, subjectid, description)
}

func (s *service) DeleteCurriculum(ctx context.Context, curriculumid int32) error {
	return s.repo.DeleteCurriculum(ctx, curriculumid)
}

func (s *service) GetCurriculum(ctx context.Context, curriculumid int32) (db.Curriculum, error) {
	return s.repo.GetCurriculum(ctx, curriculumid)
}

func (s *service) ListCurriculums(ctx context.Context) ([]db.Curriculum, error) {
	return s.repo.ListCurriculums(ctx)
}
