package handlers

import (
	"bytes"
	"fmt"
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
		return nil, fmt.Errorf("error fetching image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
	}

	// Read the image data from the response body into a buffer
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image data: %v", err)
	}

	img, err := jpeg.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %v", err)
	}

	log.Println("Image decoded successfully")
	return img, nil
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
