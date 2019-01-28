package goTourPic

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8{

	a := make([][]uint8, dy)

	for i := range a {
		for j := 0; j < dx; j++ {
			a[i] = append(a[i], uint8(i*j))
		}
	}

	return a
}

func goTourPic() {
	pic.Show(Pic)
}

