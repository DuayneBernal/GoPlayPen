package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	kr := KnightRider{}
	kr.Initialize()

	for lcv := 0; lcv < 50; lcv++ {
		kr.Next()
		time.Sleep(time.Millisecond * 50)
	}
	kr.Done()

	fmt.Print("\rDone\n")
}

type KnightRider struct {
	Character  string
	Lcv        int
	Buffer     [17]string
}

func (kr *KnightRider) Initialize() {
	kr.Character = "."
	kr.Lcv = 0
}

func (kr *KnightRider) Next() {
	for lcv := 0; lcv < len(kr.Buffer); lcv++ {
		kr.Buffer[lcv] = " "
	}

	offset := kr.Lcv
	if kr.Lcv > 12 {
		offset = 24 - kr.Lcv
	}

	for lcv := 0; lcv < 5; lcv++ {
		kr.Buffer[offset + lcv] = kr.Character
	}

	kr.Lcv++
	kr.Lcv %= 24

	fmt.Printf("%s\r", strings.Join(kr.Buffer[4:13], ""))
}

func (kr *KnightRider) Done() {
	fmt.Printf("\r%s\r", strings.Repeat(" ", 9))
}
