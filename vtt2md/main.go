package main

// vtt2md - A simple tool to convert WebVTT subtitles to Markdown format.

import (
	"fmt"
	"log"
	"os"
	"strings"

	webvtt "github.com/Songmu/go-webvtt"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	vtt, err := webvtt.ParseAll(os.Stdin)
	if err != nil {
		return err
	}

	var currentSpeaker string
	var currentTexts []string

	flush := func() {
		if currentSpeaker != "" && len(currentTexts) > 0 {
			fmt.Printf("### %s\n", currentSpeaker)
			fmt.Println(strings.Join(currentTexts, " "))
			fmt.Println()
		}
	}

	for _, cue := range vtt.Cues {
		for _, voice := range cue.Voices {
			if voice.Speaker != currentSpeaker {
				flush()
				currentSpeaker = voice.Speaker
				currentTexts = nil
			}
			if text := strings.TrimSpace(voice.Text); text != "" {
				currentTexts = append(currentTexts, text)
			}
		}
	}
	flush()

	return nil
}
