package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// Function to print the top column labels (A, B, C...) with more spacing
func printColumnLabels(width int) {
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ')

	for i := 0; i < width; i++ {
		ap.PutRune(' ')
		ap.PutRune(' ')
		ap.PutRune(' ')
		ap.PutRune(rune('A' + i)) // Column labels
		ap.PutRune(' ')
		ap.PutRune(' ')
		ap.PutRune(' ')

	}
	ap.PutRune('\n')
}

// Function to print the top border
func printBorder(width int) {
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ')

	for i := 0; i < width; i++ {
		for _, char := range "_______" {
			ap.PutRune(char)
		}
	}
	ap.PutRune('\n')
}

// Function to print each row with **labels outside the grid**
func printRow(row string, rowIndex int) {
	// Print row label before the row (1, 2, 3...)
	ap.PutRune(' ')

	ap.PutRune(rune('0' + rowIndex))
	ap.PutRune(' ')

	// empty upper cells
	for range row {
		ap.PutRune('|')
		for _, char := range "      " {
			ap.PutRune(char)
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')

	// cell content
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ')
	for _, cell := range row {
		ap.PutRune('|')
		if cell == '2' {
			for _, char := range "  >   " {
				ap.PutRune(char)
			}
		} else if cell == '0' {
			for _, char := range "XXXXXX" {
				ap.PutRune(char)
			}
		} else if cell == '3' {
			for _, char := range "  @   " {
				ap.PutRune(char)
			}
		} else {
			for _, char := range "      " {
				ap.PutRune(char)
			}
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')

	// closing row
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ') // Align with labels
	for _, cell := range row {
		ap.PutRune('|')
		if cell == '0' {
			for _, char := range "XXXXXX" {
				ap.PutRune(char)
			}
		} else {
			for _, char := range "______" {
				ap.PutRune(char)
			}
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')
}

// Function to print the entire map
func printMap(mapData []string, h, w int) {
	printColumnLabels(w) // Print shifted column labels
	printBorder(w)       // Print top border
	for i, row := range mapData {
		printRow(row, i+1) // Print row labels outside the grid
	}
}

// Main function
func main() {
	var h, w int

	// Read height and width
	fmt.Scanf("%d %d", &h, &w)

	// Read map data
	mapData := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scanf("%s", &mapData[i])
	}

	// Print the final map with **properly aligned labels**
	printMap(mapData, h, w)
}
