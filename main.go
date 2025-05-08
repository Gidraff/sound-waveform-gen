package main

import (
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/xigh/go-waveform"
	"github.com/xigh/go-wavreader"
)

func main() {
	// TODO:

	// bring in the package
	// upload an audio file and open it
	// read audio file with waveform package
	// generate a colorful png image

	audioFile, err := os.Open("./test/test_file1.wav")
	if err != nil {
		log.Fatal("Error opening the file")
		return
	}
	defer audioFile.Close()

	w0, err := wavreader.New(audioFile)
	if err != nil {
		log.Fatal("Error reading the file")
	}
	img := waveform.MinMax(w0, &waveform.Options{
		Width:   1800,
		Height:  400,
		Zoom:    1.7,
		Half:    false,
		MarginL: 20,
		MarginR: 15,
		MarginT: 15,
		MarginB: 20,
		Front: &color.NRGBA{
			R: 255,
			G: 128,
			B: 0,
			A: 150,
		},
		Back: &color.NRGBA{
			A: 0, // opacity
		},
	})

	imageFile, err := os.Create("./test_image.png")
	if err != nil {
		log.Fatal("failed to create the file")
	}

	defer imageFile.Close()

	err = png.Encode(imageFile, img)
	if err != nil {
		log.Fatal(err)
	}
}
