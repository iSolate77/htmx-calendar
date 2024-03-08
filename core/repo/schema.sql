CREATE TABLE IF NOT EXISTS AcademicYears (
    YearID SERIAL PRIMARY KEY,
    StartDate DATE NOT NULL,
    EndDate DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS Semesters (
    SemesterID SERIAL PRIMARY KEY,
    YearID INTEGER NOT NULL REFERENCES AcademicYears(YearID),
    Name VARCHAR(255) NOT NULL,
    StartDate DATE NOT NULL,
    WeeksDuration INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS Subjects (
    SubjectID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Categories (
    CategoryID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Curriculum (
    CurriculumID SERIAL PRIMARY KEY,
    SubjectID INTEGER NOT NULL REFERENCES Subjects(SubjectID),
    Description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Exams (
    ExamID SERIAL PRIMARY KEY,
    SubjectID INTEGER NOT NULL REFERENCES Subjects(SubjectID),
    CategoryID INTEGER NOT NULL REFERENCES Categories(CategoryID),
    Description TEXT NOT NULL,
    Date VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Holidays (
    HolidayID SERIAL PRIMARY KEY,
    CategoryID INTEGER NOT NULL REFERENCES Categories(CategoryID),
    Date VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS AcademicCalendar (
    CalendarID SERIAL PRIMARY KEY,
    SemesterID INTEGER NOT NULL REFERENCES Semesters(SemesterID),
    DayNumber INTEGER NOT NULL,
    SubjectID INTEGER REFERENCES Subjects(SubjectID),
    CurriculumID INTEGER REFERENCES Curriculum(CurriculumID),
    ExamID INTEGER REFERENCES Exams(ExamID),
    HolidayID INTEGER REFERENCES Holidays(HolidayID),
    Description TEXT
);

CREATE TABLE IF NOT EXISTS CurriculumSchedule (
    ScheduleID SERIAL PRIMARY KEY,
    CurriculumID INTEGER NOT NULL REFERENCES Curriculum(CurriculumID),
    SubjectID INTEGER NOT NULL REFERENCES Subjects(SubjectID),
    SemesterID INTEGER NOT NULL REFERENCES Semesters(SemesterID),
    DayOffset INTEGER NOT NULL,
    Description TEXT NOT NULL,
    UNIQUE (CurriculumID, SubjectID, SemesterID, DayOffset)
);
