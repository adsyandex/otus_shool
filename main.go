package main

import (
	"fmt"
)

// createChessBoard создает шахматную доску заданного размера.
func createChessBoard(size int) string {
	board := ""
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if (row+col)%2 == 0 {
				board += " " // Белая клетка
			} else {
				board += "#" // Чёрная клетка
			}
		}
		board += "\n" // Переход на новую строку
	}
	return board
}

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scan(&size)
	fmt.Println(createChessBoard(size))
}
