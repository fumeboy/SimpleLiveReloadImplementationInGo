package SimpleLiveReloadImplementationInGo

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func (s *server_) handleLiveReloadJS(r *http.ServeMux){
	// Handle JS
	r.HandleFunc("/livereload.js", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/javascript")
		if _, err := rw.Write([]byte(fmt.Sprintf(script, s.Addr))); err != nil {
			s.ErrorLog.Println(err)
		}
	})
}

func (s *server_) handleLiveReloadRequest(r *http.ServeMux){
	// Handle reload requests
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	r.HandleFunc("/livereload", func(rw http.ResponseWriter, req *http.Request) {
		if conn, err := upgrader.Upgrade(rw, req, nil); err != nil {
			s.ErrorLog.Println(err)
			return
		}else{
			s.add(conn)
		}
	})
}

func (s *server_) handleStaticSiteRequest(r *http.ServeMux, dstPath string){
	// return .html
	r.Handle("/", http.FileServer(http.Dir(dstPath)))
}
