package main

import (
	"algos_practice/euclid"
	"algos_practice/factorial"
	"algos_practice/models/linkedlist"
	"algos_practice/power"
	"algos_practice/sorting/bubblesort"
	"algos_practice/sorting/mergesort"
	"algos_practice/sorting/quicksort"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("------ euclid -------")
	fmt.Println(euclid.GenerateGCD(96, 60))
	fmt.Println(euclid.GenerateGCD(20, 8))
	fmt.Println(euclid.GenerateGCDRecursive(96, 60))
	fmt.Println(euclid.GenerateGCDRecursive(20, 8))

	fmt.Println("------linkedlist-----")
	child2 := linkedlist.NewNode(8)
	child1 := linkedlist.NewNodeWithNext(6, child2)
	root := linkedlist.NewNodeWithNext(5, child1)
	findExists := root.Find(6)
	findNotExists := root.Find(69)
	fmt.Printf("Find node 6 %v \n", findExists)
	fmt.Printf("Find node 69 %v \n", findNotExists)
	fmt.Println("Append node 69 since it's not existing")
	child3 := linkedlist.NewNode(69)
	root.Append(child3)
	findSixtyNinde := root.Find(69)
	fmt.Printf("Find node 69 %v \n", findSixtyNinde)

	fmt.Println("------factorial-----")
	fmt.Printf("%v! = %v\n", 5, factorial.Factorial(5))
	fmt.Printf("%v! = %v\n", 10, factorial.Factorial(10))

	fmt.Println("------power------")
	fmt.Printf("%v^%v = %v\n", 5, 3, power.Power(5, 3))
	fmt.Printf("%v^%v = %v\n", 10, 3, power.Power(10, 3))

	fmt.Println("------Sorting-------")
	arr := []int{42, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10}
	copy := bubblesort.Sort(arr)
	fmt.Println(copy)
	arrMergeSort := []int{7, 6, 5, 4, 3, 2, 1}
	sortedMergeSort := mergesort.Sort(arrMergeSort)
	fmt.Println(sortedMergeSort)
	largeDataset := make([]int, 100)
	for i := 1; i < 100; i++ {
		largeDataset[i-1] = rand.Intn(i * 10000)
	}
	largeMergeSortedData := mergesort.Sort(largeDataset)
	fmt.Println(largeMergeSortedData)

	largeDatasetTwo := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}
	quicksort.Sort(largeDatasetTwo)
	fmt.Println(largeDatasetTwo)
}
