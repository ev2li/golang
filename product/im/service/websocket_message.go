package service

import (
	"im/define"
	"im/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常:" + err.Error(),
		})
		return
	}
	defer conn.Close()
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	wc[uc.Identity] = conn
	for {
		ms := new(define.MessageStruct)
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("Read Error:%v\n", err)
			return
		}
		for _, cc := range wc {
			err = cc.WriteMessage(websocket.TextMessage, []byte(ms.Message))
			log.Printf("Write Message Error:%v\n", err)
			return
		}
	}

}
