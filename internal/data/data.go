package data

import (
	"image"
	"mime/multipart"
	"os"

	"path/filepath"

	"github.com/chai2010/webp"
)

func CreateWEBP(file *multipart.FileHeader) ([]byte, error) {
	// Opening
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Reading
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	// Creating WebP
	webpPath := filepath.Join("uploads", file.Filename+".webp")
	webpFile, err := os.Create(webpPath)
	if err != nil {
		return nil, err
	}
	defer webpFile.Close()

	data, err := webp.EncodeExactLosslessRGBA(img)

	return data, err
}

func isAllowedExt(ext string) bool {
	allowedExt := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
		".webp": true,
	}
	return allowedExt[ext]
}
