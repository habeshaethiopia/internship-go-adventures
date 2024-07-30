package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func helper(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+( [a-zA-Z]+)*$`)
	return re.MatchString(s)

}

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
	cyan := color.New(color.FgCyan)
	boldcyan := color.New(color.FgCyan, color.Bold)
	boldGreen.Print("\n \t\tWelcome to the 🎓 Grade Calculator 🎓\n\n")
	// fmt.Println("welcome to the Grade Calculator")
e1:
	cyan.Print("Enter Your Name:")
	var name string
	read := bufio.NewReader(os.Stdin)
	name, _ = read.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" || !helper(name) {
		color.Red("Error in input name must be alphabets only")
		goto e1
	}
e3:

	cyan.Print("Enter the number of subjects: ")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n <= 0 {
		color.Red("Error in input or invalid number of subjects")
		read := bufio.NewReader(os.Stdin)
		read.ReadString('\n')
		goto e3
	}

	total := 0.0
	marks := make(map[string]float64)
	boldcyan.Println("Enter the subjects and marks separated by space.")
	for i := 0; i < n; i++ {
		var subject string
		var mark float64
	e2:
		cyan.Printf("Enter the subject %d and mark: ", i+1)
		_, err := fmt.Scanln(&subject, &mark)
		if err != nil {
			color.Red("Error in input")
			read := bufio.NewReader(os.Stdin)
			read.ReadString('\n')
			fmt.Println("")
			goto e2
		}
		if mark < 0 || mark > 100 {
			color.Red("Invalid marks marks should be between 0 and 100")
			i--
			continue
		}
		marks[subject] = mark
		total += mark
	}
	percentage := total / float64(n)
	boldGreen.Printf("Name: %s\n", name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Subject", "Marks", "Grade"})
	for sub, mar := range marks {
		table.Append([]string{sub, fmt.Sprintf("%.2f", mar), Grade(mar)})
	}

	table.Render()

	color.Blue("Total Average: %.2f\n", percentage)
	color.Yellow("Do you want to calculate again? (y/n)")
	var a string
	fmt.Scanln(&a)
	a = strings.ToLower(a)
	if a == "y" || a == "yes" {
		goto again
	} else {
		print("\033[H\033[2J")

	}

}
