/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur'


* Channels: 
 */

package main

import (
	"fmt"
	"time"
)
 
 func main() {
	 c := make(chan string)
	 go count("sheep", c)
 
	 for msg := range c {
		 fmt.Println(msg)
	 }
 }
 
 func count(thing string, c chan string) {
	 for i := 1; ; i++ {
		 c <- fmt.Sprintf("%d %s", i, thing)
		 time.Sleep(time.Second)
	 }
 }
 