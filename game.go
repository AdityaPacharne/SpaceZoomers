package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"golang.org/x/sys/unix"
)

func spaceship (spaceshipPosition *[][]string, currentHeight int, currentWidth int, direction string) int {
	if direction == "k" {
		(*spaceshipPosition)[currentHeight][currentWidth] = ".";
		(*spaceshipPosition)[currentHeight][currentWidth+1] = "H";
		return currentWidth + 1;
	} else if direction == "j" {
		(*spaceshipPosition)[currentHeight][currentWidth] = ".";
		(*spaceshipPosition)[currentHeight][currentWidth-1] = "H";
		return currentWidth - 1;
	} else {
		fmt.Println("Invalid Input");
	}
	return currentWidth;
}

func render(screen [][]string) {
    fmt.Print("\033[H\033[2J") // Clear screen
    for _, row := range screen {
        fmt.Println(strings.Join(row, ""))
    }
}

func getTerminalSize() (width, height int) {
	// var ws unix.Winsize
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Col), int(ws.Row)
}
			
func main() {
	fmt.Println("Welcome Player");
	fmt.Println("Enter q to exit");

	player_stty_settings, _ := exec.Command("stty", "-g").Output()
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run();
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run();

	reader := bufio.NewReader(os.Stdin);

	var width, height int = getTerminalSize();
	fmt.Println("Width: ", width);
	fmt.Println("Height: ", height);

	screen := make([][]string, height);
	for i := range screen {
    		screen[i] = make([]string, width)
		for j := range screen[i] {
			screen[i][j] = ".";
		}
	}
	screen[height-1][width/2] = "H";
	var currentHeight int = height-1;
	var currentWidth int = width/2;
	fmt.Println("Spaceship Height: ", currentHeight);
	fmt.Println("Spaceship Width: ", currentWidth);

	for {
		render(screen);
		input, _ := reader.ReadByte();
		var inputString string = string(input);
		if inputString == "q" {
			break;
		}
		currentWidth = spaceship(&screen, currentHeight, currentWidth, inputString);
	}
	exec.Command("stty", string(player_stty_settings)).Run();
	fmt.Println("Your stty settings have been restored");
	fmt.Println("Thanks for playing the game");
}
