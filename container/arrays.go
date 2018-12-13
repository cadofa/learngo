package container

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 88
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println( arr3)
	//fmt.Println(grid)

	printArray(&arr3)
	fmt.Println(arr3)
}
