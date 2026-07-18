package main

import (
	"fmt"

	sc "github.com/nilese1/Ascii-Rasterizer/scene"
)

const CHARS = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

type pixel struct {
	r int
	g int
	b int

	light float64
}

func moveCursor(lines uint32, move_up bool) {
	char := 'B'
	if move_up {
		char = 'A'
	}

	fmt.Printf("\033[%v%c", lines, char)
}

func setColour(r int, g int, b int) {
	fmt.Printf("\033[38;2;%v;%v;%vm", r, g, b)
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func getChar(light float64) string {
	chosen_char := int(light * float64(len(CHARS)))
	char_inx := len(CHARS) - chosen_char - 1

	return string(CHARS[char_inx])
}

func printScreen(pixels [][]pixel, scene *sc.Scene, headers ...string) {
	setColour(255, 255, 255)
	for _, header := range headers {
		fmt.Print(header, "\n")
	}

	for _, row := range pixels {
		for _, pixel := range row {
			setColour(pixel.r, pixel.g, pixel.b)
			char := getChar(pixel.light)

			fmt.Print(char)
		}

		fmt.Print("\n")
	}

	moveCursor(scene.ScreenHeight+uint32(len(headers)), true)
}
