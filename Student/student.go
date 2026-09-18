package student

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

type UUID [16]byte

func (u UUID) String() string {
	return hex.EncodeToString(u[:])
}

type StatusFlags uint8

const (
	FlagActive StatusFlags = 1 << iota
	FlagHonors
	FlagFinancialDebt
	FlagAcademicDebt
	FlagNonResident
)

type GradeRecord struct {
	Value uint8
	Date  time.Time
}

type Student struct {
	ID         UUID          // id, по которому студент хранится в журнале
	FullName   string        // ФИО
	Grades     []GradeRecord // оценки с датами
	EnterDate  time.Time     // дата зачисления
	CourseYear uint8
	AVG        float64
	Flags      StatusFlags
}

var (
	ErrName             = errors.New("ФИО слишком короткое")
	ErrorGeneratingUUID = errors.New("ошибка генерации UUID для студента")
)

func GenerateUUID() (UUID, error) {
	var id UUID
	if _, err := rand.Read(id[:]); err != nil {
		return UUID{}, ErrorGeneratingUUID
	}
	id[6] = (id[6] & 0x0f) | 0x40 // version 4
	id[8] = (id[8] & 0x3f) | 0x80 // variant 10
	return id, nil
}

func NewStudent(name string) (*Student, error) {
	name = strings.TrimSpace(name)

	if utf8.RuneCountInString(name) < 3 {
		return nil, ErrName
	}

	id, err := GenerateUUID()
	if err != nil {
		return nil, err
	}

	return &Student{
		ID:         id,
		FullName:   name,
		Grades:     make([]GradeRecord, 0),
		AVG:        0,
		EnterDate:  time.Now().UTC(),
		CourseYear: 1,
		Flags:      FlagActive, // по умолчанию активен
	}, nil
}

func (s *Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	var sum uint
	for _, g := range s.Grades {
		sum += uint(g.Value)
	}

	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) IsActive() bool {
	return s.Flags&FlagActive != 0
}