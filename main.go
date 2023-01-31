package main

import (
	"github.com/Serpantes/GoWOL/config"
	"github.com/Serpantes/GoWOL/server"
)
func main() {

	config.InitConfig()
	config.InitLogger()

	server.InitServer()


	// var Mac = flag.String("mac", "B4:2E:99:1B:5C:57", "Set mac to wake")
	// flag.Parse()
	// println("Waking " + *Mac)
	// if packet, err := wol.NewMagicPacket(*Mac); err == nil {
	// 	packet.Send("255.255.255.255") // send to broadcast
	// 	//	packet.SendPort("255.255.255.255", "7") // specify receiving port
	// } else {
	// 	fmt.Println(err.Error())
	// }
}
