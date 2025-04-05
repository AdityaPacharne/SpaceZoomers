package main

import (
    "os"
    "bufio"
    "golang.org/x/term"
)

func PlayerInput(spaceshipDirection chan string, spaceship *spaceshipstruct, terminalHeight int, terminalWidth int, quit chan bool) {
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
                    spaceshipDirection <- "right";
                } else if next2 == 68 {
                    spaceshipDirection <- "left";
                }
            }
        } else if input == 'q' {
            for i:=0; i<5; i++ {
                quit <- true;
            }
            break;
        }
    }
}
