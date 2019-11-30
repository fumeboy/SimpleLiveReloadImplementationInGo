package SimpleLiveReloadImplementationInGo

import (
	"github.com/gorilla/websocket"
	"gopkg.in/fsnotify.v1"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// server_ contains a single lrserver instance's data
type server_ struct {
	*http.Server
	*log.Logger
	*ConnSet
}

// New creates a new server instance
func server(port string, dstPath string) (server *server_) {
	router := http.NewServeMux()
	logPrefix := "[lr] "
	// Create server
	server = &server_{}
	server.Server = &http.Server{
		Addr:port,
		Handler:router,
		ErrorLog:log.New(os.Stderr, logPrefix, 0),
	}
	server.Logger = log.New(os.Stderr, logPrefix, 0)
	server.ConnSet =&ConnSet{
		connections: make(map[*websocket.Conn]struct{}),
		m:           sync.Mutex{},
		reloadChan:  make(chan *reload_),
	}
	server.handleLiveReloadJS(router)
	server.handleLiveReloadRequest(router)
	server.handleStaticSiteRequest(router, dstPath)
	return server
}

func (s *server_) run() {
	if l, err := net.Listen("tcp", s.Addr); err != nil {
		panic("")
	} else {
		s.Println("listening on " + s.Addr)
		if s.Serve(l) != nil {
			panic("")
		}
	}
}

func (s *server_) watch(srcPath string) {
	go s.reload()
	debounceLock := false
	if watcher, err := fsnotify.NewWatcher(); err != nil {
		log.Fatalln(err)
	}else{
		defer watcher.Close()
		// Watch dir
		if err := watcher.Add(srcPath); err != nil {
			log.Fatalln(err)
		}
		for {
			select {
			case event := <-watcher.Events:
				if !debounceLock{
					s.reloadChan <- reload(event.Name)
					debounceLock = true
					go func() {
						time.Sleep(time.Millisecond*500)
						debounceLock = false
					}()
				}
			case err := <-watcher.Errors:
				log.Println(err)
			}
		}
	}
}
