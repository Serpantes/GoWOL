package main

import "github.com/linde12/gowol"

func main() {
	if packet, err := gowol.NewMagicPacket("03:AA:FF:67:64:05"); err == nil {
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
	}
}
