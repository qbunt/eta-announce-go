package ifttt

import (
	"github.com/jamesmillerio/go-ifttt-maker"
	"os"
)

func Notify(timeInTraffic string, eta string) {
	maker := new(GoIFTTTMaker.MakerChannel)
	key, event := os.Getenv("MAKER_KEY"), os.Getenv("MAKER_EVENT")

	maker.Value1 = timeInTraffic
	maker.Value2 = eta
	maker.Send(key, event)
}
