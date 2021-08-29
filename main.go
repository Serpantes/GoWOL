package main

import (
	"flag"

	"github.com/linde12/gowol"
)

func main() {
	var Mac = flag.String("mac", "B4:2E:99:1B:5C:57", "Set mac to wake")
	flag.Parse()
	println("Waking " + *Mac)
	if packet, err := gowol.NewMagicPacket(*Mac); err == nil { //B4:2E:99:1B:5C:57
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
	} else {
		println("Error in gowol lib while waking " + *Mac)
	}
}
