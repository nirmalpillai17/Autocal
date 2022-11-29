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

func stoi(str string) (bool, int) {
	// - Function to convert string input from
	//   get_input() to integer datatype.
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

func print_info() {

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

func print_months() {

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

func clr_buf() {
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

func get_input(prompt string) int {
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
				clr_buf()
			}
		}

		if input == "exit" {
			fmt.Println("[Abort]")
			return -1
		} else {
			isint, num = stoi(input)
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

func get_year() int {

	var year int

	for true {
		year = get_input("year")

		if year < 1000 || year > 9999 {
			fmt.Println("[Error] : Input not in range 1000 - 9999")
		} else {
			return year
		}
	}

	return -1
}

func get_month() (int, int) {

	var year int
	var month int

	year = get_year()

	for true {
		month = get_input("month")

		if month < 1 || month > 12 {
			fmt.Println("[Error] : Input not in range 1 - 12")
		} else {
			return year, (month - 1)
		}
	}

	return -1, -1
}

func get_day() (int, int, int) {

	var year int
	var month int
	var day int
	var nod int

	year, month = get_month()

	nod = find_nod(month, if_leap(year))

	for true {
		day = get_input("day")

		if day < 1 || day > nod {
			fmt.Printf("[Error] : Input not in range 1 - %d\n", nod)
		} else {
			return year, month, day
		}
	}

	return -1, -1, -1
}

func if_leap(year int) bool {

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}

func first_day(year int) int {

	var century int = year / 100
	var century_offset int = 4
	var day_offset int = 1
	var month_offset int = 0
	var year_offset int = (((year % 100) * 5) / 4) % 7

	var day_no int

	for century_offset <= century {
		century_offset += 4
	}

	century_offset -= (century + 1)
	century_offset *= 2

	if if_leap(year) {
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

func find_nod(month int, leap_year bool) int {

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

func gen_cal(day_no int, year int) map[string]map[int][7]int {

	cal := make(map[string]map[int][7]int)

	var week_no int = 1
	var date_lst [7]int
	var leap_year bool
	var nod int

	leap_year = if_leap(year)

	for i, month := range months {
		cal[month] = map[int][7]int{}

		nod = find_nod(i, leap_year)

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

	return cal
}

func print_week(k, lm, m int, keys []int, cal map[string]map[int][7]int) {

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

func print_cal(cal map[string]map[int][7]int) {

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

		var keys_m1 []int = make([]int, 0, len(cal[months[m1]]))
		var keys_m2 []int = make([]int, 0, len(cal[months[m2]]))

		for k := range cal[months[m1]] {
			keys_m1 = append(keys_m1, k)
		}

		for k := range cal[months[m2]] {
			keys_m2 = append(keys_m2, k)
		}

		sort.Ints(keys_m1)
		sort.Ints(keys_m2)

		lm1 = len(cal[months[m1]])
		lm2 = len(cal[months[m2]])

		if lm1 > lm2 {
			now = lm1
		} else {
			now = lm2
		}

		for k := 0; k < now; k++ {
			print_week(k, lm1, m1, keys_m1, cal)
			fmt.Printf("|%s", "   ")

			print_week(k, lm2, m2, keys_m2, cal)
			fmt.Printf("|\n")
		}

		fmt.Println(hline, " ", hline)
	}

	return
}

func main() {

	var year int
	var day_no int
	var month int
	var day int
	var cal map[string]map[int][7]int

	/* testing function stoi */
	b, n := stoi("3468")
	fmt.Println(b, n)
	/* testing function stoi */

	year, month, day = get_day()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	if year == -1 {
		return
	}
	print_months()
	fmt.Println()
	day_no = first_day(year)
	cal = gen_cal(day_no, year)
	print_cal(cal)

}
