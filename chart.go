package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

const layout = "2006-01-02"

type Chart interface {
	Render()
}

type ContributionChart struct {
	Data map[string]int
	Year int
}

func (cc ContributionChart) Render() {
	// 7 (days) x 53 (weeks) grid 0s for a year with 0s as default
	var grid [7][53]int

	for commitDate, count := range cc.Data {
		parsed, err := time.Parse(layout, commitDate)
		if err != nil {
			log.Println("date parsing error for: ", commitDate)
			continue
		}

		year, weekOfYear := parsed.ISOWeek()
		dayOfWeek := parsed.Weekday()

		if year == cc.Year {
			grid[dayOfWeek][weekOfYear-1] = count
		}
	}

	for _, row := range grid {
		for _, col := range row {
			cc.colorizer(strconv.Itoa(col), ColorBlue)
		}
		fmt.Println()
	}
}

func (cc ContributionChart) colorizer(message string, color Color) {
	fmt.Print(string(color), message, string(ColorReset))
}
