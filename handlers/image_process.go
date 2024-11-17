package handlers

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func DownloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching image: %v", err)
	}
	defer resp.Body.Close()

	// Read the image data from the response body into a buffer
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading image data: %v", err)
	}

	img, err := jpeg.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
	}

	log.Println("Image decoded successfully")
	_ = img
	// fmt.Print(imageData)
	return img, err
}

func ProcessImage(url string) (int, error) {
	img, err := DownloadImage(url)
	if err != nil {
		return 0, err
	}

	// Get image dimensions
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	// Calculate perimeter
	perimeter := 2 * (height + width)

	// Simulate GPU processing
	sleepDuration := time.Duration(rand.Intn(300)+100) * time.Millisecond
	time.Sleep(sleepDuration)

	return perimeter, nil
}
