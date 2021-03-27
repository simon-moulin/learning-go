package main

import "fmt"

func main() {
	printInfoNoParam()
	printInfoParam("Simon", 20, "sm@gmail.com")

	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, 64)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, -8)

	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v (len=%d, cap=%d)", letters, len(letters), cap(letters))

}

func printInfoNoParam() {
	fmt.Printf("Name:%s, age=%d, email=%s \n", "Bob", 10, "sm@golang.org")
}

func printInfoParam(name string, age int, email string) {
	fmt.Printf("Name:%s, age=%d, email=%s \n", name, age, email)
}

func avg(x float32, y float32) float32 {
	return (x + y) / 2
}

func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return sum
}
