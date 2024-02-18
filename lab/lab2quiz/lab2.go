package main

import (
	"log"
	"time"
)

func main() {
	log.Println("I am going to write an interesting program!")
	go func() {
		go func() {
			go func() {
				go func() {
					go func() {
						go func() {
							go func() {

								log.Println("well.. maybe I shuold have thought about this program a little more")
							}()
						}()
					}()
				}()
			}()
		}()
	}()
	time.Sleep(8 * time.Second)
}
