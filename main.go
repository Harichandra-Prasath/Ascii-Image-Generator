package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
}

func getPixel(R uint32, G uint32, B uint32, a uint32) Pixel {
	return Pixel{int(R / 257), int(G / 257), int(B / 257)}
}

func getPixelsArray(file io.Reader) ([][]Pixel, error) {

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Print("error in decoding")
		return nil, err
	}
	var pixels [][]Pixel

	// get the dimensions
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		var curr_row []Pixel
		for x := 0; x < width; x++ {
			curr_row = append(curr_row, getPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, curr_row)
	}
	return pixels, nil
}
func main() {
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open("image.jpg")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	pixels, err := getPixelsArray(file)
	if err != nil {
		fmt.Print("error")
	}
	fmt.Print(pixels)

}
