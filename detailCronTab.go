package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cronIntf interface {
	findValues(string, int, int) (string, error)
}
type cronDetails struct {
	minute     string
	hour       string
	dayOfMonth string
	month      string
	dayOfWeek  string
	command    string
}

func (cd *cronDetails) findValues(str string, start, end int) (string, error) {
	var retStr string
	var step int
	var err error
	ch := ""
	// fmt.Println("debug : ", str, "  is  : ", str, " ", start, " ", end)
	if strings.Contains(str, "*") {
		ch = "*"
		if str == "*" {
			step = 1
			goto lable
		}
		matched, err := regexp.MatchString("^\\*\\/[1-9]+$", str)
		if err != nil {
			return retStr, err
		}
		if matched {
			step, err = strconv.Atoi(strings.Replace(str, "*/", "", 1))
			if err != nil {
				return retStr, err
			}
		} else {
			return retStr, errors.New("invaid string format")
		}
	} else if strings.Contains(str, "-") {
		step = 1
		ch = "-"
		sp := strings.Split(str, "-")
		start, err = strconv.Atoi(sp[0])

		if err != nil {
			return retStr, err
		}
		end, err = strconv.Atoi(sp[1])
		if err != nil {
			return retStr, err
		}
	} else if strings.Contains(str, ",") {
		ch = ","
		sp := strings.Split(str, ",")
		start, err = strconv.Atoi(sp[0])
		if err != nil {
			return retStr, err
		}
		end, err = strconv.Atoi(sp[1])
		if err != nil {
			return retStr, err
		}
	} else {
		isdigit, err := regexp.MatchString(`^\d+$`, str)
		if err != nil {
			return retStr, err
		} else if !isdigit {
			return retStr, errors.New("invalid string format")
		} else {
			ch = "digit"
		}
	}

lable:
	switch ch {
	case "*", "-":
		for i := start; i <= end; i = i + step {
			retStr = retStr + strconv.Itoa(i) + " "
		}
	case ",":
		retStr = strconv.Itoa(start) + " " + strconv.Itoa(end)

	case "digit":
		retStr = str
	}

	return retStr, nil

}

func main() {
	var ci cronIntf
	if len(os.Args) != 2 {
		fmt.Println("Input request is not valid")
		return
	}

	inputStr := os.Args[1]

	items := strings.Split(inputStr, " ")
	if len(items) < 6 {
		fmt.Println("Cron format is not valid")
		return
	}

	cr1 := cronDetails{
		minute:     items[0],
		hour:       items[1],
		dayOfMonth: items[2],
		month:      items[3],
		dayOfWeek:  items[4],
		command:    strings.Join(items[5:], " "),
	}

	ci = &cr1
	minute, err := ci.findValues(cr1.minute, 0, 59)
	if err != nil {
		fmt.Println("Error Validating Minute: ", err)
	}
	hour, err := ci.findValues(cr1.hour, 0, 23)
	if err != nil {
		fmt.Println("Error Validating Hour: ", err)
	}
	dayOfMonth, err := ci.findValues(cr1.dayOfMonth, 1, 31)
	if err != nil {
		fmt.Println("Error Validating Day of Month: ", err)
	}
	month, err := ci.findValues(cr1.month, 1, 12)
	if err != nil {
		fmt.Println("Error Validating Month: ", err)
	}
	dayOfWeek, err := ci.findValues(cr1.dayOfWeek, 0, 6)
	if err != nil {
		fmt.Println("Error Validating Day of week: ", err)
	}

	fmt.Println("minute\t", minute)
	fmt.Println("hour\t", hour)
	fmt.Println("day of month\t", dayOfMonth)
	fmt.Println("month\t", month)
	fmt.Println("day of week\t", dayOfWeek)
	fmt.Println("command\t", cr1.command)

}
