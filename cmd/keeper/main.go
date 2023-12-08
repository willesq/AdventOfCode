package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type DayMeta struct {
	Year string
	Day  string
}

func main() {
	curDate := time.Now()
	year := flag.Int("year", curDate.Year(), fmt.Sprintf("target year"))
	day := flag.Int("day", curDate.Day(), fmt.Sprintf("target day"))

	flag.Parse()
	if *day < 1 || *day > 25 {
		log.Printf("Cannot make target folder. %s", *day)
		os.Exit(1)
	}
	dayStr := strconv.Itoa(*day)
	if len(dayStr) == 1 {
		dayStr = fmt.Sprintf("0%d", *day)
	}
	dayM := DayMeta{Year: strconv.Itoa(*year), Day: dayStr}
	// all the file paths
	targetPath := fmt.Sprintf("cmd/%s/%s", dayM.Year, dayM.Day)
	inputFile := strings.Join([]string{targetPath, "input.txt"}, "/")
	dayGoFile := strings.Join([]string{targetPath, "internal", "day.go"}, "/")
	mainGoFile := strings.Join([]string{targetPath, "main.go"}, "/")
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		// Directory does not exist, so create it
		err := os.MkdirAll(targetPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v\n", err)
		} else {
			fmt.Printf("Directory %s created successfully.\n", targetPath)
		}
	} else {
		fmt.Printf("Directory %s already exists.\n", targetPath)
	}

	// make the input file (empty) if not exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		file, err := os.Create(inputFile)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
		} else {
			_ = file.Close()
		}
	} else {
		fmt.Print("Input file already exists.\n")
	}

	// make go day file
	if _, err := os.Stat(dayGoFile); os.IsNotExist(err) {
		err := os.MkdirAll(targetPath+"/internal", os.ModePerm)
		file, err := os.Create(dayGoFile)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
		} else {
			sourceFile, err := os.Open("./templates/day.tmpl")
			if err != nil {
				log.Fatal(err)
			}
			defer sourceFile.Close()
			_, err = io.Copy(file, sourceFile)
			_ = file.Close()
		}
	} else {
		fmt.Print("Day Go file already exists.\n")
	}

	// make main.go day file
	tmpl, err := template.ParseFiles("./templates/main.tmpl")
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(mainGoFile); os.IsNotExist(err) {
		file, err := os.Create(mainGoFile)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
		} else {
			writer := io.Writer(file)
			err = tmpl.Execute(writer, dayM)
			_ = file.Close()
		}
	} else {
		fmt.Print("Main Go file already exists.\n")
	}

}

//func ParseTemplates(path string) *template.Template {
//	templ := template.New("")
//	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
//		if strings.Contains(path, ".html") {
//			_, err = templ.ParseFiles(path)
//			if err != nil {
//				log.Println(err)
//			}
//		}
//
//		return err
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	return templ
//}
