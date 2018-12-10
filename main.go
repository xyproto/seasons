package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

func main() {
	WMs := []WM{
		&Sway{},
		&Generic{},
	}
	for _, wm := range WMs {
		if wm.ExecutablesExists() && wm.Running() {
			fmt.Println("WindowManager: " + wm.Name())
			matches, err := filepath.Glob("/usr/share/pixmaps/*.png")
			if err != nil {
				panic(err)
			}
			rand.Seed(time.Now().Unix())
			randomIndex := rand.Int() % len(matches)
			imageFilename := matches[randomIndex]
			fmt.Println("Setting background image to: " + imageFilename)
			wm.SetWallpaper(imageFilename)
			break
		}
	}
}
