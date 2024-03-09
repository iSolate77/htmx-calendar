package core

type Subject struct {
	ID   int
	Name string
}

type Exam struct {
	ID          int
	subjectID   int
	categoryID  int
	Description string
	Date        string
}

type Holiday struct {
	ID         int
	categoryID int
	Date       string
}

type Semester struct {
	ID            int
	yearID        int
	Name          string
	StartDate     string
	WeeksDuration int
}

type AcademicCalendar struct {
	ID           int
	semesterID   int
	DayNumber    int
	subjectID    int
	curriculumID int
	examID       int
	holidayID    int
	Description  string
}

type AcademicYear struct {
	ID        int
	StartDate string
	EndDate   string
}

type Category struct {
	ID   int
	Name string
}

type Curriculum struct {
	ID          int
	SubjectID   int
	Description string
}

type CurriculumSchedule struct {
	ID           int
	curriculumID int
	subjectID    int
	semesterID   int
	Dayoffset    int
	description  string
}
