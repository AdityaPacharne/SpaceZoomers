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
            rockMutex.Lock();
            var newRock rocks = rocks{
                height: 0,
                width: tempRockWidth,
                state: "*",
            }
            (*activeRocks) = append((*activeRocks), newRock);
            rockMutex.Unlock();
        }
        time.Sleep(400 * time.Millisecond);
    }
}

func RocksLocation(activeRocks *[]rocks, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return;
        default:
            rockMutex.Lock();
            var newRocks []rocks;
            for i := range len(*activeRocks) {
                if (*activeRocks)[i].height < terminalHeight - 1 {
                    (*activeRocks)[i].height++;
                    newRocks = append(newRocks, (*activeRocks)[i]);
                }
            }
            *activeRocks = newRocks;
            rockMutex.Unlock();
        }
        time.Sleep(150 * time.Millisecond);
    }
}


