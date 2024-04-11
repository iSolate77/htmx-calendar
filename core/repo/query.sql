-- name: GetSubject :one
SELECT * FROM Subjects
WHERE id = $1
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
WHERE id = $1;

-- name: DeleteSubject :exec
DELETE FROM Subjects
WHERE id = $1;

-- name: GetAcademicCalendarForSemester :many
SELECT * FROM AcademicCalendar
WHERE Semester_id = $1
ORDER BY DayNumber ASC;

-- name: InsertAcademicCalendarEntry :exec
INSERT INTO AcademicCalendar (Semester_id, DayNumber, Subject_id, Curriculum_id, Exam_id, Holiday_id, Description)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateAcademicCalendarEntry :exec
UPDATE AcademicCalendar
SET DayNumber = $2, Subject_id = $3, Curriculum_id = $4, Exam_id = $5, Holiday_id = $6, Description = $7
WHERE id = $1;

-- name: DeleteAcademicCalendarEntry :exec
DELETE FROM AcademicCalendar
WHERE id = $1;

-- name: GetCurriculum :one
SELECT * FROM Curriculum
WHERE id = $1
LIMIT 1;

-- name: ListCurriculum :many
SELECT * FROM Curriculum
ORDER BY Description;

-- name: CreateCurriculum :one
INSERT INTO Curriculum (Subject_id, Description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCurriculum :exec
UPDATE Curriculum
SET Subject_id = $2, Description = $3
WHERE id = $1;

-- name: DeleteCurriculum :exec
DELETE FROM Curriculum
WHERE id = $1;

-- name: GetExam :one
SELECT * FROM Exams
WHERE id = $1
LIMIT 1;

-- name: ListExams :many
SELECT * FROM Exams
ORDER BY Date;

-- name: CreateExam :one
INSERT INTO Exams (Date, Subject_id, Description, Category_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateExam :exec
UPDATE Exams
SET Date = $2, Subject_id = $3, Description = $4, Category_id = $5
WHERE id = $1;

-- name: DeleteExam :exec
DELETE FROM Exams
WHERE id = $1;
