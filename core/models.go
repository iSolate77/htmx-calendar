package core

type subject struct {
	ID   int
	Name string
}

type exam struct {
	ID          int
	subjectID   int
	categoryID  int
	Description string
	Date        string
}

type holiday struct {
	ID         int
	categoryID int
	Date       string
}

type semester struct {
	ID            int
	yearID        int
	Name          string
	StartDate     string
	WeeksDuration int
}

type academicCalendar struct {
	ID           int
	semesterID   int
	DayNumber    int
	subjectID    int
	curriculumID int
	examID       int
	holidayID    int
	Description  string
}

type academicYear struct {
	ID        int
	StartDate string
	EndDate   string
}

type category struct {
	ID   int
	Name string
}

type curriculum struct {
	ID          int
	subjectID   int
	Description string
}

type curriculumSchedule struct {
	ID           int
	curriculumID int
	subjectID    int
	semesterID   int
	Dayoffset    int
	description  string
}
