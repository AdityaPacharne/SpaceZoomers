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

type Bullet struct {
	existence bool
	xaxis int
	yaxis int
}

func (r *Bullet) bulletLocation(screen *[][]string, terminalHeight int, terminalWidth int) {
	var bulletX int = r.xaxis;
	var bulletY int = r.yaxis;
	if bulletY == 0 {
		r.existence = false;
	}
}

func bulletCreate (screen *[][]string, currentHeight int, currentWidth int, activeBullets *[]Bullet, quit) {


func spaceship (screen *[][]string, currentHeight int, currentWidth int, direction string, terminalHeight int, terminalWidth int) int {
	if direction == "k" {
		if currentWidth < terminalWidth - 1 {
			(*screen)[currentHeight][currentWidth] = "_";
			(*screen)[currentHeight][currentWidth+1] = "H";
			return currentWidth + 1;
		}
	} else if direction == "j" {
		if currentWidth > 0 {
			(*screen)[currentHeight][currentWidth] = "_";
			(*screen)[currentHeight][currentWidth-1] = "H";
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

func getTerminalSize() (terminalWidth, terminalHeight int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Col), int(ws.Row)
}

func playerInput(screen *[][]string, currentHeight int, currentWidth int, terminalHeight int, terminalWidth int, quit chan bool) {
	reader := bufio.NewReader(os.Stdin);
	for {
		input, _ := reader.ReadByte();
		var inputString string = string(input);
		if inputString == "q" {
			quit <- true;
			break;
		}
		currentWidth = spaceship(screen, currentHeight, currentWidth, inputString, terminalHeight, terminalWidth);
	}
}
			
func main() {
	fmt.Println("Welcome Player");
	fmt.Println("Enter q to exit");

	player_stty_settings, _ := exec.Command("stty", "-g").Output()
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run();
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run();


	var terminalWidth, terminalHeight int = getTerminalSize();
	fmt.Println("Width: ", terminalWidth);
	fmt.Println("Height: ", terminalHeight);

	screen := make([][]string, terminalHeight);
	for i := range screen {
    		screen[i] = make([]string, terminalWidth)
		for j := range screen[i] {
			screen[i][j] = "_";
		}
	}
	screen[terminalHeight-1][terminalWidth/2] = "H";
	var currentHeight int = terminalHeight-1;
	var currentWidth int = terminalWidth/2;

	quit := make(chan bool);
	
	var wg sync.WaitGroup
	wg.Add(3);

	go func() {
		defer wg.Done();
		render(&screen, quit);
	}();

	go func() {
		defer wg.Done();
		playerInput(&screen, currentHeight, currentWidth, terminalHeight, terminalWidth, quit);
	}();

	go func() {
		defer wg.Done();
		bulletCreate(&screen, currentHeight, currentWidth, quit)
	}();

	wg.Wait();

	exec.Command("stty", string(player_stty_settings)).Run();
	fmt.Println("Your stty settings have been restored");
	fmt.Println("Thanks for playing the game");
}
