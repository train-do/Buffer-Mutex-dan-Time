package main

import (
	"fmt"
	"time"
)
func main() {
	sCh := make(chan string, 5)
	kCh := make(chan string, 5)
	tCh := make(chan string, 5)
	go sensorSensing(sCh, "\033[31mSuhu")
	go sensorSensing(kCh, "\033[33mKelembapan")
	go sensorSensing(tCh, "\033[32mTekanan")
	var arr [][]string
	for i := 0; i < 5; i++ {
		var temp []string
		suhu := <- sCh
		temp = append(temp, suhu)
		kelembapan := <- kCh
		temp = append(temp, kelembapan)
		tekanan := <- tCh
		temp = append(temp, tekanan)
		arr = append(arr, temp)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
	// fmt.Println(arr)
}
func sensorSensing(ch chan <- string, sensor string)  {
	ticker := time.NewTicker(100*time.Microsecond)
	layout := "03:04:05 PM"
	var start, end int
	switch sensor {
	case "\033[31mSuhu":
		start = 30
		end = 35
	case "\033[33mKelembapan":
		start = 40
		end = 45
	case "\033[32mTekanan":
		start = 0
		end = 5
	}
	for i := start; i < end; i++ {
		select {
		case t:= <-ticker.C:
			ch <- fmt.Sprintf("%s : %d \033[0m, Time : %s,", sensor, i, t.Format(layout))
		}
	}
	timeout := time.After(1 * time.Second)
	select {
	case <-timeout:
		ch <- "Timeout"
	}
	ticker.Stop()
	close(ch)
}