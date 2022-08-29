package greeting

import (
	"errors"
	"fmt"
	"sync"
)

func New(wg *sync.WaitGroup) (*Greetings, error) {
	if wg == nil {
		return nil, errors.New("nil waitgroup")
	}
	return &Greetings{wg: wg}, nil
}

type Greetings struct {
	wg *sync.WaitGroup // This is unexported so only through new can be intantiated
}

// There is waitgroup so use wg.Add(1) before calling waitgroup
func (g *Greetings) Greet(msg string) {
	fmt.Println(msg)
	g.wg.Done()
}
