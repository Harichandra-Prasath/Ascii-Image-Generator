package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Harichandra-Prasath/Ascii-Image-Generator/utils"
)

func main() {
	path := flag.String("path", "test.jpeg", "Path of the image (Jpg,Jpeg,Png)")
	method := flag.String("method", "average", "Method for brightness conversion (Average,luminosity,lightness)")
	//save := flag.Bool("save", false, "Option to save the ascii in a text file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: Ascii-generator -path path of the image")
		os.Exit(1)
	}

	*method = strings.ToLower(*method)
	if *method != "average" && *method != "luminosity" && *method != "lightness" {
		fmt.Println("Invalid method.Use -h tag for help")
		os.Exit(1)
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
	brightness_array := utils.GetBrightnessArray(pixels, method)
	ascii_array := utils.Brit_to_ascii(brightness_array)
	Art := utils.Generate(ascii_array)
	fmt.Print(Art)
}
