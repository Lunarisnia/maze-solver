package ds

import "math"

type Vector2 []int

func (v Vector2) Sub(y Vector2) Vector2 {
	return Vector2{v[0] - y[0], v[1] - y[1]}
}

func (v Vector2) Length() float64 {
	x := float64(v[0])
	y := float64(v[1])
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func (v Vector2) MoveTo(y Vector2) Vector2 {
	return v.Sub(y)
}

func (v Vector2) Equal(y Vector2) bool {
	return v[0] == y[0] && v[1] == y[1]
}
