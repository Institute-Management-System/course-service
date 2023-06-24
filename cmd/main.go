package main

import "github.com/Institute-Management-System/course-service/boot"

func main() {
	if err := boot.Boot(); err != nil {
		panic(err)
	}
}
