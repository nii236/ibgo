package main

import (
	"fmt"
	"sync"

	"github.com/gofinance/ib"
)

func main() {

	e, err := ib.NewEngine(ib.EngineOptions{
		Gateway: "127.0.0.1:4002",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	im, err := ib.NewInstrumentManager(e, ib.Contract{
		Symbol:       "USD",
		SecurityType: "CASH",
		Exchange:     "IDEALPRO",
		Currency:     "JPY",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	go HandleInstrument(im)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func HandleInstrument(im *ib.InstrumentManager) {
	for {
		select {
		case t := <-im.Refresh():
			if err := im.FatalError(); err != nil {
				fmt.Println(err)
				break
			}
			if t {
				// fmt.Println(im.Last())
			}

		}
	}
}

func HandleCurrentTime(ctm *ib.CurrentTimeManager) {
	for {
		select {
		case t := <-ctm.Refresh():
			if err := ctm.FatalError(); err != nil {
				fmt.Println(err)
				break
			}
			if t {
				fmt.Println(ctm.Time())
			}

		}
	}
}
