package main

import (
	"fmt"
	"os"

	"github.com/Harichandra-Prasath/Ascii-Image-Generator/utils"
)

func main() {

	file, err := os.Open("image.jpg")
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
	fmt.Print(brightness_array)
}
