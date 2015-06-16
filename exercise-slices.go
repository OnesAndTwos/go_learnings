package main

import "golang.org/x/tour/pic"

// Pic - pic
func Pic(dx, dy int) [][]uint8 {
	var a = make([][]uint8, dy)

	for i := range a {
		a[i] = make([]uint8, dx)

		for j := range a[i] {
			a[i][j] = uint8(i * j % 255)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
