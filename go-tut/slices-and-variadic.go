package main

import "fmt"

func sum(numbers ...int) int {
	num := 0;
	for i := 0; i < len(numbers); i++ {
		num += numbers[i];
	}

	return num;
}

func _() {
	mySlice1 := make([]int, 5);
	fmt.Println(mySlice1);

	mySlice2 := []string{"bonjour", "et", "bienvenue"};
	fmt.Println(mySlice2);

	total := sum(1, 2, 3, 4);
	fmt.Println(total);

	mySlice2 = append(mySlice2, "life", "is", "good");
	fmt.Println(mySlice2);
}