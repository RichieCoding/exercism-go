// Package triangle determines if given 3 sides of a triangle is it a valid triangle
package triangle

import "math"

type Kind int

const (
	NaT = iota 
	Equ        
	Iso        
	Sca        
)

// KindFromSides determines given its sides if it is a triangle
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isTriangle(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b, a == c, b == c:
		return Iso
	default:
		return Sca
	}
}

func isValidNumber(n float64) bool {
	return n != 0 && !math.IsInf(n, 0)
}

func isTriangle(a, b, c float64) bool {
	return isValidNumber(a*b*c) && a+b >= c && a+c >= b && b+c >= a
}
