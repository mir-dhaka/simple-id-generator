package main

import (
	"fmt"
	"strconv"
)

func getId(seed int, pre string) func() string {
	i := seed
	return func() string {
		i += 1
		return pre + "-" + strconv.Itoa(i)
	}
}

func main() {
	/* employeeId is now a function with i as 0 */
	employeeId := getId(0, "E")

	/* invoke employeeId to increase id by 1 and return it with added prefix */
	fmt.Println(employeeId())
	fmt.Println(employeeId())
	fmt.Println(employeeId())

	/* create id sequence for attendance*/
	attendanceId1 := getId(10, "A")
	
	/*some other calls*/
	fmt.Println(attendanceId1())
	fmt.Println(attendanceId1())
	fmt.Println(employeeId())
	fmt.Println(attendanceId1())
}
