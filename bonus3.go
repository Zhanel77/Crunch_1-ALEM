package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// ANSI-коды цветов для раскрашивания клеток карты
const (
	Reset  = "\033[0m"  // Сброс цвета (возвращает стандартный цвет)
	Red    = "\033[41m" // Красный фон (стены)
	White  = "\033[47m" // Белый фон (свободные клетки)
	Blue   = "\033[44m" // Синий фон (игрок)
	Yellow = "\033[43m" // Желтый фон (награда)
)

// printBorder печатает верхнюю и нижнюю границу карты
func printBorder(width int) {
	for i := 0; i < width; i++ {
		for _, char := range "_______" { // 7 символов для каждой ячейки
			ap.PutRune(char)
		}
	}
	ap.PutRune('\n')
}

// getColor возвращает строку с ANSI-кодом цвета для заданной ячейки карты.
func getColor(cell rune) string {
	// Определяем цвет клетки на основе её значения
	if cell == '0' {
		return Red // Стена
	} else if cell == '1' {
		return White // Свободное пространство
	} else if cell == '2' {
		return Blue // Игрок
	} else if cell == '3' {
		return Yellow // Награда
	} else {
		return Reset // По умолчанию (без цвета)
	}
}

// printRow печатает строку карты, окрашивая каждую клетку
func printRow(row string, width int) {
	// Верхняя часть ячеек (две строки высоты)
	for i := 0; i < 2; i++ {
		ap.PutRune('|') // Левый край строки
		for _, cell := range row {
			color := getColor(cell)                         // Получаем цвет для текущей клетки
			for _, char := range color + "      " + Reset { // Заполняем цветом всю клетку (6 пробелов)
				ap.PutRune(char)
			}
			ap.PutRune('|') // Разделитель между клетками
		}
		ap.PutRune('\n') // Переход на новую строку
	}

	// Нижняя часть ячеек (граница между строками)
	ap.PutRune('|') // Левый край строки
	for _, cell := range row {
		color := getColor(cell)                         // Получаем цвет
		for _, char := range color + "______" + Reset { // Заполняем цветом нижнюю часть клетки (6 символов)
			ap.PutRune(char)
		}
		ap.PutRune('|') // Разделитель между клетками
	}
	ap.PutRune('\n') // Переход на новую строку
}

// printMap печатает всю карту
func printMap(mapData []string, h, w int) {
	printBorder(w) // Печатаем верхнюю границу карты
	for _, row := range mapData {
		printRow(row, w) // Печатаем каждую строку карты
	}
}

// main — основная функция, в которой считываются данные и запускается вывод карты
func main() {
	var h, w int // Высота и ширина карты

	// Вводим размеры карты (h - высота, w - ширина)
	for {
		if _, err := fmt.Scanf("%d %d", &h, &w); err == nil && h > 0 && w > 0 {
			break // Если данные введены корректно, выходим из цикла
		}
	}

	// Создаем карту (двумерный массив строк)
	mapData := make([]string, h)
	for i := 0; i < h; i++ {
		for {
			if _, err := fmt.Scanf("%s", &mapData[i]); err == nil && len(mapData[i]) == w {
				break // Убеждаемся, что длина строки соответствует ширине карты
			}
		}
	}

	// Выводим карту на экран
	printMap(mapData, h, w)
}
