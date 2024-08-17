package main

import "fmt"

func main() {
	fmt.Println("Buat persegi")
	createSquare(5)
	fmt.Println("---")

	fmt.Println("Buat garis border kanan kotak")
	createRightBorderSquare(5)
	fmt.Println("---")

	fmt.Println("Buat kotak kosong")
	emptySquare(5)
	fmt.Println("---")

	fmt.Println("Buat segitiga")
	createTriangle(5)
	fmt.Println("---")

	fmt.Println("Buat diagonal")
	createDiagonal(5)
	fmt.Println("---")

	fmt.Println("Buat diagonal 2")
	createDiagonal2(5)
	fmt.Println("---")

	fmt.Println("Buat belah ketupat")
	createRhombus(3)
}

func createSquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func createRightBorderSquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func emptySquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func createTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func createDiagonal(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("#")
	}
}

func createDiagonal2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("#")
	}
}

func createRhombus(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
