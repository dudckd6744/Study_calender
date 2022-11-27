package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
	volume() float64
}

type RectangularParallelepiped struct {
	line1, line2, line3 float64
}

type Cylinder struct {
	radius, height float64
}

func (r RectangularParallelepiped) area() float64 {
	return (2*r.line1*r.line2) + (2*r.line1*r.line3) + (2*r.line2*r.line3)
}

func (r RectangularParallelepiped) volume() float64 {
	return r.line1*r.line2*r.line3
}

func (c Cylinder) area() float64 {
	return (math.Pi * c.radius * c.radius) * 2 + (2*math.Pi*c.radius)*c.height
}
func (c Cylinder) volume() float64 {
	return (math.Pi * c.radius * c.radius) * c.height
}

func main() {
	r1 := RectangularParallelepiped{10.5, 20.2,20}
	c1 := Cylinder{10, 10}
	r2 := RectangularParallelepiped{4,10,23}
	c2 := Cylinder{4.2,15.6}

	printMeasure(c1,c2 ,r1, r2)
}

func printMeasure(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Printf("%.2f, ",val.area()) //인터페이스의 메소드 호출
		fmt.Printf("%.2f\n",val.volume()) //인터페이스의 메소드 호출
	}
}