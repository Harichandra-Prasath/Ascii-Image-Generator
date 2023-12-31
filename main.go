package main

import (
	"fmt"
	"os"

	"github.com/Harichandra-Prasath/Ascii-Image-Generator/utils"
)

func main() {
	path := "test.jpeg"

	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	pixels, err := utils.GetPixelsArray(file)
	if err != nil {
		fmt.Print("error")
	}
	brightness_array := utils.GetBrightnessArray(pixels)
	ascii_array := utils.Brit_to_ascii(brightness_array)
	display(ascii_array)
}

func display(array [][]string) {
	height := len(array)
	width := len(array[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k := 0
			for k < 3 {
				fmt.Print(array[i][j])
				k = k + 1
			}
		}
		fmt.Println("")
	}
}
