CREATE TABLE IF NOT EXISTS AcademicYears (
    ID SERIAL PRIMARY KEY,
    StartDate DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS Semesters (
    ID SERIAL PRIMARY KEY,
    Year_ID INTEGER NOT NULL REFERENCES AcademicYears(ID),
    Name VARCHAR(255) NOT NULL,
    StartDate DATE NOT NULL,
    WeeksDuration INTEGER NOT NULL,
    UNIQUE (Year_ID, Name)
);

CREATE INDEX idx_semesters_startdate ON Semesters(StartDate);

CREATE TABLE IF NOT EXISTS Subjects (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS CategoriesTyps (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Categories (
    ID SERIAL PRIMARY KEY,
    Type_ID INTEGER NOT NULL REFERENCES CategoriesTyps(ID),
    Name VARCHAR(255) NOT NULL UNIQUE
);

CREATE INDEX idx_categories_type_id ON Categories(Type_ID);

CREATE TABLE IF NOT EXISTS Curriculum (
    ID SERIAL PRIMARY KEY,
    Subject_id INTEGER NOT NULL REFERENCES Subjects(ID),
    Description TEXT NOT NULL
);

CREATE INDEX idx_curriculum_subject_id ON Curriculum(Subject_ID);

CREATE TABLE IF NOT EXISTS Exams (
    ID SERIAL PRIMARY KEY,
    Subject_ID INTEGER NOT NULL REFERENCES Subjects(ID),
    Category_ID INTEGER NOT NULL REFERENCES Categories(ID),
    Description TEXT NOT NULL,
    Date DATE NOT NULL
);

CREATE INDEX idx_exams_subject_id ON Exams(Subject_ID);
CREATE INDEX idx_exams_date ON Exams(Date);

CREATE TABLE IF NOT EXISTS Holidays (
    ID SERIAL PRIMARY KEY,
    Category_ID INTEGER NOT NULL REFERENCES Categories(ID),
    Date DATE NOT NULL
);

CREATE INDEX idx_holidays_date ON Holidays(Date);

CREATE TABLE IF NOT EXISTS AcademicCalendar (
    ID SERIAL PRIMARY KEY,
    Semester_ID INTEGER NOT NULL REFERENCES Semesters(ID),
    DayNumber INTEGER NOT NULL,
    Subject_ID INTEGER REFERENCES Subjects(ID),
    Curriculum_ID INTEGER REFERENCES Curriculum(ID),
    Exam_ID INTEGER REFERENCES Exams(ID),
    Holiday_ID INTEGER REFERENCES Holidays(ID),
    Description TEXT,
    UNIQUE (Semester_ID, DayNumber)
);

CREATE INDEX idx_academiccalendar_semester_day ON AcademicCalendar(Semester_ID, DayNumber);

CREATE TABLE IF NOT EXISTS CurriculumSchedule (
    ID SERIAL PRIMARY KEY,
    Curriculum_ID INTEGER NOT NULL REFERENCES Curriculum(ID),
    Subject_ID INTEGER NOT NULL REFERENCES Subjects(ID),
    Semester_ID INTEGER NOT NULL REFERENCES Semesters(ID),
    DayOffset INTEGER NOT NULL,
    Description TEXT NOT NULL,
    UNIQUE (Curriculum_ID, Subject_ID, Semester_ID, DayOffset)
);

CREATE INDEX idx_curriculumschedule_semester_subject ON CurriculumSchedule(Semester_ID, Subject_ID);
