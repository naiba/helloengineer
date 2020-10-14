package handler

import (
	"encoding/json"

	"github.com/gofiber/websocket/v2"

	"github.com/naiba/helloengineer/internal/model"
	"github.com/naiba/helloengineer/pkg/hub"
	"github.com/naiba/helloengineer/pkg/util"
)

func WS(pubsub *hub.Hub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		c.SetPingHandler(c.PingHandler())
		c.SetPongHandler(c.PongHandler())
		roomID := c.Params("meetingID")
		pubsub.Subscribe <- hub.Subscription{
			Conn:  c,
			Topic: roomID,
		}
		var m model.WsMsg
		var err error
		for {
			err = c.ReadJSON(&m)
			if err != nil {
				if !websocket.IsCloseError(err) && !websocket.IsUnexpectedCloseError(err) {
					util.Infof(0, "websocket error: %+v", err)
				}
				c.Close()
				break
			}
			if err == nil {
				data, err := json.Marshal(m)
				if err != nil {
					util.Errorf(0, "%+v", err)
					continue
				}
				pubsub.Broadcast <- hub.Message{
					Topic: roomID,
					Data:  data,
					From:  c,
				}
			}
		}
	}
}
