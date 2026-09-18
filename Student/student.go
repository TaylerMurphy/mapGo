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
	ID         UUID
	FullName   string
	Grades     []GradeRecord
	EnterDate  time.Time
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
	id[6] = (id[6] & 0x0f) | 0x40
	id[8] = (id[8] & 0x3f) | 0x80
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
		Flags:      FlagActive,
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

func (s *Student) setFlag(f StatusFlags, v bool) {
	if v {
		s.Flags |= f
	} else {
		s.Flags &^= f
	}
}

func (s *Student) hasFlag(f StatusFlags) bool {
	return s.Flags&f != 0
}

func (s *Student) SetActive(v bool) {
	s.setFlag(FlagActive, v)
}

func (s *Student) SetHonors(v bool) {
	s.setFlag(FlagHonors, v)
}

func (s *Student) SetFinancialDebt(v bool) {
	s.setFlag(FlagAcademicDebt, v)
}

func (s *Student) SetAcademicDebt(v bool) {
	s.setFlag(FlagAcademicDebt, v)
}

func (s *Student) SetNonResident(v bool) {
	s.setFlag(FlagNonResident, v)
}

func (s *Student) IsActive() bool {
	return s.hasFlag(FlagActive)
}

func (s *Student) IsHonor() bool {
	return s.hasFlag(FlagHonors)
}

func (s *Student) IsFinancialDebt() bool {
	return s.hasFlag(FlagFinancialDebt)
}

func (s *Student) IsAcatemicDebt() bool {
	return s.hasFlag(FlagAcademicDebt)
}

func (s *Student) IsNonResident() bool {
	return s.hasFlag(FlagNonResident)
}
