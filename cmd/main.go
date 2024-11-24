package main

import (
	"io"
	"os"
	"time"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	fyneContainer "fyne.io/fyne/v2/container"
	fyneWidget "fyne.io/fyne/v2/widget"
)

const (
	Title  = "GoUtil"
	Width  = 200
	Height = 100
)

func copyDir(src string, dst string) error {

	if err := os.RemoveAll(dst); err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := src + "/" + entry.Name()
		dstPath := dst + "/" + entry.Name()

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func main() {

	app := fyneApp.New()

	// create window
	window := app.NewWindow(Title)

	// set window size
	window.Resize(fyne.NewSize(Width, Height))

	textWidget := fyneWidget.NewLabel("")
	textWidget.Alignment = fyne.TextAlignCenter

	window.SetContent(fyneContainer.NewVBox(
		textWidget,
		fyneWidget.NewButton("Transfer chart", func() {

			textWidget.SetText("")
			err := copyDir("C:/Users/root/Desktop/lab projects/binary-option/option-chart/dist", "C:/Users/root/Desktop/lab projects/binary-option/web-terminal/packages/option-chart/dist")
			if err != nil {
				textWidget.SetText("Error: " + err.Error())
			} else {
				textWidget.SetText("File copied successfully!")
				go func() {

					// clear the text after 3 seconds
					<-time.After(3 * time.Second)
					textWidget.SetText("")
				}()
			}
		}),
	))

	window.ShowAndRun()
}
