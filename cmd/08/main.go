/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur'


* Goroutines
 */

package main

import (
	"fmt"
	"sync"
	"time"
)
 
 var wg = sync.WaitGroup{}
 var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
 var results = make([]string, len(dbData))
 
 func main() {
	 tB := time.Now()
	 for i := 0; i < len(dbData); i++ {
		 wg.Add(1)
		 go dbCall(i)
	 }
	 wg.Wait()
	 fmt.Printf("\nTotal execution time: %v", time.Since(tB))
	 fmt.Printf("\nThe results: %v", results)
 }
 
 func dbCall(i int) {
	 var delay = 2000
	 time.Sleep(time.Duration(delay) * time.Millisecond)
	 fmt.Println("The result from the databases is: ", dbData[i])
	 results[i] = dbData[i]
	 wg.Done()
 }
 