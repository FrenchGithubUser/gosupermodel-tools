package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func findItemGroups() (map[string][][][2]int, bool) {
	imageFile, err := os.Open("currentStep-cropped.png")
	if err != nil {
		log.Fatal(err)
	}
	currentStepScreenshot, err := png.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	// matrix representing the image of the current game state
	currentStateBase := [rows][columns]string{}
	currentState := make([][]string, rows)
	for i := 0; i < rows; i++ {
		currentState[i] = make([]string, columns)
		for j := 0; j < columns; j++ {
			currentState[i][j] = currentStateBase[i][j]
		}
	}

	items := map[string][]int{
		"bag":   []int{164, 78, 22},
		"shoe":  []int{177, 115, 204},
		"pants": []int{223, 113, 172},
		"shirt": []int{165, 16, 90},
	}

	for row := 1; row <= rows; row++ {
		for column := 1; column <= columns; column++ {

			x0 := (column-1)*itemSize + (column-1)*itemSpacing
			y0 := (row-1)*itemSize + (row-1)*itemSpacing
			x1 := x0 + itemSize
			y1 := y0 + itemSize

			subImage := currentStepScreenshot.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(image.Rect(x0, y0, x1, y1))
			imgColors := subImage.At(x0+itemSize/2, y0+itemSize/2).(color.RGBA)
			currentRGBValues := []int{int(imgColors.R), int(imgColors.G), int(imgColors.B)}

			var correspondingItem bool
			var currentItem string
			for itemName, values := range items {
				currentItem = ""
				for index, value := range values {
					correspondingItem = true
					// fmt.Println(currentRGBValues[index], value)
					if currentRGBValues[index] != value {
						correspondingItem = false
					}
				}
				if correspondingItem {
					currentItem = itemName
					break
				}
			}
			// fmt.Println(currentItem == ""+"\n")

			currentState[row-1][column-1] = currentItem

			// save item image file to disk
			// outputFile, outputErr := os.Create(fmt.Sprintf("items/%d-%d-%s.png", row, column, currentItem))
			// if outputErr != nil {
			// 	log.Fatal(outputErr)
			// }
			// png.Encode(outputFile, subImage)
		}
	}

	// "name": [groups][group][item(row,column)]
	itemGroups := map[string][][][2]int{}
	for item, _ := range items {
		itemGroups[item] = [][][2]int{}
	}

	var totalLength int
	var noGroups bool = true
	for itemRow, row := range currentState {
		for itemColumn, item := range row {
			if !strings.Contains(item, "done") && item != "" {

				newGroup := findSimilarNeighbours(&currentState, item, itemRow, itemColumn)

				totalLength += len(newGroup) + 1
				// add the newly found group to the list of groups
				// fmt.Println(" item : '" + item + "'\n")
				if len(newGroup) != 0 /*&& item != ""*/ {
					noGroups = false
					// add the current item to the group
					newGroup = append(newGroup, [2]int{itemRow, itemColumn})
					itemGroups[item] = append(itemGroups[item], newGroup)
				}

			}
		}
	}
	fmt.Printf("Current state : %v\n", currentState)
	fmt.Printf("Groups found : %v\n", itemGroups)
	fmt.Printf("Found groups for %d items\n", totalLength)

	return itemGroups, noGroups

}

func findSimilarNeighbours(currentState *[][]string, item string, itemRow int /*index*/, itemColumn int /*index*/) [][2]int {
	// fmt.Println("current: ", item)

	similarNeighbours := [][2]int{}

	if strings.Contains(item /*(*currentState)[itemRow][itemColumn]*/, "done") {
		return similarNeighbours
	}

	// above
	if itemRow-1 != -1 {
		neighbour := (*currentState)[itemRow-1][itemColumn]
		if neighbour == item {
			similarNeighbours = append(similarNeighbours, [2]int{itemRow - 1, itemColumn})
			(*currentState)[itemRow-1][itemColumn] += " (done)"
		}
	}

	// bottom
	if itemRow+1 != rows {
		neighbour := (*currentState)[itemRow+1][itemColumn]
		if neighbour == item {
			similarNeighbours = append(similarNeighbours, [2]int{itemRow + 1, itemColumn})
			(*currentState)[itemRow+1][itemColumn] += " (done)"
		}
	}

	// right
	if itemColumn+1 != columns {
		neighbour := (*currentState)[itemRow][itemColumn+1]
		if neighbour == item {
			similarNeighbours = append(similarNeighbours, [2]int{itemRow, itemColumn + 1})
			(*currentState)[itemRow][itemColumn+1] += " (done)"
		}
	}

	// left
	if itemColumn-1 != -1 {
		neighbour := (*currentState)[itemRow][itemColumn-1]
		if neighbour == item {
			similarNeighbours = append(similarNeighbours, [2]int{itemRow, itemColumn - 1})
			(*currentState)[itemRow][itemColumn-1] += " (done)"
		}
	}

	(*currentState)[itemRow][itemColumn] += " (done)"

	for _, neighbour := range similarNeighbours {
		// fmt.Println(neighbour)
		similarNeighbours = append(similarNeighbours, findSimilarNeighbours(currentState, item, neighbour[0], neighbour[1])...)
	}

	return similarNeighbours

}
