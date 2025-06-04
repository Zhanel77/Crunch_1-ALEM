package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// printBorder печатает верхнюю и нижнюю границы карты
func printBorder(width int) {
	for i := 0; i < width; i++ {
		for _, char := range "_______" { // 7 символов на каждую ячейку
			ap.PutRune(char)
		}
	}
	ap.PutRune('\n')
}

// printRow печатает строку карты с заданными символами
func printRow(row string, width int, wall, player, award rune) {
	// Верхняя часть ячеек (две строки высоты)
	for i := 0; i < 2; i++ {
		ap.PutRune('|') // Левый край строки
		for _, cell := range row {
			if cell == '0' { // Стена
				for _, char := range []rune{wall, wall, wall, wall, wall, wall} {
					ap.PutRune(char)
				}
			} else if cell == '2' { // Игрок
				if i == 1 {
					for _, char := range []rune{' ', ' ', player, ' ', ' ', ' '} {
						ap.PutRune(char)
					}
				} else {
					for _, char := range "      " {
						ap.PutRune(char)
					}
				}
			} else if cell == '3' { // Награда
				if i == 1 {
					for _, char := range []rune{' ', ' ', award, ' ', ' ', ' '} {
						ap.PutRune(char)
					}
				} else {
					for _, char := range "      " {
						ap.PutRune(char)
					}
				}
			} else { // Пустые клетки
				for _, char := range "      " {
					ap.PutRune(char)
				}
			}
			ap.PutRune('|') // Разделитель между клетками
		}
		ap.PutRune('\n') // Переход на новую строку
	}

	// Нижняя часть ячеек (граница между строками)
	for _, cell := range row {
		if cell == '0' {
			for _, char := range []rune{'|', wall, wall, wall, wall, wall, wall} {
				ap.PutRune(char)
			}
		} else {
			for _, char := range "|______" {
				ap.PutRune(char)
			}
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')
}

// printMap печатает карту с заданными символами
func printMap(mapData []string, h, w int, wall, player, award rune) {
	printBorder(w) // Печатаем верхнюю границу
	for _, row := range mapData {
		printRow(row, w, wall, player, award) // Печатаем каждую строку
	}
}

func main() {
	var h, w int // Размер карты

	// Ввод размеров карты (h - высота, w - ширина)
	for {
		if _, err := fmt.Scanf("%d %d", &h, &w); err == nil && h > 0 && w > 0 {
			break
		}
	}

	// Чтение карты
	mapData := make([]string, h)
	for i := 0; i < h; i++ {
		for {
			if _, err := fmt.Scanf("%s", &mapData[i]); err == nil && len(mapData[i]) == w {
				break
			}
		}
	}

	// Ввод новых символов для стены, игрока и награды
	var wall, player, award rune
	fmt.Scanf("%c%c%c", &wall, &player, &award)

	// Вывод карты с новыми символами
	printMap(mapData, h, w, wall, player, award)
}
