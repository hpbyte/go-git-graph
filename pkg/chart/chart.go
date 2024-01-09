package chart

import (
	"fmt"
	"log"
	"time"
)

const layout = "2006-01-02"
const numOfDaysInWeek = 7
const numOfweeksInYear = 53

const (
	blankColor       = "\033[48;5;238m"
	dimGreenColor    = "\033[48;5;22m"
	midGreenColor    = "\033[48;5;28m"
	brightGreenColor = "\033[48;5;34m"
	resetColor       = "\033[0m"
	highlightColor   = "\033[48;5;226m"
)

type Chart interface {
	Render()
}

type ContributionChart struct {
	Data map[string]int
	Year int
	grid [numOfDaysInWeek][numOfweeksInYear]int // 7 (days) x 53 (weeks) grid 0s for a year with 0s as default
}

func (cc *ContributionChart) Render() {
	cc.parse()

	fmt.Printf("\n\n\n")

	printMonths()
	for i, row := range cc.grid {
		printDayCol(i)
		for _, col := range row {
			printCell(col, false)
		}
		fmt.Println()
	}
}

func (cc *ContributionChart) parse() {
	for commitDate, count := range cc.Data {
		parsed, err := time.Parse(layout, commitDate)
		if err != nil {
			log.Fatalf("[Err]: parsing date error for: %s", commitDate)
		}

		year, weekOfYear := parsed.ISOWeek()
		dayOfWeek := parsed.Weekday()

		if year == cc.Year {
			cc.grid[dayOfWeek][weekOfYear-1] = count
		}
	}
}

func printMonths() {
	var months [53]string
	offset := 0
	for month := time.January; month <= time.December; month++ {
		months[offset] = month.String()[:3]
		offset += 4
	}

	fmt.Print("     ")
	for _, month := range months {
		fmt.Printf("%s ", month)
	}
	fmt.Println()
}

func printDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}

	fmt.Printf(out)
}

func printCell(count int, today bool) {
	var colorCode string

	switch {
	case count == 0:
		colorCode = blankColor
	case count <= 5:
		colorCode = dimGreenColor
	case count <= 10:
		colorCode = midGreenColor
	default:
		colorCode = brightGreenColor
	}

	if today {
		colorCode = highlightColor
	}

	fmt.Printf("%s  %s", colorCode, resetColor)
}
