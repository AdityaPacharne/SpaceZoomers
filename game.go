package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"golang.org/x/sys/unix"
)

func spaceship (spaceshipPosition *[]string, current int, direction string) int {
	if direction == "k" {
		(*spaceshipPosition)[current] = "";
		(*spaceshipPosition)[current+1] = "H";
		return current + 1;
	} else if direction == "j" {
		(*spaceshipPosition)[current] = "";
		(*spaceshipPosition)[current-1] = "H";
		return current - 1;
	}
	return current;
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
	var current int = 5;

	var width, height int = getTerminalSize();

	screen := make([][]string, height)
	for i := range screen {
    		screen[i] = make([]string, width)
	}

	for {
		render(screen);
		input, _ := reader.ReadByte();
		var inputString string = string(input);
		if inputString == "q" {
			break;
		}
		current = spaceship(&spaceshipPosition, current, inputString);
	}
	exec.Command("stty", string(player_stty_settings)).Run();
	fmt.Println("Your stty settings have been restored");
	fmt.Println("Thanks for playing the game");
}
