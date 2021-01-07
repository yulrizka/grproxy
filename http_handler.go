package grproxy

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type HTTPHandler struct {
	upgrader websocket.Upgrader
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}
func (h *HTTPHandler) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	ctx := r.Context()

	go h.writer(ctx, ws)
	h.reader(ctx, ws)
}

func (h *HTTPHandler) reader(ctx context.Context, ws *websocket.Conn) {
	defer ws.Close()
	const pongWait = 60 * time.Second

	ws.SetReadLimit(512)
	_ = ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		_ = ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (h *HTTPHandler) writer(ctx context.Context, ws *websocket.Conn) {
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()

	const writeWait = 10 * time.Second
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-tick.C:
			_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
			err := ws.WriteJSON(map[string]string{
				"msg": t.String(),
			})
			if err != nil {
				log.Printf("write json failed: %v", err)
				return
			}
		}
	}
}
