package main

import (
    "fmt"
    "strings"
    "time"
)

// Function that renders the screen
func checkOutOfBound(terminalWidth int, spaceship *spaceshipstruct, spaceshipDirection string) int {
    fmt.Println("Inside checker");
    if spaceshipDirection == "right" && spaceship.width < terminalWidth-1 {
        (*spaceship).width++;
        return (*spaceship).width + 1;
    } else if spaceshipDirection == "left" && spaceship.width > 0 {
        (*spaceship).width--;
        return (*spaceship).width - 1;
    } else {
        return (*spaceship).width;
    }
}

func Render(actualScreen [][]string, activeBullets *[]bullet, terminalWidth int, spaceship *spaceshipstruct, spaceshipDirection chan string, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            screen := make([][]string, len(actualScreen))
            for i := range actualScreen {
                screen[i] = append([]string(nil), actualScreen[i]...) // Deep copy each row
            }

			fmt.Print("\033[H\033[2J")

            // Adds bullet onto the screen from activeBullets slice
            for _, tempBullet := range *activeBullets {
                screen[tempBullet.height][tempBullet.width] = "^";
            }

            select {
            case dir := <- spaceshipDirection:
                var newCurrentWidth int = checkOutOfBound(terminalWidth, spaceship, dir);
                screen[spaceship.height][newCurrentWidth] = "H";
            default:
                screen[spaceship.height][spaceship.width] = "H";
            }

            // Printing the screen
			for _, row := range (screen) {
				fmt.Println(strings.Join(row, ""))
			}

			time.Sleep(50 * time.Millisecond);
		}
	}
}

