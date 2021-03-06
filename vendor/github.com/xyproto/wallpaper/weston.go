package wallpaper

import (
	"errors"
	"fmt"
	"os"
)

// Weston windowmanager detector
type Weston struct {
}

func (s *Weston) Name() string {
	return "Weston"
}

func (s *Weston) ExecutablesExists() bool {
	return which("weston") != ""
}

func (s *Weston) Running() bool {
	return containsE("GDMSESSION", "weston") || containsE("XDG_SESSION_DESKTOP", "weston") || containsE("XDG_CURRENT_DESKTOP", "weston")
}

// SetWallpaper sets the desktop wallpaper, given an image filename.
// The image must exist and be readable.
func (s *Weston) SetWallpaper(imageFilename string) error {

	// TODO: Add the following to ~/.config/weston.ini (or whichever configuration file Weston uses)
	fmt.Println("WESTON CONFIG FILE: ", os.Getenv("WESTON_CONFIG_FILE"))

	// [shell]
	// background-image=/home/user/somewhere/image.jpg
	// background-type=scale

	return errors.New("Weston currently does not support changing the desktop wallpaper at runtime")
}
