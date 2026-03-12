package main

import "fmt"

func main() {

	// Задание 1
	// numbers := [3][5]int{}

	// for i, row := range numbers {
	// 	for j, _ := range row {
	// 		fmt.Scan(&numbers[i][j])
	// 	}
	// }

	// fmt.Println(numbers)

	// max_i := 0
	// max_j := 0

	// max_value := 0

	// for i, row := range numbers {
	// 	for j, column := range row {
	// 		if column > max_value {
	// 			max_value = column
	// 			max_i = i
	// 			max_j = j
	// 		}
	// 	}
	// }

	// fmt.Printf("Максимальная число %d, находятся на строке %d и на столбце %d", max_value, max_i, max_j)

	// Задание 2

	// matrix := [11][11]string{}

	// for i, row := range matrix {
	// 	for j, _ := range row {
	// 		matrix[i][j] = "."
	// 		if i == j || i+j == len(matrix)-1 || i == len(matrix)/2 || j == len(matrix)/2 {
	// 			matrix[i][j] = "*"
	// 		}
	// 		fmt.Print(matrix[i][j], " ")

	// 	}
	// 	fmt.Println()
	// }

	// Задание 3

	// chessboard := [8][8]string{}

	// for i := 0; i < len(chessboard); i++ {
	// 	for j := 0; j < len(chessboard[i]); j++ {
	// 		if (i+j)%2 == 0 {
	// 			chessboard[i][j] = "."
	// 		} else {
	// 			chessboard[i][j] = "*"
	// 		}
	// 		fmt.Print(chessboard[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	// Задание 4

	// matrix := [4][4]int{}

	// for i, row := range matrix {
	// 	for j, _ := range row {
	// 		fmt.Print("Введите число: ")
	// 		fmt.Scan(&matrix[i][j])
	// 	}
	// }
	// fmt.Println(matrix)
	// fmt.Println()
	// fmt.Println()

	// var i1, i2 int

	// fmt.Print("Введите номер первой строки: ")
	// fmt.Scan(&i1)

	// fmt.Print("Введите номер второй строки: ")
	// fmt.Scan(&i2)

	// matrix[i1], matrix[i2] = matrix[i2], matrix[i1]

	// for i := 0; i < len(matrix); i++ {
	// 	fmt.Print(matrix[i], " ")
	// }

	// Задание 5

	matrix := [4][4]int{}

	for i, row := range matrix {
		for j := range row {
			fmt.Print("Введите число: ")
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println(matrix)
	fmt.Println()
	fmt.Println()

	var j1, j2 int

	fmt.Print("Введите номер первого столбца: ")
	fmt.Scan(&j1)

	fmt.Print("Введите номер второго столбца: ")
	fmt.Scan(&j2)

	for i := 0; i < len(matrix); i++ {
		matrix[i][j1], matrix[i][j2] = matrix[i][j2], matrix[i][j1]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
