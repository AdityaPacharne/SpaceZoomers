package main

import (
    "os"
    "bufio"
    "golang.org/x/term"
    "time"
)

func PlayerInput(spaceshipDirection chan string, currentHeight *int, currentWidth *int, terminalHeight int, terminalWidth int, quit chan bool) {
    fd := int(os.Stdin.Fd());
    oldState, _ := term.MakeRaw(fd);
    defer term.Restore(fd, oldState);

    reader := bufio.NewReader(os.Stdin);
    for {
        input, _ := reader.ReadByte();
        if input == 27 {
            next1, _ := reader.ReadByte();
            next2, _ := reader.ReadByte();
            if next1 == 91 {
                if next2 == 67 {
                    (*currentWidth)++;
                    spaceshipDirection <- "right";
                } else if next2 == 68 {
                    (*currentWidth)--;
                    spaceshipDirection <- "left";
                }
            }
        } else if input == 'q' {
            for i:=0; i<3; i++ {
                quit <- true;
            }
            break;
        }
        time.Sleep(100 * time.Millisecond);
    }
}
