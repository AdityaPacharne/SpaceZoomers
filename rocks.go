package main

import (
    "math/rand/v2"
    "time"
)

func RocksCreate(activeRocks *[]rocks, terminalHeight int, terminalWidth int, quit chan bool) {
    for {
        select {
        case <- quit:
            return;
        default:
            var tempRockWidth int = rand.IntN(terminalWidth);
            (*activeRocks) = append((*activeRocks), rocks{height: 0, width: tempRockWidth, state: "*"});
        }
        time.Sleep(3 * time.Second);
    }
}

func RocksLocation(activeRocks *[]rocks, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return;
        default:
            var newRocks []rocks;
            for i := range len(*activeRocks) {
                if (*activeRocks)[i].height < terminalHeight - 1 {
                    (*activeRocks)[i].height++;
                    newRocks = append(newRocks, (*activeRocks)[i]);
                }
            }
            *activeRocks = newRocks;
        }
        time.Sleep(1 * time.Second);
    }
}







