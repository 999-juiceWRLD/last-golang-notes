package main

import "fmt"

func concatter() func(string) string {
    doc := ""
    return func(word string) string {
        doc += word + " "
        return doc;
    }
}

func _() {
	aggregator := concatter();
	aggregator("mr.");
	aggregator("beast");

	fmt.Println(aggregator("harmonie granger"));
}