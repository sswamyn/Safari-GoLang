package main

import "fmt"

func main() {
	nums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	someNums := nums[3:7]
	fmt.Printf("someNums value: %#v, its a %T at %p\n", someNums, someNums, &someNums)
	moreNums := nums[2:8]
	fmt.Printf("moreNums value: %#v, its a %T at %p\n", moreNums, moreNums, &moreNums)
	someNums[0] = 1000
	fmt.Printf("someNums value: %#v, its a %T at %p\n", someNums, someNums, &someNums)
	fmt.Printf("moreNums value: %#v, its a %T at %p\n", moreNums, moreNums, &moreNums)
	fmt.Printf("nums value: %#v, its a %T at %p\n", nums, nums, &nums)
	fmt.Printf("someNums cap is %d len is %d\n", cap(someNums), len(someNums))

	//sl3 := append(someNums, -1, -2, -3, -4, -5, -6)

	sl3 := make([]int, 0, 20)
	copy(sl3, someNums)
	sl3 = append(sl3, -1, -2, -3, -4, -5, -6)

	fmt.Printf("sl3 is %#v cap is %d len is %d\n", sl3, cap(sl3), len(sl3))
	fmt.Printf("nums value: %#v, its a %T at %p\n", nums, nums, &nums)
	fmt.Printf("someNums %#v, cap is %d len is %d\n", someNums, cap(someNums), len(someNums))

}
