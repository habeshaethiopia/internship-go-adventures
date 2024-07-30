package main

import (
	"fmt"
	"strings"
	 "bufio"
	 "os"
	"github.com/fatih/color"
)


func Grade(marks float64) string {
	if marks >= 90 {
		return "A+"
		} else if marks >= 80 {
			return "A"
	} else if marks >= 70 {
		return "B"
	} else if marks >= 60 {
		return "C"
	} else if marks >= 50 {
		return "D"
		} else {
		return "F"
	}
}

func main() {
	again:
	print("\033[H\033[2J")
	boldGreen := color.New(color.FgGreen, color.Bold)
	cyan:=color.New(color.FgCyan)
	boldGreen.Print("\n \t\tWelcome to the 🎓 Grade Calculator 🎓\n\n")
	// fmt.Println("welcome to the Grade Calculator")
	e1:
	cyan.Print("Enter the number of subjects: ")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n <= 0 {
		color.Red("Error in input")
		goto e1
	}

	total := 0.0
	marks := make(map[string]float64)
	cyan.Println("Enter the subjects and marks separated by space.")
	for i := 0; i < n; i++ {
		var subject string
		var mark float64
		e2:
		cyan.Printf("Enter the subject %d and mark: ", i+1)
		_, err := fmt.Scanln(&subject, &mark)
		if err != nil {
			color.Red("Error in input")
			read:=bufio.NewReader(os.Stdin)
			read.ReadString('\n')
			goto e2
		}
		if mark < 0 || mark > 100 {
			color.Red("Invalid marks")
			i--
			continue
		}
		marks[subject] = mark
		total += mark
	}
	percentage := total / float64(n)
	for sub, mar := range marks {

		color.Green("Subject: %s \t Grade: %s\n", sub, Grade(mar))
	}
	color.Blue("Total Average: %.2f\n", percentage)
	color.Yellow("Do you want to calculate again? (y/n)")
	var a string
	fmt.Scanln(&a)
	a = strings.ToLower(a)
	if a == "y" || a == "yes" {
		goto again
	}
}
