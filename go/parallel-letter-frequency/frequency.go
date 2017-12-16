/*
Package letter implements sequential and parallel
letter frequency computation.
*/
package letter

import (
	"sync"
	"time"
)

// FreqMap used to store individual letter frequency
type FreqMap map[rune]int

// Frequency used to sequentially count letter frequency
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency used to concurrently count letter frequency
func ConcurrentFrequency(s []string) FreqMap {
	mutex := &sync.Mutex{}
	c := 0
	m := FreqMap{}
	for _, x := range s {
		go func(m FreqMap, s string, c *int) {
			for l, freq := range Frequency(s) {
				mutex.Lock()
				m[l] += freq
				mutex.Unlock()
			}
			*c++
		}(m, x, &c)
	}
	for {
		if c == len(s) {
			return m
		}
		time.Sleep(time.Millisecond * 10)
	}
}
