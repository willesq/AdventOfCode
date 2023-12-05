package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func main() {
	time.Now().Year()
	year := flag.String("year", strconv.Itoa(time.Now().Year()), "target year")
	dayStr := strconv.Itoa(time.Now().Day())
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}
	day := flag.String("day", dayStr, "target day")

	rootDir := "cmd"
	dirPrefix := fmt.Sprintf("%s/%s/%s", rootDir, *year, *day)
	internalDir := fmt.Sprintf("%s/internal", dirPrefix)
	err := os.MkdirAll(internalDir, os.ModePerm)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}
	templatePrefix := rootDir + "/templates/"
	os.Link(templatePrefix+"main.go.tmp", dirPrefix+"/main.go")

}
