package main

import (
	"fmt"
	"time"
)

func main() {
	s := Spinner{}
	s.Initialize()

	for lcv := 0; lcv < 50; lcv++ {
		s.Next()
		time.Sleep(time.Millisecond * 50)
	}
	s.Done()

	fmt.Printf("Done\n")
}

type Spinner struct {
	Lcv        int
	Iterations []string
}

func (s *Spinner) Initialize() {
	s.Lcv = 0
	s.Iterations = []string{
		"-",
		"\\",
		"|",
		"/",
	}
}

func (s *Spinner) Next() {
	iteration := s.Iterations[s.Lcv]

	s.Lcv++
	s.Lcv %= 4

	fmt.Printf("\r%s", iteration)
}

func (s *Spinner) Done() {
	fmt.Print("\r \r")
}
