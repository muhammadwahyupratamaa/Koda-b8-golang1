package main

import "fmt"

func luas(jari float32, phi float32) { 

	result := phi * float32(jari) * float32(jari)

  fmt.Printf("Luas Keliling lingkaran  :%f\n" , result )
}

func keliling(jari int , phi float32) {
	result := 2 * float32(jari) * phi
 fmt.Printf("Keliling Lingkaran : %f\n", result)
}

func main() {
	phi := float32 (22/7)
	luas(21,phi)
	keliling(21,phi)
}
