package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"

	"github.com/nfnt/resize"
)

// pixel struct
type Pixel struct {
	R int
	G int
	B int
}

func getPixel(R uint32, G uint32, B uint32, a uint32) Pixel {
	return Pixel{int(R / 257), int(G / 257), int(B / 257)}
}

// getting the pixel array of scaled image
func GetPixelsArray(file io.Reader) ([][]Pixel, error) {
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Print("error in decoding")
		return nil, err
	}
	scaled_image := resize.Resize(120, 120, img, resize.Lanczos2) //scaling due to finite size of display

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

// fucntion to get brightness matrix from pixel matrix
func GetBrightnessArray(pixels [][]Pixel, method *string) [][]int {
	height := len(pixels)
	width := len(pixels[0])
	var selected_method conversion // based on user prefered conversion method
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

// function to map brightness to ascii characters
func convert(value int) string {
	str := "`^\",:;Il!i~+_-?][{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$}"
	// characters are chosen in reference to black background like terminal
	index := value / 4
	return string(str[index])
}

// fuction to convert the brightness matrix to ascii
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

// function to generate the ascii art from the ascii matrix
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

func get_path(path string) string {
	var output_path bytes.Buffer
	var index int
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			index = i
		}
	}
	output_path.WriteString(path[:index+1])
	output_path.WriteString("output.txt")
	return output_path.String()
}

// func to save the art in text file

func Save(art string, path string) {
	output_path := get_path(path)
	file, err := os.OpenFile(output_path, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s", art)

}
