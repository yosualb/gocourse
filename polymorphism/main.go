package main

import "fmt"
import "reflect"

// Interface
type Vehicle interface {
	Check()
	Override()
}

// Interface
type FakeVehicle interface {
	Check()
}

type Car struct {
	name string
}

// Implement function
func (c *Car) Check() {
	fmt.Println("Checking", c.name)
}

func (c *Car) Override() {
	fmt.Println("Overriding", c.name)
}

func (c *Car) Ride() {
	fmt.Println("Riding")
}

type Plane struct {
	kind string
}

// Implement function
func (p *Plane) Check() {
	fmt.Println("Checking", p.kind)
}

func (p *Plane) Override() {
	fmt.Println("Overriding", p.kind)
}

func (p *Plane) Fly() {
	fmt.Println("Flying")
}

type Jet struct {
	// Embedded interface
	Vehicle
}

func (j *Jet) Check() {
	fmt.Println("Checking Jet")
}

func (j *Jet) Fly() {
	fmt.Println("Flying")
}

func ParseToJet(v Vehicle) Vehicle {
	return &Jet{v}
}

// Empty interface
type Everything interface{}

func main() {
	var v Vehicle
	c := &Car{"Chevy"}
	p := &Plane{"Boeing 737"}

	v = c
	fmt.Println("v:", reflect.TypeOf(v))
	v.Check()
	v.Override()
	// Error, interface Vehicle doesn't have this function
	// v.Ride()

	v = p
	fmt.Println("v:", reflect.TypeOf(v))
	v.Check()
	v.Override()
	// Error, interface Vehicle doesn't have this function
	// v.Fly()

	var fV FakeVehicle
	fV = c
	fmt.Println("fV:", reflect.TypeOf(fV))
	fV.Check()
	// Error, interface FakeVehicle doesn't have these functions
	// fV.Ride()
	// fV.Override()

	fV = p
	fmt.Println("fV:", reflect.TypeOf(fV))
	fV.Check()
	// Error, interface FakeVehicle doesn't have these functions
	// fV.Fly()
	// fV.Override()

	j := ParseToJet(p)
	j.Check()
	j.Override()
	fmt.Println("j:", reflect.TypeOf(j))

	var e Everything
	e = c
	fmt.Println("e:", reflect.TypeOf(e))
	e = p
	fmt.Println("e:", reflect.TypeOf(e))
	e = j
	fmt.Println("e:", reflect.TypeOf(e))
}
