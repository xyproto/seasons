package main

type WM interface {
	Name() string
	ExecutablesExists() bool
	Running() bool
	SetWallpaper(imageFilename string) error
}
