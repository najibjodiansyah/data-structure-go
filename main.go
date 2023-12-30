package main

import (
	"fmt"

	ds "github.com/najibjodiansyah/data-structure-go/ds"
)

func main() {
	// newArray := ds.NewDynamycArray()
	// newArray.Push("hi")
	// newArray.Push("you")
	// newArray.Push("!")
	// newArray.Delete(0)
	// newArray.Delete(1)
	// newArray.Print()
	// fmt.Println(reverseString2("najib"))
	// fmt.Println(mergeSortArray([]int{0, 3, 4, 31}, []int{4, 6, 30}))

	myHashTable := ds.NewHashTable(5)
	myHashTable.Set("grapes", 10000)
	myHashTable.Set("apples", 9)
	myHashTable.Set("oranges", 2)
	fmt.Println(myHashTable)
	fmt.Println(myHashTable.Get("oranges"))
}

func reverseString1(str string) string {
	var reversed string
	for _, v := range str {
		fmt.Println(string(v), reversed)
		reversed = string(v) + reversed
	}
	return reversed
}

func reverseString2(str string) string {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func mergeSortArray(arr1, arr2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result

}
