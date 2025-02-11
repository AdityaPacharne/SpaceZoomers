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

type spaceship struct {
    health int
    xaxis int
    yaxis int
}

type Rocks struct {
	xaxis int
	yaxis int
}

type bullet struct {
    direction bool
	xaxis int
	yaxis int
}

func bulletExistence (screen *[][]string, b *bullet, terminalHeight int, terminalWidth int) bool {
	if b.yaxis == 2 {
		return false;
	}
	return true;
}

func bulletLocation (screen *[][]string, activeBullets *[]bullet) {
	for tempBullet := activeBullets {

func bulletCreate (screen *[][]string, currentHeight *int, currentWidth *int, activeBullets *[]bullet, terminalHeight int, terminalWidth int, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			newBullets := activeBullets[:0];
			for _, tempBullet := range activeBullets {
				var tempExistence bool = bulletExistence (screen, tempBullet, terminalHeight, terminalWidth);
				if !tempExistence {
					newBullets = append(newBullets, tempBullet);

			(*activeBullets) = (*activeBullets).append(bullet{existence: true, xaxis: *currentHeight - 1, yaxis: *currentWidth});
			(*screen)[*currentHeight - 1][*currentWidth] = "^";
			time.Sleep(100 * time.Millisecond);
		}
	}
}

func spaceship (screen *[][]string, currentHeight *int, currentWidth *int, direction string, terminalHeight int, terminalWidth int) {
	if direction == "k" {
		if *currentWidth < terminalWidth - 1 {
			(*screen)[*currentHeight][*currentWidth] = "_";
			(*screen)[*currentHeight][*currentWidth+1] = "H";
			*currentWidth += 1;
		}
	} else if direction == "j" {
		if *currentWidth > 0 {
			(*screen)[*currentHeight][*currentWidth] = "_";
			(*screen)[*currentHeight][*currentWidth-1] = "H";
			*currentWidth -= 1;
		}
	} else {
		fmt.Println("Invalid Input");
	}
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

func playerInput(screen *[][]string, *currentHeight int, *currentWidth int, terminalHeight int, terminalWidth int, quit chan bool) {
	fd := int(os.Stdin.Fd())
	oldState, _ := term.MakeRaw(fd) // Enable raw mode
	defer term.Restore(fd, oldState) // Restore on exit

	reader := bufio.NewReader(os.Stdin);
	for {
		input, _ := reader.ReadByte();
		if input == 27 { // Escape key
			// Check for the arrow keys
			next1, _ := reader.ReadByte()
			next2, _ := reader.ReadByte()
			if next1 == 91 { // '['
				if next2 == 67 {
					fmt.Println("Right Arrow")
				} else if next2 == 68 {
					fmt.Println("Left Arrow")
				}
			}
		} else if input == 'q' { // Quit on 'q'
			break
		} else {
			fmt.Printf("You typed: %c\n", input)
		}

		var inputString string = string(input);
		if inputString == "q" {
			quit <- true;
			break;
		}
		spaceship(screen, currentHeight, currentWidth, inputString, terminalHeight, terminalWidth);
	}
}
			
func getTerminalSize() (terminalWidth, terminalHeight int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Col), int(ws.Row)
}

func main() {
    // Player Enters
	fmt.Println("Welcome Player");
	fmt.Println("Enter q to exit");

    // We store the players stty settings and make some changes but we will restore the user's settings when they quit
	player_stty_settings, _ := exec.Command("stty", "-g").Output()

    // stty settings to take only one element as input and to show it to the user
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run();
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run();

    // Fetching the terminal size
	var terminalWidth, terminalHeight int = getTerminalSize();

    // Creating a nested slice of strings to mimic a screen
	screen := make([][]string, terminalHeight);
	for i := range screen {
    	screen[i] = make([]string, terminalWidth)
		for j := range screen[i] {
			screen[i][j] = "_";
		}
	}
    
    // Placing our spaceship as a single character for now at the centre of the screen
    screen[terminalHeight-1][terminalWidth/2] = "H";
	var currentHeight int = terminalHeight-1;
	var currentWidth int = terminalWidth/2;

    // Creating a slice of type bullet struct to save the activeBullets present on Users screen
	var activeBullets []bullet;

    // Bool Channel to check if user has to quit
	quit := make(chan bool);
    spaceshipDirection := make(chan string);
	
    // A waitgroup to synchronize multiple goroutines, it waits for 3 goroutines to finish 
	var wg sync.WaitGroup
	wg.Add(3);

    // Goroutine to Print the screen
	go func() {
		defer wg.Done();
		render(&screen, &activeBullets, quit);
	}();

    // Goroutine to take player input and modify the position of spaceship accordingly
	go func() {
		defer wg.Done();
		playerInput(&screen, &currentHeight, &currentWidth, terminalHeight, terminalWidth, quit);
	}();

    // Go routine to create bullet every specific duration
    // It is made so the user dont have to shoot, the spaceship will keep firing and the user has to aim
	go func() {
		defer wg.Done();
		bulletCreate(&screen, &currentHeight, &currentWidth, &activeBullets, terminalHeight, quit)
	}();

    // WaitGroup waits....
	wg.Wait();

    // We restore the stty settings of player
	exec.Command("stty", string(player_stty_settings)).Run();
	fmt.Println("Your stty settings have been restored");

    // A little thanks for playing
    // Thanks for playing my game
    // And thanks to you too, to take time from your day to read my dumb code
	fmt.Println("Thanks for playing the game");
}
