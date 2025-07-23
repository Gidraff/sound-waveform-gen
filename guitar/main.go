package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	sampleRate = 48000
	duration   = 5 // seconds
	freq       = 440.0
	volume     = 0.5
)

func main() {
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	spec := sdl.AudioSpec{
		Freq:     sampleRate,
		Format:   sdl.AUDIO_F32SYS, // 32-bit float samples
		Channels: 2,
		Samples:  1024,
	}
	obtainedSpec := &sdl.AudioSpec{}

	deviceID, err := sdl.OpenAudioDevice("", false, &spec, obtainedSpec, 0)
	if err != nil {
		log.Fatal("Failed to open audio device:", err)
	}
	defer sdl.CloseAudioDevice(deviceID)

	// Generate audio samples
	sampleCount := sampleRate * duration
	samples := make([]float32, sampleCount*2) // Stereo
	phase := 0.0
	phaseInc := 2 * math.Pi * freq / float64(sampleRate)
	for i := 0; i < sampleCount; i++ {
		sample := float32(math.Sin(phase) * volume)
		samples[i*2] = sample   // Left
		samples[i*2+1] = sample // Right
		phase += phaseInc
	}

	// Convert float32 slice to byte slice using binary.Write
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, samples)
	if err != nil {
		log.Fatal("Failed to convert samples to byte buffer:", err)
	}

	err = sdl.QueueAudio(deviceID, buf.Bytes())
	if err != nil {
		log.Fatal("QueueAudio failed:", err)
	}

	sdl.PauseAudioDevice(deviceID, false)
	sdl.Delay(uint32(duration * 1000)) // Play for 'duration' seconds
}
