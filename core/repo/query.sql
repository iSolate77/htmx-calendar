-- name: GetSubject :one
SELECT * FROM Subjects
WHERE SubjectID = $1
LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM Subjects
ORDER BY Name;

-- name: CreateSubject :one
INSERT INTO Subjects (Name)
VALUES ($1)
RETURNING *;

-- name: UpdateSubject :exec
UPDATE Subjects
SET Name = $2
WHERE SubjectID = $1;

-- name: DeleteSubject :exec
DELETE FROM Subjects
WHERE SubjectID = $1;

-- name: GetAcademicCalendarForSemester :many
SELECT * FROM AcademicCalendar
WHERE SemesterID = $1
ORDER BY DayNumber ASC;

-- name: InsertAcademicCalendarEntry :exec
INSERT INTO AcademicCalendar (SemesterID, DayNumber, SubjectID, CurriculumID, ExamID, HolidayID, Description)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateAcademicCalendarEntry :exec
UPDATE AcademicCalendar
SET DayNumber = $2, SubjectID = $3, CurriculumID = $4, ExamID = $5, HolidayID = $6, Description = $7
WHERE CalendarID = $1;

-- name: DeleteAcademicCalendarEntry :exec
DELETE FROM AcademicCalendar
WHERE CalendarID = $1;

-- name: GetCurriculum :one
SELECT * FROM Curriculum
WHERE CurriculumID = $1
LIMIT 1;

-- name: ListCurriculum :many
SELECT * FROM Curriculum
ORDER BY Description;

-- name: CreateCurriculum :one
INSERT INTO Curriculum (SubjectID, Description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCurriculum :exec
UPDATE Curriculum
SET SubjectID = $2, Description = $3
WHERE CurriculumID = $1;

-- name: DeleteCurriculum :exec
DELETE FROM Curriculum
WHERE CurriculumID = $1;
