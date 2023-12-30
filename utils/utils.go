package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
)

func getPixel(R uint32, G uint32, B uint32, a uint32) Pixel {
	return Pixel{int(R / 257), int(G / 257), int(B / 257)}
}

func GetPixelsArray(file io.Reader) ([][]Pixel, error) {
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
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

type Pixel struct {
	R int
	G int
	B int
}
