package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func InstanciateBrowser() string {
	possibleBrowsers := []string{"brave", "chromium", "chrome", "edge"}
	var browserPath string

	for _, browser := range possibleBrowsers {
		browserPath, _ = exec.LookPath(browser)
		fmt.Println("found chromium-based browser, using : ", browserPath)
		if browserPath != "" {
			break
		}
	}
	if browserPath == "" {
		fmt.Println("no browser found")
	}

	var l string
	var browserErr error
	// change to headless(false) to see the browser in action
	l, browserErr = launcher.New().Headless(false).Bin(browserPath).Launch()
	if browserErr != nil {
		fmt.Println("error when launching browser : ", browserErr)
	}

	return l
}

func screenshotChallenge(page *rod.Page) {
	// takes a screenshot of the canvas, cropps it, and saves it in a global variable
	canvasScreenshotBytes := page.MustElement(gameCanvas).MustScreenshot("currentStep.png")

	buf := bytes.NewBuffer(canvasScreenshotBytes)

	// Decode the image using image.Decode
	canvasScreenshot, _, err := image.Decode(buf)
	if err != nil {
		// handle error
	}

	// imageFile, err := os.Open("currentStep.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imageObject, err := png.Decode(imageFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// topLeft := []int{356, 19}
	// bottomRight := []int{731, 429}
	currentStepScreenshot = canvasScreenshot.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(357, 20, 731, 429))

	outputFile, outputErr := os.Create("currentStep-cropped.png")
	if outputErr != nil {
		log.Fatal(outputErr)
	}
	png.Encode(outputFile, currentStepScreenshot)

}

func getRandomCoordinatesInGroup(group *[][2]int) [2]int {
	randomItem := rand.Intn(len(*group))
	// + and - are for safety and also no one clicks on the very edge of the item
	randomX0Padding := rand.Intn(itemSize-10) + 5
	randomY0Padding := rand.Intn(itemSize-10) + 5

	fmt.Println("Item position in grid : ", (*group)[randomItem])

	fmt.Printf("random padding on item : %v, %v\n", randomX0Padding, randomY0Padding)

	randomCoordinates := [2]int{
		(*group)[randomItem][1]*(itemSize+itemSpacing) + randomX0Padding,
		(*group)[randomItem][0]*(itemSize+itemSpacing) + randomY0Padding,
	}

	fmt.Println("Coordinates in grid : ", randomCoordinates)
	return randomCoordinates
}
