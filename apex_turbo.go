package ApexTurbo

import "sync"

type (
	App struct {
		// Route, Reference Router Type
		Router *Router

		// Configuration Data, Reference Config Type
		Config *Config

		// ConnectionQueue
		ConnectionQueue map[string]*Connection

		// ConnectionQueue Lock
		ConnectionQueueLock sync.RWMutex
	}
)

func Run() {

}
