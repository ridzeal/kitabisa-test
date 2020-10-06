package main

import (
	"fmt"
	"kitabisa-test/controller"
)

func _main() {
	total, apple, cake := controller.BundleIt(25, 23)

	fmt.Printf("Summary %v %v %v\n", total, apple, cake)
}