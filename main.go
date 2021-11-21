package main

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"
)

var input string = "tester"
var hash = md5.Sum([]byte(input))
var length int32 = int32(len(input))
var lowerLimit int32 = 97
var upperLimit int32 = 127
var wg sync.WaitGroup

func stringEnumerator(depth int32, runeSlice []rune) {
	if depth < length {
		for currentCharacter := lowerLimit; currentCharacter <= upperLimit; currentCharacter++ {
			runeSlice[depth] = currentCharacter
			stringEnumerator(depth + 1, runeSlice)
		}
	} else {
		if md5.Sum([]byte(string(runeSlice))) == hash {
			println("Hash Match: ", string(runeSlice))
		}
	}
}

func stringStarter(startingCharacter rune) {
	defer wg.Done()
	runeSlice := make([]rune, length)
	runeSlice[0] = startingCharacter
	stringEnumerator(1, runeSlice)
}

func calculateHashingDomain() {
	for startingCharacter := lowerLimit; startingCharacter <= upperLimit; startingCharacter++ {
		wg.Add(1)
		go stringStarter(startingCharacter)
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	calculateHashingDomain()
	duration := time.Since(start)

	fmt.Printf("Execution Time: %s", duration.Seconds())
}
