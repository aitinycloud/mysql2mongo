package handle

import (
	"time"
)

func SyncHandle(sqlinfo SqlInfoT) {

}

func GetInfoHandle(msg chan interface{}) {
	lastSavedTime := time.Now()
	for v := range msg {
		if val, ok := v.(SqlInfoT); ok {
			//wait...
			now := time.Now()
			if now.Sub(lastSavedTime) < 20*time.Millisecond {
				//time.Sleep(1 * time.Millisecond)
				lastSavedTime = now
			}
			SyncHandle(val)
		}
	}
}

func SyncInit() {
	clientch := MSGPB.Sub(TOPIC)
	go GetInfoHandle(clientch)
}
