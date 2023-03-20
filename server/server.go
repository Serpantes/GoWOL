package server

import (
	"net/http"
	"strings"

	"github.com/Serpantes/GoWOL/config"
	"github.com/Serpantes/GoWOL/wol"

	log "github.com/sirupsen/logrus"
)

func InitServer() {

	http.HandleFunc("/wake", func(w http.ResponseWriter, r *http.Request) {
		wakeHandler(w, r)
	})

	log.Info("Starting web server on " + config.Config.ServerHost)

	err := http.ListenAndServe(config.Config.ServerHost, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func wakeHandler(w http.ResponseWriter, r *http.Request){
	log.Debug("Wake up request")
	var target = strings.TrimPrefix(r.URL.String(), "/wake?")
	var mac = parseTarget(target)
	if mac == "" {
		log.Warn("Unknown target: ", target)
		w.Write([]byte("Unknown target"))
		return
	}
	log.Info("Waking up ", mac, " (",target,")")

	if packet, err := wol.NewMagicPacket(mac); err == nil {
		packet.Send("192.168.100.255") // send to broadcast
		//	packet.SendPort("255.255.255.255", "7") // specify receiving port
	} else {
		log.Error(err.Error())
		return
	}

	log.Info("Done")

	w.Write([]byte(target + " is awake"))
}

func parseTarget(target string) string {
	if strings.ContainsAny(target, ":") {
		return target
	} else {
		switch strings.ToLower(target) {
		case "serpantes":
			return "22:20:5c:06:13:f9"
		case "takero":
			return "B4:2E:99:1B:5C:57"
		}
	}
	return ""
}