package main

import (
    "time"
)

func abs(x int, y int) int {
    if x > y {
        return x - y;
    }
    return y - x;
}

func BulletLocation (activeBullets *[]bullet, activeRocks *[]rocks, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return
        default:
            bulletMutex.Lock();
            rockMutex.Lock();
            var first int = 0;
            for second:=0; second < len(*activeBullets); second++ {
                var tempBullet bullet = (*activeBullets)[second];

                if tempBullet.direction && tempBullet.height >= 5 {
                    var collided bool = false;
                    for i := range *activeRocks {
                        var tempRock *rocks = &(*activeRocks)[i]
                        if tempRock.width == tempBullet.width && abs(tempRock.height,tempBullet.height) <= 1 {
                            tempRock.state = "O";
                            collided = true;
                            break;
                        }
                    }
                    if !collided {
                        tempBullet.height--;
                        (*activeBullets)[first] = tempBullet;
                        first++;
                    }
                }
            }
            *activeBullets = (*activeBullets)[:first];

            rockMutex.Unlock();
            bulletMutex.Unlock();
        }
        time.Sleep(100 * time.Millisecond);
    }
}

func BulletCreate (activeBullets *[]bullet, spaceship *spaceshipstruct, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            spaceshipMutex.Lock();
            var newBullet bullet = bullet{
                direction: true,
                height: (*spaceship).height - 1,
                width: (*spaceship).width,
            }
            spaceshipMutex.Unlock();

            bulletMutex.Lock();
            (*activeBullets) = append((*activeBullets), newBullet);
            bulletMutex.Unlock();
			time.Sleep(200 * time.Millisecond);
		}
	}
}
