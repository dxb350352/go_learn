package main
import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6}
	var temp []int = []int{}
	for i := len(arr) - 3; i < len(arr); i++ {
		temp = append(temp, arr[i])
	}
	fmt.Println(arr)
	arr = append(arr[:len(arr) - 3], 0)
	arr = append(arr, temp...)
	fmt.Println(arr)
}
