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

            var first = 0;
            for second := 0; second < len(*activeRocks); second++ {
                var tempRock rocks = (*activeRocks)[second];
                if tempRock.height < terminalHeight - 1 {
                    if tempRock.state != "O" {
                        tempRock.height++;
                        (*activeRocks)[first] = tempRock;
                        first++;
                    }
                }
            }
            *activeRocks = (*activeRocks)[:first];
            rockMutex.Unlock();
        }
        time.Sleep(200 * time.Millisecond);
    }
}


