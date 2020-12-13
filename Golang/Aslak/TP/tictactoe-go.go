package main

import "github.com/tadvi/winc"

var cases []*winc.PushButton
var player string

func main() {
	player = "X"
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 400)
	mainWindow.SetText("Morpion By Aslak")
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	initializeButtons(mainWindow)
	winc.RunMainLoop()
}

func initializeButtons(mainWindow *winc.Form) {
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			button := winc.NewPushButton(mainWindow)
			button.SetSize(50, 50)
			button.SetText("")
			button.SetPos(20+y*50, 20+i*50)
			button.OnClick().Bind(func(arg *winc.Event) {
				if button.Text() != "X" && button.Text() != "O" {
					button.SetText(player)
					end := checkEnd()
					if end {
						gameEnded(end, player)
					} else {
						end = checkDraw()
						if !end {
							gameEnded(end, player)
						}
					}
					if player == "X" {
						player = "O"
					} else {
						player = "X"
					}
				}
			})
			cases = append(cases, button)
		}
	}
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func checkEnd() bool {
	if cases[0].Text() == player && cases[1].Text() == player && cases[2].Text() == player {
		return true
	} else if cases[3].Text() == player && cases[4].Text() == player && cases[5].Text() == player {
		return true
	} else if cases[6].Text() == player && cases[7].Text() == player && cases[8].Text() == player {
		return true
	} else if cases[0].Text() == player && cases[3].Text() == player && cases[6].Text() == player {
		return true
	} else if cases[1].Text() == player && cases[4].Text() == player && cases[7].Text() == player {
		return true
	} else if cases[2].Text() == player && cases[5].Text() == player && cases[8].Text() == player {
		return true
	} else if cases[0].Text() == player && cases[4].Text() == player && cases[8].Text() == player {
		return true
	} else if cases[2].Text() == player && cases[4].Text() == player && cases[6].Text() == player {
		return true
	}
	return false
}

func checkDraw() bool {
	if cases[0].Text() != "" && cases[1].Text() != "" && cases[2].Text() != "" && cases[3].Text() != "" && cases[4].Text() != "" && cases[5].Text() != "" && cases[6].Text() != "" && cases[7].Text() != "" && cases[8].Text() != "" {
		return false
	}
	return true
}

func gameEnded(isWon bool, player string) {
	winWindow := winc.NewForm(nil)
	winWindow.SetSize(400, 400)
	if isWon {
		winWindow.SetText("Morpion By Aslak - Victoire")
		label := winc.NewLabel(winWindow)
		label.SetPos(10, 10)
		label.SetText("Victoire du joueur " + player)
	} else {
		winWindow.SetText("Morpion By Aslak- Égalité")
		label := winc.NewLabel(winWindow)
		label.SetPos(10, 10)
		label.SetSize(200, 200)
		label.SetText("Égalité : pas de gagnant...")
	}
	winWindow.Show()
	winWindow.OnClose().Bind(wndOnClose)
}
