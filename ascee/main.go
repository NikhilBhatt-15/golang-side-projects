package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"
	"strings"
)

func main() {
    file, err := os.Open("C:\\Users\\nikhi\\less-go\\ascee\\goku.jpg")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error in opening file: %v\n", err)
        return
    }
    defer file.Close()

    img, err := jpeg.Decode(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error decoding JPEG: %v\n", err)
        return
    }

    // Choose your mode
    ascii := imageToASCIIWithEdges(img, 80) // Combined approach
    // ascii := imageToASCII(img, 80)        // Original
    // ascii := edgeDetectionASCII(img, 80)  // Pure edge detection
    
    fmt.Println(ascii)
}

func imageToASCIIWithEdges(img image.Image, width int) string {
    bounds := img.Bounds()
    imgWidth := bounds.Max.X - bounds.Min.X
    imgHeight := bounds.Max.Y - bounds.Min.Y
    height := (imgHeight * width) / (imgWidth * 2)
    
    var result strings.Builder
    
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            imgX := bounds.Min.X + (x * imgWidth) / width
            imgY := bounds.Min.Y + (y * imgHeight) / height
            
            // Get edge strength
            edgeStrength := getSobelEdge(img, imgX, imgY)
            
            // Get grayscale for areas without edges
            r, g, b, _ := img.At(imgX, imgY).RGBA()
            gray := (299*r + 587*g + 114*b + 500) / 1000
            gray8 := uint8(gray >> 8)
            
            var char byte
            if edgeStrength > 100 { // Strong edge threshold
                // Use edge characters
                chars := "|/-\\+*"
                charIndex := int(edgeStrength) * (len(chars) - 1) / 255
                if charIndex >= len(chars) {
                    charIndex = len(chars) - 1
                }
                char = chars[charIndex]
            } else {
                // Use regular ASCII characters for non-edge areas
                chars := "@#W$9876543210?!abc;:+=-,._"
                charIndex := int(gray8) * (len(chars) - 1) / 255
                if charIndex >= len(chars) {
                    charIndex = len(chars) - 1
                }
                char = chars[charIndex]
            }
            
            result.WriteByte(char)
        }
        result.WriteByte('\n')
    }
    
    return result.String()
}

func edgeDetectionASCII(img image.Image, width int) string {
    bounds := img.Bounds()
    imgWidth := bounds.Max.X - bounds.Min.X
    imgHeight := bounds.Max.Y - bounds.Min.Y
    height := (imgHeight * width) / (imgWidth * 2)
    
    // Characters for different edge strengths
    chars := " .-:=+*#%@"
    
    var result strings.Builder
    
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            imgX := bounds.Min.X + (x * imgWidth) / width
            imgY := bounds.Min.Y + (y * imgHeight) / height
            
            edgeStrength := getSobelEdge(img, imgX, imgY)
            
            charIndex := int(edgeStrength) * (len(chars) - 1) / 255
            if charIndex >= len(chars) {
                charIndex = len(chars) - 1
            }
            
            result.WriteByte(chars[charIndex])
        }
        result.WriteByte('\n')
    }
    
    return result.String()
}

func getSobelEdge(img image.Image, x, y int) uint8 {
    bounds := img.Bounds()
    
    // Sobel kernels
    sobelX := [3][3]int{
        {-1, 0, 1},
        {-2, 0, 2},
        {-1, 0, 1},
    }
    
    sobelY := [3][3]int{
        {-1, -2, -1},
        { 0,  0,  0},
        { 1,  2,  1},
    }
    
    var gx, gy float64
    
    // Apply Sobel operators
    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            px := x + dx
            py := y + dy
            
            // Handle boundaries
            if px < bounds.Min.X {
                px = bounds.Min.X
            }
            if px >= bounds.Max.X {
                px = bounds.Max.X - 1
            }
            if py < bounds.Min.Y {
                py = bounds.Min.Y
            }
            if py >= bounds.Max.Y {
                py = bounds.Max.Y - 1
            }
            
            r, g, b, _ := img.At(px, py).RGBA()
            gray := (299*r + 587*g + 114*b + 500) / 1000
            intensity := float64(gray >> 8)
            
            gx += intensity * float64(sobelX[dy+1][dx+1])
            gy += intensity * float64(sobelY[dy+1][dx+1])
        }
    }
    
    // Calculate gradient magnitude
    magnitude := math.Sqrt(gx*gx + gy*gy)
    
    // Normalize to 0-255
    if magnitude > 255 {
        magnitude = 255
    }
    
    return uint8(magnitude)
}

// Original ASCII function (for comparison)
func imageToASCII(img image.Image, width int) string {
    chars := " .,:-=+*#%@"
    
    bounds := img.Bounds()
    imgWidth := bounds.Max.X - bounds.Min.X
    imgHeight := bounds.Max.Y - bounds.Min.Y
    height := (imgHeight * width) / (imgWidth * 2)
    
    var result strings.Builder
    
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            imgX := bounds.Min.X + (x * imgWidth) / width
            imgY := bounds.Min.Y + (y * imgHeight) / height
            
            r, g, b, _ := img.At(imgX, imgY).RGBA()
            gray := (299*r + 587*g + 114*b + 500) / 1000
            gray8 := uint8(gray >> 8)
            
            charIndex := int(gray8) * (len(chars) - 1) / 255
            if charIndex >= len(chars) {
                charIndex = len(chars) - 1
            }
            
            result.WriteByte(chars[charIndex])
        }
        result.WriteByte('\n')
    }
    
    return result.String()
}