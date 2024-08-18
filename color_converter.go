package main

import (
	"fmt"
	"math"
	"strconv"
)

type XY struct {
  X, Y float64
}

func HexToRGB(hex string) (float64, float64, float64, error) {
  if (hex[0] != '#') {
    hex = fmt.Sprintf("#%v", hex)
  }

	if len(hex) != 7 || hex[0] != '#' {
		return 0, 0, 0, fmt.Errorf("invalid hex color")
	}
	r, err := strconv.ParseInt(hex[1:3], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	return float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0, nil
}

func RGBToXY(r, g, b float64) (float64, float64) {
	r = gammaCorrect(r)
	g = gammaCorrect(g)
	b = gammaCorrect(b)

	X := r*0.649926 + g*0.103455 + b*0.197109
	Y := r*0.234327 + g*0.743075 + b*0.022598
	Z := r*0.000000 + g*0.053077 + b*1.035763

	if X+Y+Z == 0 {
		return 0, 0
	}
	x := X / (X + Y + Z)
	y := Y / (X + Y + Z)
	return x, y
}

func gammaCorrect(value float64) float64 {
	if value > 0.04045 {
		return math.Pow((value+0.055)/(1.055), 2.4)
	}
	return value / 12.92
}

func getXYFromHex(hex string) XY{
  red, green, blue, err := HexToRGB(hex)

  if err != nil {
    fmt.Println("Invalid hex color:", err)
  }

  x, y := RGBToXY(red, green, blue) 
  
  return XY{
    X: x,
    Y: y,
  }
}

// Method to convert XY to Hex
func (xy XY) ToHex() string {
    // Convert XY to RGB (this is a placeholder, actual conversion needed)
    r, g, b := xyToRGB(xy.X, xy.Y)

    // Convert RGB to Hex
    hex := fmt.Sprintf("#%02X%02X%02X", int(r*255), int(g*255), int(b*255))
    return hex
}

// Placeholder function for XY to RGB conversion
func xyToRGB(x, y float64) (float64, float64, float64) {
    // Assuming a fixed brightness for simplicity
    brightness := 1.0

    if y == 0 {
        return 0.0, 0.0, 0.0
    }

    z := 1.0 - x - y
    Y := brightness
    X := (Y / y) * x
    Z := (Y / y) * z

    r := X*1.656492 - Y*0.354851 - Z*0.255038
    g := -X*0.707196 + Y*1.655397 + Z*0.036152
    b := X*0.051713 - Y*0.121364 + Z*1.011530

    // Apply gamma correction
    r = gammaCorrect(r)
    g = gammaCorrect(g)
    b = gammaCorrect(b)

    // Clamp values to [0, 1]
    r = math.Max(0, math.Min(1, r))
    g = math.Max(0, math.Min(1, g))
    b = math.Max(0, math.Min(1, b))

    return r, g, b
}
