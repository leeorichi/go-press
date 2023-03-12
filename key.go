package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/faiface/beep"
    "github.com/faiface/beep/speaker"
    "github.com/faiface/beep/wav"
)

func main() {
    // Set sample rate and buffer size for speaker
    sampleRate := beep.SampleRate(44100)
    bufferSize := beep.BufferSize(1024)

    // Initialize speaker
    err := speaker.Init(sampleRate, bufferSize)
    if err != nil {
        panic(err)
    }

    // Load sound file
    f, err := os.Open("ding.wav")
    if err != nil {
        panic(err)
    }

    // Decode sound file
    streamer, format, err := wav.Decode(f)
    if err != nil {
        panic(err)
    }

    // Play sound when key is pressed
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Press any key: ")
        _, _, err := reader.ReadRune()
        if err != nil {
            panic(err)
        }
        speaker.Play(streamer)
    }
}
