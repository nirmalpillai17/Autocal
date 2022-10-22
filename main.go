package main

import (
	"fmt"
)

var months [12]string = [12]string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}

func qsort(arr []int) {

	var high int = len(arr) - 1

	var pivot int
	var pi int

	var if_sorted func() bool
	var partition func(lo, hi int) int
	var init_sort func(lo, hi int)

	if_sorted = func() bool {

		for i := 0; i < high; i++ {
			if arr[i] > arr[i+1] {
				return false
			}
		}

		return true
	}

	partition = func(lo, hi int) int {

		pivot = arr[lo]

		for true {
			for arr[lo] < pivot {
				lo++
			}

			for arr[hi] > pivot {
				hi--
			}

			if lo >= hi {
				return hi
			}

			arr[lo], arr[hi] = arr[hi], arr[lo]
		}

		return 0
	}

	init_sort = func(lo, hi int) {

		if lo < hi {
			pi = partition(lo, hi)

			init_sort(lo, pi)
			init_sort(pi+1, hi)
		}

		return
	}

	if !if_sorted() {
		init_sort(0, high)
	}

	return
}

func str_rep(str string, not int) string {

	var str_out string

	for i := 1; i <= not; i++ {
		str_out += str
	}

	return str_out
}

/* function to convert string to integer */
/* under development */

func stoi(str string) (bool, int) {

	var num int
	var is_num bool = true

	_, e := fmt.Sscan(str, &num)

	if e != nil {
		is_num = false
	}

	return is_num, num
}

/* under development */

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

	var cont bool = true

	for cont {
		_, error := fmt.Scanln()
		if error == nil {
			cont = false
		}
	}

	return
}

func get_input() int {

	var year int

	for true {
		fmt.Printf("Enter the year: ")
		_, error := fmt.Scanln(&year)

		/* clear input buffer stream if any */

		if error != nil {
			if error.Error() != "unexpected newline" {
				clr_buf()
			}
			fmt.Println("[Error] : Input not an year")
		} else if year < 1000 || year > 9999 {
			fmt.Println("[Error] : Input not in range 1000 - 9999")
		} else {
			break
		}

		fmt.Println()
	}

	return year
}

func first_day(year int) (int, bool) {

	var century int = year / 100
	var century_offset int = 4
	var day_offset int = 1
	var month_offset int = 0
	var year_offset int = (((year % 100) * 5) / 4) % 7

	var leap_year bool
	var day_no int

	for century_offset <= century {
		century_offset += 4
	}

	century_offset -= (century + 1)
	century_offset *= 2

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		leap_year = true
		year_offset -= 1
	} else {
		leap_year = false
	}

	day_no = (century_offset + year_offset + month_offset + day_offset) % 7

	if day_no == 0 {
		day_no = 6
	} else {
		day_no--
	}

	return day_no, leap_year
}

func gen_cal(day_no int, leap_year bool) map[string]map[int][7]int {

	cal := make(map[string]map[int][7]int)

	var week_no int = 1
	var date_lst [7]int
	var nod int

	for i, month := range months {
		cal[month] = map[int][7]int{}

		if (i == 3) || (i == 5) || (i == 8) || (i == 10) {
			nod = 30
		} else if i == 1 {
			if leap_year {
				nod = 29
			} else {
				nod = 28
			}
		} else {
			nod = 31
		}

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

		qsort(keys_m1)
		qsort(keys_m2)

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
	var leap_year bool
	var cal map[string]map[int][7]int

	/* testing function stoi */
	b, n := stoi("abcd")
	fmt.Println(b, n)
	/* testing function stoi */
	
	year = get_input()
	print_months()
	fmt.Println()
	day_no, leap_year = first_day(year)
	cal = gen_cal(day_no, leap_year)
	print_cal(cal)

}
