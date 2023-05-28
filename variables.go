package main

import "image"

var username string = "komas44486"
var password string = "azerty123"

var currentStepScreenshot image.Image

// values in px
const rows int = 12
const columns int = 11
const itemSize int = 33
const itemSpacing int = 1
const gridMarginFromCanvasTop int = 19
const gridMarginFromCanvasLeft int = 357

// const canvasMarginFromPageTop int = 182
// const canvasMarginFromPageLeft int = 143

// css selectors
const gameCanvas string = "#gsmbanneriframe #canvas canvas"
