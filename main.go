package main

import (
	"fmt"
	"sort"
	"strings"
)

var months [12]string = [12]string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}

func str_rep(str string, not int) string {
	// This function is meant as an alias
	// to strings.Repeat() for ease of use.

	return strings.Repeat(str, not)
}

func strToInt(str string) (bool, int) {
	// - Function to convert string input from
	//   getInput() to integer datatype.
	// - Uses fmt.Sscanf() to read bytes from
	//   variable str to variable num.
	// - Returns (False, 0) if str is not a
	//   number and (True, num) otherwise.

	var num int = 0
	var isint bool = true

	_, e := fmt.Sscanf(str, "%d\n", &num)

	if e != nil {
		isint = false
	}

	return isint, num
}

func printInfo() {

	fmt.Println("1. Generate the annual calendar of an year.")
	fmt.Println("2. Generate the calendar of a month.")
	fmt.Println("3. Find the day of a date.")
	fmt.Println("4. Find the number of days between two dates.")
	fmt.Println("5. Find the Zodiac sign of a date.")
	fmt.Println("6. Exit the program.")
	fmt.Println("Press <Enter> to enter input stored in history.")
	fmt.Println("Enter <i> and press <Enter> to list possible commands.")
	fmt.Println("Type 'exit' at any promp and press <Enter> to exit the program.")

}

func printMonths() {

	const hline string = "\n+----------+----------+"
	var rep string

	for i := 1; i <= 6; i++ {
		rep = "  "

		fmt.Println(hline)
		fmt.Printf("|")

		for c, j := range [2]int{i, i + 6} {
			if (c == 1) && (j > 9) {
				rep = " "
			}

			fmt.Printf(" %s   %d%s|", months[j-1], j, rep)
		}
	}

	fmt.Println(hline)

	return
}

func clearBuffer() {
	// - Function to clear input buffer stream
	//   by repeatedly calling fmt.Scanln()
	//   until it returns no errors.
	// - This function is invoked only if input
	//   contains more than one space separated
	//   value.

	var cont bool = true

	for cont {
		_, error := fmt.Scanln()
		if error == nil {
			cont = false
		}
	}

	return
}

/* function under development */

func getInput(prompt string) int {
	// Function gets the input value and
	// returns an integer. Handles input
	// data type errors.

	var input string
	var isint bool
	var num int

	for true {
		fmt.Printf("Enter the %s: ", prompt)
		_, error := fmt.Scanln(&input)

		if error != nil {
			if error.Error() != "unexpected newline" {
				clearBuffer()
			}
		}

		if input == "exit" {
			fmt.Println("[Abort]")
			return -1
		} else {
			isint, num = strToInt(input)
			if isint {
				return num
			} else {
				fmt.Printf("[Error] : '%s' is not a valid input!\n", input)
				input = ""
			}
		}

		fmt.Println()
	}

	return -1
}

func getYear(year *int) {

	for true {
		*year = getInput("year")

		if *year < 1000 || *year > 9999 {
			fmt.Println("[Error] : Input not in range 1000 - 9999")
		} else {
			return
		}
	}
}

func getMonth(year, month *int) {

	getYear(year)

	for true {
		*month = getInput("month")

		if *month < 1 || *month > 12 {
			fmt.Println("[Error] : Input not in range 1 - 12")
		} else {
			*month -= 1
			return
		}
	}
}

func getDay(year, month, day *int) {

	var nod int

	getMonth(year, month)

	nod = findNOD(*month, ifLeap(*year))

	for true {
		*day = getInput("date")

		if *day < 1 || *day > nod {
			fmt.Printf("[Error] : Input not in range 1 - %d\n", nod)
		} else {
			return
		}
	}
}

func getChoice() int {

	var choice int

	for true {
		choice = getInput("choice")

		if choice < 1 || choice > 6 {
			fmt.Println("[Error] : Input not in range 1 - 6")
		} else {
			return choice
		}
	}

	return -1
}

func ifLeap(year int) bool {

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}

func firstDay(year int) int {

	var century int = year / 100
	var century_offset int = 4
	var day_offset int = 1
	var month_offset int = 0
	var day_no int = 0
	var year_offset int = (((year % 100) * 5) / 4) % 7

	for century_offset <= century {
		century_offset += 4
	}

	century_offset -= (century + 1)
	century_offset *= 2

	if ifLeap(year) {
		year_offset -= 1
	}

	day_no = (century_offset + year_offset + month_offset + day_offset) % 7

	if day_no == 0 {
		day_no = 6
	} else {
		day_no--
	}

	return day_no
}

func findNOD(month int, leap_year bool) int {
	// Function to find the number of days
	// in a given month of a given year.

	var nod int

	if (month == 3) || (month == 5) || (month == 8) || (month == 10) {
		nod = 30
	} else if month == 1 {
		if leap_year {
			nod = 29
		} else {
			nod = 28
		}
	} else {
		nod = 31
	}

	return nod
}

func genCal(year, m int) map[string]map[int][7]int {
	// The function returns calendar for entire year
	// if second parameter m = -1. Else returns the
	// calendar for the month specified by m.

	var day_no int = firstDay(year)
	cal := make(map[string]map[int][7]int)

	var week_no int = 1
	var date_lst [7]int
	var leap_year bool
	var nod int

	leap_year = ifLeap(year)

	var skipCal func(i int)
	var monthCal func(i int, month string)

	skipCal = func(i int) {

		nod = findNOD(i, leap_year)

		for j := 1; j <= nod; j++ {
			if day_no == 6 {
				day_no = 0
				week_no++
			} else if j == nod {
				day_no++
			} else {
				day_no++
			}
		}
	}

	monthCal = func(i int, month string) {

		cal[month] = map[int][7]int{}
		nod = findNOD(i, leap_year)

		for j := 1; j <= nod; j++ {
			date_lst[day_no] = j

			if day_no == 6 {
				cal[month][week_no] = date_lst
				date_lst = [7]int{}
				day_no = 0
				week_no++
			} else if j == nod {
				cal[month][week_no] = date_lst
				date_lst = [7]int{}
				day_no++
			} else {
				day_no++
			}
		}
	}

	for i, month := range months {
		if m == -1 || i == m {
			monthCal(i, month)
		} else {
			skipCal(i)
		}
	}

	return cal
}

func printWeek(k, lm, m int, keys []int, cal map[string]map[int][7]int) {

	var day int

	if k < lm {
		if keys[k] < 10 {
			fmt.Printf("|%d ", keys[k])
		} else {
			fmt.Printf("|%d", keys[k])
		}

		for j := 0; j < 7; j++ {
			day = cal[months[m]][keys[k]][j]

			if day == 0 {
				fmt.Printf("    ")
			} else if day < 10 {
				fmt.Printf("   %d", day)
			} else {
				fmt.Printf("  %d", day)
			}
		}
	} else {
		fmt.Printf("|%s", str_rep(" ", 30))
	}

	return
}

func calOfYear() {

	var year int
	var cal map[string]map[int][7]int

	getYear(&year)

	cal = genCal(year, -1)

	const hline string = "+------------------------------+"
	const dline string = "|Wk  Mo  Tu  We  Th  Fr  Sa  Su|"

	var now int /* short for no. of weeks */
	var m1 int  /* first month */
	var m2 int  /* second month */
	var lm1 int /* length of m1 */
	var lm2 int /* length of m2 */

	for i := 0; i < 6; i++ {
		m1 = 2 * i
		m2 = m1 + 1

		fmt.Println(str_rep(" ", 14), months[m1], str_rep(" ", 30), months[m2])
		fmt.Println(hline, " ", hline)
		fmt.Println(dline, " ", dline)
		fmt.Println(hline, " ", hline)

		lm1 = len(cal[months[m1]])
		lm2 = len(cal[months[m2]])

		var keys_m1 []int = make([]int, 0, lm1)
		var keys_m2 []int = make([]int, 0, lm2)

		for k := range cal[months[m1]] {
			keys_m1 = append(keys_m1, k)
		}

		for k := range cal[months[m2]] {
			keys_m2 = append(keys_m2, k)
		}

		sort.Ints(keys_m1)
		sort.Ints(keys_m2)

		if lm1 > lm2 {
			now = lm1
		} else {
			now = lm2
		}

		for k := 0; k < now; k++ {
			printWeek(k, lm1, m1, keys_m1, cal)
			fmt.Printf("|   ")

			printWeek(k, lm2, m2, keys_m2, cal)
			fmt.Println("|")
		}

		fmt.Println(hline, " ", hline)
	}

	return
}

func calOfMonth() {

	var year, month int
	var cal map[string]map[int][7]int

	const hline string = "+------------------------------+"
	const dline string = "|Wk  Mo  Tu  We  Th  Fr  Sa  Su|"

	getMonth(&year, &month)

	cal = genCal(year, month)

	var length int = len(cal[months[month]])
	var keys []int = make([]int, 0, length)

	for k := range cal[months[month]] {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println(str_rep(" ", 14), months[month])

	fmt.Println(hline)
	fmt.Println(dline)
	fmt.Println(hline)

	for key := range keys {
		printWeek(key, length, month, keys, cal)
		fmt.Println("|")
	}

	fmt.Println(hline)
}

func main() {

	//var day int

	/* testing function strToInt */
	b, n := strToInt("3468")
	fmt.Println(b, n)
	/* testing function strToInt */

	printMonths()
	fmt.Println()
	calOfMonth()

}
