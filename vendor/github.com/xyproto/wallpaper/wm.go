package wallpaper

import (
	"errors"
)

type WM interface {
	Name() string
	ExecutablesExists() bool
	Running() bool
	SetWallpaper(imageFilename string) error
}

var WMs = []WM{
	&Sway{},
	&Plasma{},
	&Weston{},
	&Generic{},
}

// Set the desktop wallpaper (filled/stretched), for any supported windowmanager.
// The fallback is to use `feh`.
func Set(imageFilename string) error {
	var lastErr error
	// Loop through all available WM structs
	for _, wm := range WMs {
		if wm.ExecutablesExists() && wm.Running() {
			if err := wm.SetWallpaper(imageFilename); err != nil {
				lastErr = err
				switch wm.Name() {
				case "Weston":
					// If the current windowmanager is Weston, no method is currently available
					return err
				default:
					// Try the next one
					continue
				}
			} else {
				return nil
			}
		}
	}
	if lastErr != nil {
		return errors.New("found no working way of setting the desktop wallpaper: " + lastErr.Error())
	}
	return errors.New("found no working way of setting the desktop wallpaper")
}
