package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	// sessionID := "B32CE4CC2D11AD00611839E0F631F0DD"

	// cookies := []*proto.NetworkCookieParam{}
	// cookies = append(cookies, &proto.NetworkCookieParam{Name: "JSESSIONID", Value: sessionID})

	l := InstanciateBrowser()
	browser := rod.New().ControlURL(l).MustConnect()

	// browser.SetCookies(cookies)

	page := browser.MustPage("https://gosupermodel.com/games/wardrobegame.jsp")
	page.MustWaitLoad()

	//login
	loginFields := page.MustElements("input.loginpage_loginbox_input")
	loginFields[0].MustClick().MustSelectAllText().MustInput(username)
	loginFields[1].MustClick().MustSelectAllText().MustInput(password)
	page.MustElement("a.button.coloridx2.floatLeft").MustClick()

	page.MustWaitLoad()

	pinball, _ := page.Elements(".pinball_frame")
	if len(pinball) > 0 {
		fmt.Println("Pinball is available, please play it and re-run the program")
		return
	}

	playLevel(page)

	time.Sleep(time.Hour)
	browser.Close()
}
