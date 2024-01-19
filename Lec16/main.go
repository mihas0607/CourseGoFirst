package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(newX, newY float64) *Point {
	return &Point{newX, newY}
}

// Return X coord of Point
func (p *Point) GetX() float64 {
	return p.x
}

// Return Y coord of Point
func (p *Point) GetY() float64 {
	return p.y
}

// Set new X coord for Existing point
func (p *Point) SetX(newX float64) {
	p.x = newX
}

// Set new Y coord for Existing point
func (p *Point) SetY(newY float64) {
	p.y = newY
}

// Retunrt string info about point in format Point{X: value, Y: value}
func (p *Point) Stringify() string {
	return fmt.Sprintf("Point{X: %.0f, Y: %.0f}", p.GetX(), p.GetY())
}

type Line struct {
	p1 Point
	p2 Point
}

func NewLine(newP1, newP2 Point) *Line {
	return &Line{newP1, newP2}
}

// Return p1 of line
func (l *Line) GetP1() Point {
	return l.p1
}

// Return p2 of line
func (l *Line) GetP2() Point {
	return l.p2
}

// Set new p1 for Existing Line
func (l *Line) SetP1(newP1 Point) {
	l.p1 = newP1
}

// Set new p2 for Existing Line
func (l *Line) SetP2(newP2 Point) {
	l.p2 = newP2
}

// Calc length of segment p1p2
func (l *Line) Length() float64 {
	return math.Sqrt(math.Pow(l.p2.x-l.p1.x, 2) + math.Pow(l.p2.y-l.p1.y, 2))
}

// Method return True if p3 belongs to p1p2 segment
//
//	False otherwise
func (l *Line) IsOnSegment(p3 Point) bool {
	p1 := l.GetP1()
	p2 := l.GetP2()
	l2 := NewLine(p1, p3)
	l2.SetP1(p1)
	l2.SetP2(p3)
	l3 := NewLine(p3, p2)
	l3.SetP1(p3)
	l3.SetP2(p2)

	if l.Length() == l2.Length()+l3.Length() {
		return true
	}
	return false
}

// //Retunrt string info about line in format Line{P1: value, P2: value}
func (l *Line) Stringify() string {
	return fmt.Sprintf("Line{P1: %v, P2: %v}", l.p1, l.p2)
}

func main() {
	pointSlice := make([]float64, 0, 6)
	var coord float64

	for i := 0; i < 6; i++ {
		fmt.Scan(&coord)
		pointSlice = append(pointSlice, coord)
	}

	p1 := NewPoint(0, 0)
	p1.SetX(pointSlice[0])
	p1.SetY(pointSlice[1])
	p2 := NewPoint(0, 0)
	p2.SetX(pointSlice[2])
	p2.SetY(pointSlice[3])
	p3 := NewPoint(0, 0)
	p3.SetX(pointSlice[4])
	p3.SetY(pointSlice[5])
	l1 := NewLine(*p1, *p2)
	l1.SetP1(*p1)
	l1.SetP2(*p2)

	if l1.IsOnSegment(*p3) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// fmt.Println(p1.Stringify())
	// fmt.Println(p2.Stringify())
	// fmt.Println(l1.Stringify())
	// fmt.Println(l1)
	// fmt.Println(l1.Length())
	// fmt.Println(l1.IsOnSegment(*p3))

}
