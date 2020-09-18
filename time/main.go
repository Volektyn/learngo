// You have to perform the following steps to get the desired output:
// 1.
//  Use the time.Now() function to capture the current date in a variable.
// 2. Dissect the captured date to day, month, year, hour, minute, and seconds variables by
// converting them into strings.
// 3. Print out the concatenated variables in order.
// Once the script is complete, the output should appear as follows (note that this
// depends on when you run the code):
// 15:32:30 2019/10/17
// By the end of this activity, you should have learned how you can craft your
// custom time variables and use strconv.Itoa() to convert a number to a string and
// concatenate the result.
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	day := strconv.Itoa(int(now.Day()))
	month := strconv.Itoa(int(now.Month()))
	year := strconv.Itoa(int(now.Year()))
	hour := strconv.Itoa(int(now.Hour()))
	minute := strconv.Itoa(int(now.Minute()))
	seconds := strconv.Itoa(int(now.Second()))
	fmt.Println(hour + ":" + minute + ":" + seconds + " " + year + "/" + month + "/" + day)
}
