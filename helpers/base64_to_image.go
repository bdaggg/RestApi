package helpers

import (
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"strings"
)

func Base64ToImage(base64String string, outputPath string) error {
	base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = png.Encode(outFile, img)
	if err != nil {
		return err
	}
	return nil
}
