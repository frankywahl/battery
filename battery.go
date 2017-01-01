package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	now := time.Now()
	percentage := "101%"

	//WriteHeaders("results.csv")
	cmd := exec.Command("pmset", "-g", "batt")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	matched := regexp.MustCompile("([0-9]*)%").FindString(out.String())

	if matched != percentage {
		percentage = matched

		AppendRestults(now, time.Now(), percentage)
	}
}

func WriteHeaders(filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	data := []string{"Percentage", "Elapsed Time", "Time", "Time: Human readable"}
	err = writer.Write(data)
	if err != nil {
		log.Println("Error: ", err)
	}

}

func AppendResults(filename, startTime, currentTime, percentage) {
	file, err := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
}
