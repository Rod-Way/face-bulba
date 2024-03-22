package data

func c() {

}

// func CreateWEBP(file *multipart.FileHeader) ([]byte, error) {
// 	// Opening
// 	src, err := file.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer src.Close()

// 	// Reading
// 	img, _, err := image.Decode(src)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Creating WebP
// 	webpPath := filepath.Join("uploads", file.Filename+".webp")
// 	webpFile, err := os.Create(webpPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer webpFile.Close()

// 	// data, err := webp.EncodeExactLosslessRGBA(img)

// 	return data, err
// }

func isAllowedExt(ext string) bool {
	allowedExt := map[string]bool{
		".dng":  true,
		".raw":  true,
		".png":  true,
		".jpeg": true,
		".jpg":  true,
		".gif":  true,
		".mp4":  true,
		".mov":  true,
	}
	return allowedExt[ext]
}

func isAllowedFileSize(size int64) bool {
	if size > 50*1024*1024 { // 50 MB
		return false
	}

	return true
}
