// Lab 7 -- functions, pt 3
//
// Programmer name: Nigel S Adams
// Date completed: 10/3/19

package main

import "math"

// AreaOfCircle computes the area of a circle.
func AreaOfCircle(radius float64) float64 {
	return (math.Pi * math.Pow(radius, 2))
}

// AreaOfRectangle computes the area of a rectangle.
func AreaOfRectangle(width, height float64) float64 {
	return (width * height)
}

// VolumeOfCylinder computes the volume of a cylinder.
func VolumeOfCylinder(radius, height float64) float64 {
	return (math.Pi * math.Pow(radius, 2)) * height
}

// VolumeOfBox computes the volume of a cube.
func VolumeOfBox(length, width, height float64) float64 {
	return (length * width) * height
}

// DistanceBetween computes the distance between points x1,y1 and x2,y2.
// HINT: use the Pythagorean Theorem.
func DistanceBetween(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// HoursMinutesSeconds computes the number of hours, minutes,
// and seconds given a number of seconds.
// EXAMPLE: 3661 seconds is 1 hours, 1 minutes, 1 seconds
func HoursMinutesSeconds(seconds int) (int, int, int) {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	Msec := ((seconds % 3600) % 60)

	return hours, minutes, Msec
}

func main() {
}
