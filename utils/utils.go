package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/nfnt/resize"
)

type Pixel struct {
	R int
	G int
	B int
}

func getPixel(R uint32, G uint32, B uint32, a uint32) Pixel {
	return Pixel{int(R / 257), int(G / 257), int(B / 257)}
}

func GetPixelsArray(file io.Reader) ([][]Pixel, error) {
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Print("error in decoding")
		return nil, err
	}
	scaled_image := resize.Resize(120, 120, img, resize.Lanczos2)
	var pixels [][]Pixel

	// get the dimensions
	bounds := scaled_image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	//fmt.Print(width, height)
	for y := 0; y < height; y++ {
		var curr_row []Pixel
		for x := 0; x < width; x++ {
			curr_row = append(curr_row, getPixel(scaled_image.At(x, y).RGBA()))
		}
		pixels = append(pixels, curr_row)
	}
	return pixels, nil
}

func getbrightness_Average(pixel Pixel) int {
	return (pixel.B + pixel.G + pixel.R) / 3
}

func getbrightness_luminosity(pixel Pixel) int {
	return int(0.21*float64(pixel.R) + 0.72*float64(pixel.G) + 0.07*float64(pixel.B))
}

func getbrightness_Lightness(pixel Pixel) int {
	max := math.Max(float64(pixel.B), float64(pixel.G))
	max = math.Max(max, float64(pixel.R))
	min := math.Min(float64(pixel.B), float64(pixel.G))
	min = math.Min(min, float64(pixel.R))
	return (int(max) + int(min)) / 2
}

type conversion func(Pixel) int

func GetBrightnessArray(pixels [][]Pixel, method *string) [][]int {
	height := len(pixels)
	width := len(pixels[0])
	var selected_method conversion
	switch *method {
	case "average":
		selected_method = getbrightness_Average
	case "lightness":
		selected_method = getbrightness_Lightness
	case "luminosity":
		selected_method = getbrightness_luminosity
	}

	var brightness_array [][]int
	for i := 0; i < height; i++ {
		var curr_row []int
		for j := 0; j < width; j++ {
			curr_row = append(curr_row, selected_method(pixels[i][j]))
		}
		brightness_array = append(brightness_array, curr_row)
	}

	return brightness_array

}

func convert(value int) string {
	str := "`^\",:;Il!i~+_-?][{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$}"
	index := value / 4
	return string(str[index])
}

func Brit_to_ascii(brightness_array [][]int) [][]string {
	height := len(brightness_array)
	width := len(brightness_array[0])
	var ascii_array [][]string
	for i := 0; i < height; i++ {
		var curr_row []string
		for j := 0; j < width; j++ {
			curr_row = append(curr_row, convert(brightness_array[i][j]))
		}
		ascii_array = append(ascii_array, curr_row)
	}
	return ascii_array
}

func Generate(array [][]string) string {
	var res bytes.Buffer

	height := len(array)
	width := len(array[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k := 0
			for k < 3 {
				res.WriteString(array[i][j])
				k++
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}
