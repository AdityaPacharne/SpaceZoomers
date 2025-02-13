package main

import (
    "fmt"
    "strings"
    "time"
)

// Function that renders the screen
func checkOutOfBound(terminalWidth int, copyOfCurrentWidth int, spaceshipDirection string) int {
    if spaceshipDirection == "right" && copyOfCurrentWidth < terminalWidth-1 {
        return copyOfCurrentWidth + 1;
    } else if spaceshipDirection == "left" && copyOfCurrentWidth > 0 {
        return copyOfCurrentWidth - 1;
    } else {
        return copyOfCurrentWidth;
    }
}

func Render(screen [][]string, activeBullets *[]bullet, terminalWidth int, currentHeight int, currentWidth int, spaceshipDirection chan string, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			fmt.Print("\033[H\033[2J")

            // Adds bullet onto the screen from activeBullets slice
            for _, tempBullet := range *activeBullets {
                screen[tempBullet.xaxis][tempBullet.yaxis] = "^";
            }

            select {
            case dir := <- spaceshipDirection:
                var newCurrentWidth int= checkOutOfBound(terminalWidth, currentWidth, dir);
                screen[currentHeight][newCurrentWidth] = "H";
            default:
                    screen[currentHeight][currentWidth] = "H";
            }

            // Printint the screen
			for _, row := range (screen) {
				fmt.Println(strings.Join(row, ""))
			}

			time.Sleep(10 * time.Millisecond);
		}
	}
}

