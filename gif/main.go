package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{
	color.RGBA{255, 255, 255, 255}, // white
	color.RGBA{255, 0, 0, 255},     // red
	color.RGBA{0, 255, 0, 255},     // green
	color.RGBA{0, 0, 255, 255},     // blue
	color.RGBA{255, 255, 0, 255},   // yellow
	color.RGBA{0, 255, 255, 255},   // cyan
	color.RGBA{255, 0, 255, 255},   // magenta
	color.RGBA{0, 0, 0, 255},       // black
}


func main() {
	// f,err := os.Create("ls.gif")
	// if err!=nil{
	// 	panic(err)
	// }

	// defer f.Close()
	// lissajous(f)
	startServer()
}

func lissajous(out io.Writer){
	const (
		cycles =5
		res =0.001
		size =100
		nframes =64
		delay =8
	)
	freq := rand.Float64() *3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	nColors := len(palette)
	for i:=0;i<nframes;i++{
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect,palette)
		colorIndex := uint8((i % (nColors - 1)) + 1)
		// fmt.Printf("colorIndex: %v\n", colorIndex)
		for t:=0.0;t<cycles*2*math.Pi; t+= res{
			x := math.Sin(t)
			y:= math.Sin(t*freq +phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			colorIndex)
		}
		phase+=0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out,&anim)
}

func startServer(){
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){
		lissajous(w)
	})
	fmt.Println("Server started on port 8000")
	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000",nil)
	
}