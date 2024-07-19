package main

import "fmt"

type employee interface {
	getName()	string
	getSalary()	int
}

type contractor struct {
	name 			string
	hourlyPay		int
	hoursPerYear	int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func _() {
	cont := contractor{ "jackson", 35, 160 * 12 + 60}
	name := cont.getName()
	sal := cont.getSalary()

	fmt.Printf("name: %s | salary: $%d\n", name, sal);
}