package main

import (
	"fmt"
)

func main() {
	name := "Nigel Poulton"
	course := "Getting started with Kubernetes"
	const c = 186000

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("You're currently watching", course)

	fmt.Println(c)

}

func updateCourse(course *string) string {
	*course = "geting Started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
