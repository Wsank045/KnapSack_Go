package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/* A brute force recursive implementation of 0-1 Knapsack problem
modified from: https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10 */

func Max(x, y items) items {
	if x.value > y.value {
		return x
	}
	return y
}

type items struct {
	names  string
	value  int
	weight int
}

func (item *items) add(item_1 items) items {
	names := item.names + " " + item_1.names
	value := item.value + item_1.value
	weight := item.weight + item_1.weight
	item_f := items{names, value, weight}

	return item_f
}

// Returns the maximum value that
// can be put in a knapsack of capacity W
func KnapSack(W int, myItems []items) items {

	// Base Case
	if len(myItems) == 0 || W == 0 {
		return items{"", 0, 0}
	}
	last := len(myItems) - 1

	// If weight of the nth item is more
	// than Knapsack capacity W, then
	// this item cannot be included
	// in the optimal solution
	if myItems[last].weight > W {
		return KnapSack(W, myItems[:last])

		// Return the maximum of two cases:
		// (1) nth item included
		// (2) item not included
	} else {

		return Max(myItems[last].add(KnapSack(W-myItems[last].weight, myItems[:last])), KnapSack(W, myItems[:last]))
	}
}

func KnapSackConc(w int, myItems []items, res items, valueChan chan items, wg *sync.WaitGroup, n int) {

	if n == 128 {
		if w >= 0 {
			res_1 := KnapSack(w, myItems)
			res_3 := res.add(res_1)
			valueChan <- res_3
		}
		wg.Done()
		return
	}

	if len(myItems) == 0 {
		if w >= 0 {
			valueChan <- res
		}
		wg.Done()
		return

	}
	last := len(myItems) - 1
	n = n * 2
	go KnapSackConc(w-myItems[last].weight, myItems[:last], res.add(myItems[last]), valueChan, wg, n)
	go KnapSackConc(w, myItems[:last], res, valueChan, wg, n)

}
func arrayMax(array []items) items {
	max := array[0]

	for i := 0; i < len(array); i++ {
		if max.value < array[i].value {
			max = array[i]
		}
	}
	return max
}

// Driver code
func main() {

	fmt.Println("Enter your file's name")
	var entry_file string
	fmt.Scanln(&entry_file)

	file, ferr := os.Open(entry_file)
	defer file.Close()
	if ferr != nil {
		panic(ferr)
	}

	file_reader := bufio.NewScanner(file)
	file_reader.Scan()
	number_of_items_string := file_reader.Text()
	number_of_items, err := strconv.Atoi(number_of_items_string)
	if err != nil {
		panic(err)
	}

	names := make([]string, number_of_items)
	weights := make([]int, number_of_items)
	values := make([]int, number_of_items)

	for i := 0; i < number_of_items; i++ {
		file_reader.Scan()
		line := file_reader.Text()
		line_content := strings.Split(line, " ")

		names[i] = line_content[0]
		weights[i], _ = strconv.Atoi(line_content[2])
		values[i], _ = strconv.Atoi(line_content[1])

	}
	file_reader.Scan()
	W_string := file_reader.Text()
	W, _ := strconv.Atoi(W_string)

	fmt.Println("Number of cores: ", runtime.NumCPU())

	// simple example
	//W := 40
	//names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	//weights := []int{1, 2, 3, 5, 6, 10, 5, 6, 9, 11}
	//values := []int{1, 6, 10, 15, 6, 8, 3, 9, 3, 5}

	//W := 7000
	//names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b"}
	//weights := []int{112, 80, 305, 185, 174, 665, 170, 460, 159, 197, 112, 184, 200, 138, 95, 54, 71, 147, 235, 301, 50, 174, 68, 14, 14, 30, 27, 40}
	//values := []int{40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 13, 11, 8, 20}

	n := float64(len(weights))
	goroutines := math.Pow(2, n)

	collection := make([]items, len(names))

	for i := 0; i < len(names); i++ {
		collection[i] = items{names[i], values[i], weights[i]}
	}

	chanValue := make(chan items)
	defer close(chanValue)

	goValues := make([]items, 12)

	start := time.Now()

	var wg sync.WaitGroup

	if len(weights) > 7 {
		wg.Add(128)
	} else {
		wg.Add(int(goroutines))
	}

	init_item := items{"", 0, 0}
	KnapSackConc(W, collection, init_item, chanValue, &wg, 1)

	go func() {
		for {
			goValues = append(goValues, <-chanValue)
		}
	}()

	wg.Wait()
	println()
	y := arrayMax(goValues)
	fmt.Printf("Total value: %d \nItems selected: %s", y.value, strings.TrimSpace(y.names))
	end := time.Now()
	fmt.Println()
	fmt.Printf("Total runtime 1: %s\n", end.Sub(start))
	total_runtime_1 := end.Sub(start)

	exit_file := entry_file + ".sol"
	f, err := os.Create(exit_file)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Total value: " + strconv.Itoa(y.value) + "\nItems selected: " + y.names)
	f.WriteString("\nTotal Runtime: " + total_runtime_1.String())

	start1 := time.Now()
	result := KnapSack(W, collection)

	fmt.Println()
	fmt.Printf("Items selected:  %s\nvalue %d", strings.TrimSpace(result.names), result.value)
	end1 := time.Now()
	println()
	fmt.Printf("Total runtime 2: %s\n", end1.Sub(start1))

}
