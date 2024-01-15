package chart

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const layout = "2006-01-02"
const numOfDaysInWeek = 7
const numOfweeksInYear = 53
const numOfCharsPerCell = 2
const numOfCharsPerMonthCell = 3

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

	printMonths(cc.Year)
	for i, row := range cc.grid {
		printDayCol(i)
		for _, col := range row {
			printCell(col, false)
		}
		fmt.Println()
	}
}

func (cc *ContributionChart) parse() {
	startOfYearOffset := numOfDaysInWeek - time.Date(cc.Year, 1, 1, 0, 0, 0, 0, time.UTC).Weekday()

	for commitDate, count := range cc.Data {
		parsed, err := time.Parse(layout, commitDate)
		if err != nil {
			log.Fatalf("[Err]: parsing date error for: %s", commitDate)
		}

		year, weekOfYear := parsed.ISOWeek()
		dayOfWeek := parsed.Weekday()

		// handle offset when startOfyear is not on Sunday 0
		// e.g.
		// - if the year starts on Monday 1, offset is 7 - 1 = 6
		// - if commitDate is on Sunday 0, Week 1, it should be on 2nd Week due to offset
		// - 7 - 0 > offset (6) => we move it to the next week
		// pls be aware that week starts from 1 instead of 0, hence the condition is below
		weekIndex := weekOfYear
		if numOfDaysInWeek-dayOfWeek <= startOfYearOffset {
			weekIndex -= 1
		}

		if year == cc.Year {
			cc.grid[dayOfWeek][weekIndex] = count
		}
	}
}

func findNumOfWeeksInMonth(year int, month time.Month, prevWeek *int) int {
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	log.Printf("prevWeek %d", *prevWeek)
	_, firstWeekOfMonth := firstDayOfMonth.ISOWeek()
	_, lastWeekOfMonth := lastDayOfMonth.ISOWeek()

	// adjustment for the case when the current year starts in the same week as the prev year ends
	if month == time.January && firstWeekOfMonth > 1 {
		firstWeekOfMonth = 1
	} else if month == time.December && lastDayOfMonth.Weekday() < time.Thursday {
		// adjustment for the case when the next year starts in the same week as the current year ends
		// otherwise lastWeekOfMonth would become 1
		lastWeekOfMonth = numOfweeksInYear + (lastDayOfMonth.Day()-28)/numOfDaysInWeek
	}
	log.Printf("firstDayOfMonth: %s lastDayOfMonth: %s firstWeekOfMonth: %d, lastWeekOfMonth: %d\n", firstDayOfMonth, lastDayOfMonth, firstWeekOfMonth, lastWeekOfMonth)

	numOfWeeksInMonth := lastWeekOfMonth - firstWeekOfMonth + 1

	if *prevWeek == firstWeekOfMonth {
		log.Println("firstweek and prevWeek is the same: ", firstDayOfMonth)
		numOfWeeksInMonth -= 1
	}

	*prevWeek = lastWeekOfMonth

	return numOfWeeksInMonth
}

func printMonths(year int) {
	var builder strings.Builder

	builder.WriteString("     ")

	prevWeek := 0

	for month := time.January; month <= time.December; month++ {
		monthName := month.String()[:3]
		numOfWeeks := findNumOfWeeksInMonth(year, month, &prevWeek)
		log.Printf("month: %s, numOfWeeks: %d\n", monthName, numOfWeeks)
		len := (numOfWeeks * numOfCharsPerCell) - numOfCharsPerMonthCell

		formatted := fmt.Sprintf("%-3s%*s", monthName, len, "")
		builder.WriteString(formatted)
	}

	fmt.Println(builder.String())
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
