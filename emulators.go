package emulators

/*
#include <stdlib.h>
*/
import "C"

import (
	"math/rand"
)

type CPPCore struct {
	FuckUpProbability float64
}

func createDeadlock() {
	a := make(chan struct{})
	<- a
}

func createMemoryLeak() {
	C.malloc(100000)
}

func createUB() {
	panic("It's an UB. And you got UBed!!!")
}

func createDoubleFree() {
	ptr := C.malloc(100000)
	C.free(ptr)
	C.free(ptr)
}

func createDivisionByZero() {
	a := 5
	b := 0
	_ = a / b 
}

func shootYourselfInTheFoot() {
	a := []func(){createDeadlock, createDivisionByZero, createDoubleFree, createMemoryLeak, createUB}

	a[rand.Int() % len(a)]()
}

func (c CPPCore) FuckAround() {
	if rand.Float64() > c.FuckUpProbability {
		return
	}
	shootYourselfInTheFoot()
}
