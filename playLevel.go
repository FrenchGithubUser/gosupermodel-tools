package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func playLevel(page *rod.Page) {

	page.MustWaitLoad()
	time.Sleep(time.Second)

	for {

		screenshotChallenge(page)

		groups, noGroups := findItemGroups()

		if noGroups {
			break
		}

		chosenGroup, chosenGroupName := strategyBiggestFirst(&groups)

		coordinatesToClick := getRandomCoordinatesInGroup(&chosenGroup)
		canvasBounds, canvasErr := (*page.MustElement(gameCanvas)).Shape()
		if canvasErr != nil {
			fmt.Println("Couldn't get canvas bounds : ", canvasErr)
		}
		spaceFromPageTopToCanvas := (*canvasBounds).Quads[0][0]
		spaceFromPageLeftToCanvas := (*canvasBounds).Quads[0][1]

		canvasCoordinatesToClickX := gridMarginFromCanvasLeft + coordinatesToClick[0]
		canvasCoordinatesToClickY := gridMarginFromCanvasTop + coordinatesToClick[1]

		fmt.Printf("Coordinates to click on canvas : %v, %v\n", canvasCoordinatesToClickX, canvasCoordinatesToClickY)

		pageCooridinatesToClickX := float64(spaceFromPageLeftToCanvas) + float64(canvasCoordinatesToClickX)
		pageCooridinatesToClickY := float64(spaceFromPageTopToCanvas) + float64(canvasCoordinatesToClickY)

		fmt.Printf("Clicking on coordinates in grid : %v of %s group (%v items). Coordinates in page : %v, %v\n", coordinatesToClick, chosenGroupName, len(chosenGroup), pageCooridinatesToClickX, pageCooridinatesToClickY)
		fmt.Printf("Moving mouse to item location\n")

		page.Mouse.MustMoveTo(pageCooridinatesToClickX, pageCooridinatesToClickY)
		time.Sleep(time.Second * 1)
		page.MustScreenshot("page.png")
		// (*page.MustElement(gameCanvas)).MustFrame().Mouse.MustMoveTo(float64(gridMarginFromCanvasLeft)+float64(coordinatesToClick[0]), float64(gridMarginFromCanvasTop)+float64(coordinatesToClick[1]))
		clickErr := page.Mouse.Click(proto.InputMouseButtonLeft, 1)
		if clickErr != nil {
			fmt.Println(clickErr)
		} else {
			fmt.Printf("Clicked\n")
		}

		// sleep in between actions
		time.Sleep(time.Second * 2)
	}

	fmt.Println("Level ended")

}
