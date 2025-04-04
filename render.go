package main

import (
    "fmt"
    "strings"
    "time"
)

func checkOutOfBound(terminalWidth int, spaceship *spaceshipstruct, spaceshipDirection string) int {
    if spaceshipDirection == "right" && spaceship.width < terminalWidth-3 {
        (*spaceship).width++;
    } else if spaceshipDirection == "left" && spaceship.width > 2 {
        (*spaceship).width--;
    }
    return (*spaceship).width;
}

func Render(actualScreen [][]string, activeBullets *[]bullet, activeRocks *[]rocks, terminalWidth int, spaceship *spaceshipstruct, spaceshipDirection chan string, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            screen := make([][]string, len(actualScreen))
            for i := range actualScreen {
                screen[i] = append([]string(nil), actualScreen[i]...)
            }

			fmt.Print("\033[H")

            for _, tempBullet := range *activeBullets {
                screen[tempBullet.height][tempBullet.width] = ":";
            }

            for _, tempRock := range *activeRocks {
                screen[tempRock.height][tempRock.width] = tempRock.state;
            }

            select {
            case dir := <- spaceshipDirection:
                var newCurrentWidth int = checkOutOfBound(terminalWidth, spaceship, dir);
                screen[spaceship.height][newCurrentWidth-2] = "/";
                screen[spaceship.height][newCurrentWidth-1] = "-";
                screen[spaceship.height][newCurrentWidth] = "o";
                screen[spaceship.height][newCurrentWidth+1] = "-";
                screen[spaceship.height][newCurrentWidth+2] = "\\";
            default:
                screen[spaceship.height][spaceship.width-2] = "/";
                screen[spaceship.height][spaceship.width-1] = "-";
                screen[spaceship.height][spaceship.width] = "o";
                screen[spaceship.height][spaceship.width+1] = "-";
                screen[spaceship.height][spaceship.width+2] = "\\";
            }

			var screenBuffer strings.Builder;
			for _, row := range screen {
				screenBuffer.WriteString(strings.Join(row, "") + "\n");
			}
			fmt.Print(screenBuffer.String());

			time.Sleep(10 * time.Millisecond);
		}
	}
}

