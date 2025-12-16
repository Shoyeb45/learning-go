package main

import "fmt"

func main() {
	var intSlice = []int32{123, 123, 13, 12, 13, 2, 3}
	fmt.Println(sumSlice[int32](intSlice));

	var floatSlice = []float32{1.342, 23.232};

	fmt.Println(sumSlice(floatSlice));
}


func sumSlice[T int32 | int64 | float32 | float64](slice []T) T {
	var sum T;
	for _, v := range slice {
		sum += v;
	}

	return sum;
} 