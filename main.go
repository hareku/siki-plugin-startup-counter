package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	if runErr := run(); runErr != nil {
		log.Fatal(runErr)
	}
}

func run() error {
	logFileName, err := getLogFileName()
	if err != nil {
		return fmt.Errorf("get log file name: %w", err)
	}

	if err := record(logFileName); err != nil {
		return fmt.Errorf("record: %w", err)
	}
	cnt, err := countInThePastDay(logFileName)
	if err != nil {
		return fmt.Errorf("count in the past day: %w", err)
	}

	return beeep.Notify(fmt.Sprintf("%d times you launched", cnt), "Message from Siki plugins", "")
}

func getLogFileName() (string, error) {
	executable, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("get executable path: %w", err)
	}
	dir := filepath.Dir(executable)
	return filepath.Join(dir, "counter.log"), nil
}

func record(logFileName string) error {
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, time.Now().Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("write to log file: %w", err)
	}

	return nil
}

func countInThePastDay(logFileName string) (int, error) {
	f, err := os.Open(logFileName)
	if err != nil {
		return 0, fmt.Errorf("open log file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var count int
	for scanner.Scan() {
		txt := scanner.Text()
		dt, err := time.Parse(time.RFC3339, txt)
		if err != nil {
			return 0, fmt.Errorf("parse time %q: %w", txt, err)
		}
		if time.Since(dt) <= 24*time.Hour {
			count++
		}
	}

	return count, nil
}