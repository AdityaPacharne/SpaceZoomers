package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"golang.org/x/sys/unix"
	"time"
	"sync"
)

type Rocks struct {
	existence bool
	xaxis int
	yaxis int
}

type bullet struct {
	existence bool
	xaxis int
	yaxis int
}

func spaceship (spaceshipPosition *[][]string, currentHeight int, currentWidth int, direction string, terminalHeight int, terminalWidth int) int {
	if direction == "k" {
		if currentWidth < terminalWidth - 1 {
			(*spaceshipPosition)[currentHeight][currentWidth] = "_";
			(*spaceshipPosition)[currentHeight][currentWidth+1] = "H";
			return currentWidth + 1;
		}
	} else if direction == "j" {
		if currentWidth > 0 {
			(*spaceshipPosition)[currentHeight][currentWidth] = "_";
			(*spaceshipPosition)[currentHeight][currentWidth-1] = "H";
			return currentWidth - 1;
		}
	} else {
		fmt.Println("Invalid Input");
	}
	return currentWidth;
}

func render(screen *[][]string, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			fmt.Print("\033[H\033[2J")
			for _, row := range (*screen) {
				fmt.Println(strings.Join(row, ""))
			}
			time.Sleep(100 * time.Millisecond);
		}
	}
}

func getTerminalSize() (width, height int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Col), int(ws.Row)
}

func playerInput(screen *[][]string, currentHeight int, currentWidth int, height int, width int, quit chan bool) {
	reader := bufio.NewReader(os.Stdin);
	for {
		input, _ := reader.ReadByte();
		var inputString string = string(input);
		if inputString == "q" {
			quit <- true;
			break;
		}
		currentWidth = spaceship(screen, currentHeight, currentWidth, inputString, height, width);
	}
}
			
func main() {
	fmt.Println("Welcome Player");
	fmt.Println("Enter q to exit");

	player_stty_settings, _ := exec.Command("stty", "-g").Output()
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run();
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run();


	var width, height int = getTerminalSize();
	fmt.Println("Width: ", width);
	fmt.Println("Height: ", height);

	screen := make([][]string, height);
	for i := range screen {
    		screen[i] = make([]string, width)
		for j := range screen[i] {
			screen[i][j] = "_";
		}
	}
	screen[height-1][width/2] = "H";
	var currentHeight int = height-1;
	var currentWidth int = width/2;

	quit := make(chan bool);
	
	var wg sync.WaitGroup
	wg.Add(2);

	go func() {
		defer wg.Done();
		render(&screen, quit);
	}();

	go func() {
		defer wg.Done();
		playerInput(&screen, currentHeight, currentWidth, height, width, quit);
	}();

	wg.Wait();

	exec.Command("stty", string(player_stty_settings)).Run();
	fmt.Println("Your stty settings have been restored");
	fmt.Println("Thanks for playing the game");
}
