package main

import (
    "time"
)

func BulletLocation (activeBullets *[]bullet, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return
        default:
            var newBullets []bullet;
            for i := range *activeBullets {
                if ((*activeBullets)[i].direction && (*activeBullets)[i].yaxis >= 2) || (!(*activeBullets)[i].direction && (*activeBullets)[i].yaxis <= terminalHeight - 2) {  
                    newBullets = append(newBullets, (*activeBullets)[i]);
                }
            }
            *activeBullets = newBullets;
        }
        time.Sleep(10 * time.Millisecond);
    }
}

// func BulletLocation (activeBullets *[]bullet, terminalHeight int, quit chan bool) {
//     for {
//         select {
//         case <- quit:
//             return
//         default:
//             for i := range activeBullets {
//                 var answer bool = bulletExistence((*activeBullet)[i], terminalHeight);
//             }
//         }
//     }
// }

func BulletCreate (activeBullets *[]bullet, currentHeight *int, currentWidth *int, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            if spaceshipBullet {
                (*activeBullets) = append((*activeBullets), bullet{direction: true, xaxis: *currentHeight - 1, yaxis: *currentWidth});
            } else {
                (*activeBullets) = append((*activeBullets), bullet{direction: false , xaxis: *currentHeight - 1, yaxis: *currentWidth});
            }
			time.Sleep(100 * time.Millisecond);
		}
	}
}
