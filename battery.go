package main

import (
	"bytes"
	"encoding/csv"
	"github.com/frankywahl/battery/datapoint"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type Observer interface {
	Update() bool
}

type Observable interface {
	AddObserver(o Observer)
	Observers []Observers
}

func (observable *Observable) NotifyObserver() {
	for _, observer := range observable.observers {
		observer.Update(observable)
	}
}

func main() {
	now := time.Now()
	percentage := "101%"

	datapoint := datapoint.New()
	log.Println(datapoint.Percentage)

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

		//AppendRestults(now, time.Now(), percentage)
	}
	log.Println(now)
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

func AppendResults(filename string, startTime time.Time, currentTime time.Time, percentage int) {
	file, err := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err != nil {
		log.Println("Error: ", err)
	}
}
