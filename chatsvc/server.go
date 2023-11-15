package chatsvc

import (
	"fmt"
	"sync"
)

type ChatServer struct {
}

type msgItem struct {
	Name   string
	Msg    string
	UserID int
}

type MsgHandler struct {
	qu []msgItem
	mu sync.Mutex
}

var msgHandleObj = MsgHandler{}

func (c ChatServer) Chat(server ChatService_ChatServer) error {
	//TODO implement me
	panic("implement me")
}

func receive(cs ChatService_ChatServer, userID int) {
	for {
		msg, err := cs.Recv()
		if err != nil {
			fmt.Println(err)
		} else {
			msgHandleObj.mu.Lock()

			msgHandleObj.qu = append(msgHandleObj.qu, msgItem{
				Name:   msg.Name,
				Msg:    msg.Msg,
				UserID: userID,
			})

			msgHandleObj.mu.Unlock()
		}
	}
}

func send(cs ChatService_ChatServer, userID int) {
	for {
		for {
			msgHandleObj.mu.Lock()

			if len(msgHandleObj.qu) == 0 {
				msgHandleObj.mu.Unlock()
				break
			}

			userName := msgHandleObj.qu[0].Name
			msg := msgHandleObj.qu[0].Msg
			senderUserID := msgHandleObj.qu[0].UserID

			msgHandleObj.mu.Unlock()

			if senderUserID != userID {
				err := cs.SendMsg(&FromServer{Name: userName, Msg: msg})
				if err != nil {
					fmt.Println(err)
				}

				msgHandleObj.mu.Lock()
				if len(msgHandleObj.qu) > 1 {
					msgHandleObj.qu = msgHandleObj.qu[1:] // delete the message at index 0 after sending to receiver
				}
			}
		}
	}
}
