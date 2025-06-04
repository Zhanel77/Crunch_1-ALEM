package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// Функция для печати верхней границы карты
func printBorder(width int) {
	// Проходимся по ширине карты
	for i := 0; i < width; i++ {
		// Выводим символы границы "_______"
		for _, char := range "_______" {
			ap.PutRune(char) // Функция PutRune используется для вывода символа
		}
	}
	ap.PutRune('\n') // Переход на новую строку после печати границы
}

// Функция для печати строки карты
func printRow(row string, width int) {
	// Два прохода, так как строка состоит из двух частей (верхней и нижней)
	for i := 0; i < 2; i++ {
		ap.PutRune('|') // Вывод вертикальной линии слева
		for _, cell := range row {
			if cell == '2' { // Если символ '2', рисуем стрелку '>'
				if i == 1 {
					for _, char := range "  >   " {
						ap.PutRune(char)
					}
				} else {
					for _, char := range "      " {
						ap.PutRune(char)
					}
				}
			} else if cell == '0' { // Если символ '0', это препятствие 'XXXXXX'
				for _, char := range "XXXXXX" {
					ap.PutRune(char)
				}
			} else if cell == '3' { // Если символ '3', рисуем '@'
				if i == 1 {
					for _, char := range "  @   " {
						ap.PutRune(char)
					}
				} else {
					for _, char := range "      " {
						ap.PutRune(char)
					}
				}
			} else {
				for _, char := range "      " {
					ap.PutRune(char)
				}
			}
			ap.PutRune('|') // Вертикальный разделитель между клетками
		}
		ap.PutRune('\n') // Переход на новую строку
	}

	// Нижняя граница строки
	for _, cell := range row {
		if cell == '0' {
			for _, char := range "|XXXXXX" {
				ap.PutRune(char)
			}
		} else {
			for _, char := range "|______" {
				ap.PutRune(char)
			}
		}
	}
	ap.PutRune('|') // Закрываем строку справа
	ap.PutRune('\n')
}

// Функция для вывода всей карты
func printMap(mapData []string, h, w int) {
	printBorder(w) // Печатаем верхнюю границу
	for _, row := range mapData {
		printRow(row, w) // Печатаем каждую строку карты
	}
}

func main() {
	var h, w int
	for {
		// Читаем размеры карты h (высота) и w (ширина)
		if _, err := fmt.Scanf("%d %d", &h, &w); err == nil && h > 0 && w > 0 {
			break // Выход из цикла при корректном вводе
		}
	}

	// Создаем массив строк для хранения карты
	mapData := make([]string, h)
	for i := 0; i < h; i++ {
		for {
			// Читаем строку карты
			if _, err := fmt.Scanf("%s", &mapData[i]); err == nil && len(mapData[i]) == w {
				break // Выход из цикла при корректном вводе строки
			}
		}
	}

	// Выводим карту
	printMap(mapData, h, w)
}
