package gradebook

import (
	"errors"
	"log/slog"
	"time"
	"map/Student"
)

var (
	ErrGrade        = errors.New("оценка может быть в диапазоне от 2 до 5")
	ErrFoundStudent = errors.New("студент не найден в журнале")
	ErrStudent      = errors.New("аккаунт студента неактивен")
	ErrStudentExist = errors.New("студент с заданным ID уже есть в системе")
)

type GradeBook struct {
	Students   map[student.UUID]*student.Student
	GradeScale [4]string
	logger     *slog.Logger
}

func NewGradeBook(logger *slog.Logger) *GradeBook {
	return &GradeBook{
		Students:   make(map[student.UUID]*student.Student),
		GradeScale: [4]string{"Неуд", "Удов", "Хор", "Отл"},
		logger:     logger,
	}
}


func (g *GradeBook) AddStudent(name string) (student.UUID, error) {
	s, err := student.NewStudent(name)
	if err != nil {
		return student.UUID{}, err
	}

	if _, exists := g.Students[s.ID]; exists {
		return student.UUID{}, ErrStudentExist
	}

	g.Students[s.ID] = s

	if g.logger != nil {
		g.logger.Info("Студент добавлен в журнал",
			slog.String("student_id", s.ID.String()),
			slog.String("name", s.FullName),
		)
	}

	return s.ID, nil
}

func (g *GradeBook) AddGrade(id student.UUID, grade uint8) error {
	if grade < 2 || grade > 5 {
		return ErrGrade
	}

	found, ok := g.Students[id]
	if !ok {
		return ErrFoundStudent
	}

	if !found.IsActive() {
		return ErrStudent
	}

	record := student.GradeRecord{
		Value: grade,
		Date:  time.Now().UTC(),
	}

	found.Grades = append(found.Grades, record)
	found.AVG = found.Average()

	if g.logger != nil {
		g.logger.Info("Оценка сохранена в журнале",
			slog.String("student_id", found.ID.String()),
			slog.Uint64("grade", uint64(grade)),
			slog.Time("created", record.Date),
		)
	}

	return nil
}