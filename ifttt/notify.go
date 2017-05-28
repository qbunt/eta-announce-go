package ifttt

import (
	"github.com/jamesmillerio/go-ifttt-maker"
	"os"
)

func Notify(eta string) {
	maker := new(GoIFTTTMaker.MakerChannel)
	key, event := os.Getenv("MAKER_KEY"), os.Getenv("MAKER_EVENT")

	maker.Value1 = eta
	maker.Send(key, event)
}
