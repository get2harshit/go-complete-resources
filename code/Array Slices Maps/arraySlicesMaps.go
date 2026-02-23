package main

import (
	"fmt"
)



func main(){
	fmt.Println("Learning Arrays, Slices and Maps")

	// arrays
	// var prices [4]float64 = [4]float64{1,2,3,4}
	prices2:= [4]float64{1,2,4,5}
	prices2[3] = 10
	fmt.Println(prices2)

	// Task 1: Create an array of size 4 storing random strings

	// Task 2.1 : Take user input array size and values
	// Task 2.2 : Take user input array values with given fixed size 4
	// var size int
	// fmt.Println("Enter size of array")
	// fmt.Scan(&size)
	// var arr [4]int64
	// fmt.Println(len(arr), cap(arr))

	// for i:=0; i<4; i++ {
	// 	fmt.Println("Enter value ", i)
	// 	fmt.Println(len(arr), cap(arr))
	// 	fmt.Scan(&arr[i])
	// }
	// fmt.Println(arr)
	// fmt.Println(len(arr), cap(arr))

	// slices
	arr2:= make([]int64, 2, 4)
	fmt.Println(len(arr2), cap(arr2))
	fmt.Println(arr2)
	// arr2[2] = 3
	// arr2[3] = 6

	arr2 = append(arr2, 3)

	fmt.Println(arr2)

	// slice on array
	housePrices:= [4]int{300, 400, 200, 100} 
	slicedArray:= housePrices[1:]
	fmt.Println(slicedArray)
	slicedArray[0] = 900
	fmt.Println(slicedArray)
	fmt.Println(housePrices)

	// slicing on sliced array

	finalSlicedArray:= slicedArray[0:]
	fmt.Println(finalSlicedArray)

	a:=[3]int{1,2,3}
	b:=a

	b[0] = 100
	fmt.Println(a)


	// maps
	myWebsites:= map[string]string {
		"pst" : "polariscampus.com",
		"google": "google.com",
	}

	// insertion in maps
	myWebsites["youtube"] = "youtube.com"

	// deletion in maps
	// delete(myWebsites, "google")

	fmt.Println(myWebsites)
	fmt.Println(myWebsites["pst"])


	// 
	var myArr [4]int
	myArr[0] = 1
	fmt.Println(myArr)

	// var mySlice []int
	// mySlice[0] = 1
	// fmt.Println(mySlice)

	var myMap map[string]int
	// myMap["pst"] = 5
	fmt.Println(myMap)

	// slices
	// var s []int 
	// fmt.Println(s)

	// sl:= []int{1,2,3}
	// fmt.Println(sl)

	sli:= make([]int, 2, 4)
	fmt.Println(sli)
	// [0 0 _ _ ]
	sli = append(sli, 10)
	fmt.Println(sli)

	// ...

	// insert an element at index 2 value 5
	// Hint: use append function

	// [0 0 10]
	// [0 0 5 10]
	// sli[:index] + value + sli[index:]
	index:=2
	value:=5
	sli = append(sli[:index], append([]int{value}, sli[index:]...)...)
	fmt.Println(sli)

	// remove an element at index 2

	housePrices2:= []int{599, 499, 399, 299, 199, 99}
	index2:= 2

	housePrices2 = append(housePrices2[:index2], housePrices2[index2+1:]...)
	fmt.Println(housePrices2)
}
