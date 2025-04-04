package main

import (
    "fmt"
	"golang.org/x/sys/unix"
	"sync"
)

var bulletMutex sync.Mutex;
var rockMutex sync.Mutex;
var spaceshipMutex sync.Mutex;

type spaceshipstruct struct {
    health int
    height int
    width int
}

type rocks struct {
    state string 
	height int
	width int
}

type bullet struct {
    direction bool
	height int
	width int
}

// Code to fetch terminal size
func getTerminalSize() (terminalHeight, terminalWidth int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Row), int(ws.Col)
}

func main() {
    // Fetching the terminal size
	var terminalHeight, terminalWidth int = getTerminalSize();

    // Creating a nested slice of strings to perfectly mimic user's terminal as screen
	screen := make([][]string, terminalHeight);
	for i := range screen {
    	screen[i] = make([]string, terminalWidth)
		for j := range screen[i] {
			screen[i][j] = " ";
		}
	}
    
    // Placing our spaceship as a single character for now at the centre of the screen
    var spaceship spaceshipstruct;
    spaceship.health = 100;
    spaceship.height = terminalHeight - 1;
    spaceship.width = terminalWidth / 2;

    // Creating a slice of type bullet struct to save the activeBullets present on Users screen
	var activeBullets []bullet;

    // Creating a slice of type rock struct to save the activeRocks present on Users screen
    var activeRocks []rocks;
    
    // Bool Channel to check if user has to quit
	quit := make(chan bool);
    spaceshipDirection := make(chan string);
	
    // A waitgroup to synchronize multiple goroutines, it waits for 3 goroutines to finish 
	var wg sync.WaitGroup
	wg.Add(5);

    // Goroutine to Print the screen
	go func() {
		defer wg.Done();
		Render(screen, &activeBullets, &activeRocks, terminalWidth, &spaceship, spaceshipDirection, quit);
	}();

    // Goroutine to take player input and modify the position of spaceship accordingly
	go func() {
		defer wg.Done();
		PlayerInput(spaceshipDirection, &spaceship, terminalHeight, terminalWidth, quit);
	}();

    // Go routine to create bullet every specific duration
    // It is made so the user dont have to shoot, the spaceship will keep firing and the user has to aim
	go func() {
		defer wg.Done();
		BulletCreate(&activeBullets, &spaceship, true, quit);
	}();

	go func() {
		defer wg.Done();
        BulletLocation(&activeBullets, &activeRocks, terminalHeight, quit);
	}();

    go func() {
        defer wg.Done();
        RocksCreate(&activeRocks, terminalHeight, terminalWidth, quit);
    }();

    go func() {
        defer wg.Done();
        RocksLocation(&activeRocks, terminalHeight, quit);
    }();

    // WaitGroup waits....
	wg.Wait();

    // Thanks for playing my game
    // And thanks to take time from your day to read my dumb code

    fmt.Println("Thanks for playing the game");
}
