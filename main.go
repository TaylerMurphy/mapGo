package main

import(
	"map/GradeBook"
	"fmt"
	"log/slog"
	"os"
)

func main(){

	handle := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, handle)
	logger := slog.New(jsonHandler)

	gradeBook := GradeBook.NewGradeBook(logger)

	if errStudent1 := gradeBook.AddStudent(1, "Petrov Petr"); errStudent1 != nil{
		fmt.Printf("Error: %v", errStudent1)
	}

	if errStudent2 := gradeBook.AddStudent(2, "Petrov Petr"); errStudent2 != nil{
		fmt.Printf("Error: %v", errStudent2)
	}

	if errGrade1 := gradeBook.AddGrade(1, 4); errGrade1 != nil{
			fmt.Printf("Error: %v", errGrade1)
	}

	if errGrade1 := gradeBook.AddGrade(1, 5); errGrade1 != nil{
			fmt.Printf("Error: %v", errGrade1)
	}

	if errGrade0 := gradeBook.AddGrade(0, 4); errGrade0 != nil{
			fmt.Printf("Error: %v", errGrade0)
	}

	fmt.Scanln()
}