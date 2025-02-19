package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

func fatal(message string, errors ...error) {
	if len(errors) > 0 {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %s\n", message, errors[0]); err != nil {
			panic(err)
		}
	} else {
		if _, err := fmt.Fprintf(os.Stderr, "%s\n", message); err != nil {
			panic(err)
		}
	}

	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fatal("No directory provided")
	}

	dir := args[0]

	entries, err := os.ReadDir(dir)

	if err != nil {
		fatal("Failed to list directory", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		imagePath := fmt.Sprintf("%s/%s", dir, entry.Name())

		file, err := os.Open(imagePath)

		if err != nil {
			if er := file.Close(); er != nil {
				panic(er)
			}

			fatal(fmt.Sprintf("Failed to open file '%s'", imagePath), err)
		}

		image, _, err := image.DecodeConfig(file)

		if err != nil {
			if er := file.Close(); er != nil {
				panic(er)
			}

			fatal(fmt.Sprintf("Failed to decode image '%s'", imagePath), err)
		}

		imageNameParts := strings.Split(entry.Name(), ".")

		imageName := imageNameParts[0]

		imageCode := fmt.Sprintf(":%s:", imageName)

		fmt.Printf(
			"'%s', '%d', '%d', '1', '%s', '%s', \n",
			entry.Name(),
			image.Width,
			image.Height,
			imageName,
			imageCode,
		)
	}
}
