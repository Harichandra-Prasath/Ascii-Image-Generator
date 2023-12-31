package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Harichandra-Prasath/Ascii-Image-Generator/utils"
)

func main() {
	path := flag.String("path", "test.jpeg", "Path of the image (Jpg,Jpeg,Png)")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: Ascii-generator -path path of the image")
	}

	file, err := os.Open(*path)
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
