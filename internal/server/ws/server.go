package ws

import (
	"context"
	"excel2config/internal/dao"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/gorilla/websocket"
	"net/http"
)

type Server struct {
	*http.Server
	d dao.Dao
}

var (
	mgr         *ClientMgr
	uidAutoIncr int

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func New(dao dao.Dao) *Server {
	wss := &Server{
		Server: &http.Server{
			Addr:    ":8001",
			Handler: nil,
		},
		d: dao,
	}
	wss.Handler = wss.defaultHandler()
	go func() {
		log.Info("websocket server listen at: 8001")
		if err := wss.Server.ListenAndServe(); err != nil {
			panic("websocket server start failed")
		}
	}()
	mgr = new(ClientMgr)
	mgr.Clients = make(map[int]*Client)
	return wss
}

func (wss *Server) defaultHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		gridKey := values.Get("g")
		if gridKey == "" {
			log.Errorw(context.TODO(), "ws conn without gridKey")
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Errorw(context.TODO(), "ws conn upgrade error: ", err)
			return
		}
		uidAutoIncr += 1
		client := NewClient(conn, gridKey)
		mgr.AddClient(uidAutoIncr, client)
		svr := newService(client, wss.d)
		go svr.readAndServe()
		go client.waitAndWrite()
	})
	return mux
}
