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

            bulletMutex.Lock();
            for _, tempBullet := range *activeBullets {
                screen[tempBullet.height][tempBullet.width] = ":";
            }
            bulletMutex.Unlock();

            rockMutex.Lock();
            for _, tempRock := range *activeRocks {
                if tempRock.height < len(screen) && tempRock.width < len(screen[0]) {
                    screen[tempRock.height][tempRock.width] = tempRock.state;
                }
            }
            rockMutex.Unlock();

            var newCurrentWidth int;
            select {
            case dir := <- spaceshipDirection:
                spaceshipMutex.Lock();
                newCurrentWidth = checkOutOfBound(terminalWidth, spaceship, dir);
                spaceshipMutex.Unlock();
            default:
                spaceshipMutex.Lock();
                newCurrentWidth = spaceship.width;
                spaceshipMutex.Unlock();
            }
            
            screen[spaceship.height][newCurrentWidth-2] = "/";
            screen[spaceship.height][newCurrentWidth-1] = "-";
            screen[spaceship.height][newCurrentWidth] = "o";
            screen[spaceship.height][newCurrentWidth+1] = "-";
            screen[spaceship.height][newCurrentWidth+2] = "\\";

			var screenBuffer strings.Builder;
			for _, row := range screen {
				screenBuffer.WriteString(strings.Join(row, "") + "\n");
			}
			fmt.Print(screenBuffer.String());

			time.Sleep(33 * time.Millisecond);
		}
	}
}

