package SimpleLiveReloadImplementationInGo

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ConnSet struct {
	connections map[*websocket.Conn]struct{}
	m   sync.Mutex
	reloadChan chan *reload_
}

func (cs *ConnSet) add(c *websocket.Conn) {
	cs.m.Lock()
	cs.connections[c] = struct{}{}
	cs.m.Unlock()
}

func (cs *ConnSet) remove(c *websocket.Conn) {
	cs.m.Lock()
	delete(cs.connections, c)
	cs.m.Unlock()
}

type reload_ struct {
	Path    string `json:"path"`
}

func reload(filename string) *reload_ {
	return &reload_{
		Path:    filename,
	}
}

func (s *server_) reload() {
	for {
		msg := <- s.reloadChan
		s.Println("requesting reload: " + msg.Path)
		for conn := range s.connections {
			if err := conn.WriteJSON(msg); err != nil{
				s.remove(conn)
			}
		}
	}
}
