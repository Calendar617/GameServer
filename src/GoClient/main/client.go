package main

import (
	"bufio"
	"fmt"
	"goclient/protocol/pbc"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	Id              uint64
	Name            string
	Conn            net.Conn
	SeqId           uint32
	LogicProcessMap map[uint32]func([]byte)
	LoginStatus     bool
	Position        [3]float32
}

func CreateClient(id uint64, name string, conn net.Conn) *Client {
	client := &Client{Id: id, Name: name, Conn: conn}
	client.Init()
	client.EnterWorld(id)
	go client.HandleConn()
	go client.Update()
	go client.Move()
	return client
}

func (c *Client) Init() {
	c.LoginStatus = false
	c.LogicProcessMap = make(map[uint32]func([]byte))
	c.LogicProcessMap[uint32(pbc.CMDDef_world_sc_enter_world)] = c.OnEnterWorld
	c.LogicProcessMap[uint32(pbc.CMDDef_common_udp_heartbeat)] = c.OnHeartBeat
	c.LogicProcessMap[uint32(pbc.CMDDef_world_sc_forward_data)] = c.OnPlayerMove
	c.LogicProcessMap[uint32(pbc.CMDDef_world_sc_other_enter_world)] = c.OnOtherPlayerEnterWorld
	c.LogicProcessMap[uint32(pbc.CMDDef_world_push_forward_data)] = c.OnOtherPlayerMove
	c.LogicProcessMap[uint32(pbc.CMDDef_world_push_disconnected)] = c.OnOtherPlayerLeaveWorld
}

func (c *Client) HandleConn() {
	buffer := make([]byte, 2048)
	for {
		n, err := c.Conn.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnect", c.Conn.RemoteAddr())
			return
		}

		msg := &pbc.ScMessageWrapper{}
		temp := buffer[0:n]
		err1 := proto.Unmarshal(temp, msg)
		if err1 != nil {
			fmt.Println("proto.Unmarshal CsMessageWrapper error =", err1)
		}

		processFun := c.LogicProcessMap[msg.Cmd]
		if processFun != nil {
			processFun(msg.Data)
		} else {
			fmt.Println("processFun is null cmdId = ", msg.Cmd)
		}
	}
}

func (c *Client) CloseConn() {
	c.Conn.Close()
}

func (c *Client) Update() {
	timer := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-timer.C:
			if c.LoginStatus {
				c.HeartBeat()
			}
		}
	}
}

func (c *Client) Move() {
	timer := time.NewTicker(2000 * time.Millisecond)
	for {
		select {
		case <-timer.C:
			if c.LoginStatus {
				var deltaX = rand.Float32()
				var deltaY = rand.Float32()
				var deltaZ = rand.Float32()
				c.Position[0] += deltaX
				c.Position[1] += deltaY
				c.Position[2] += deltaZ

				c.PlayerMove(c.Position[0], c.Position[1], c.Position[2])
			}
		}
	}
}

func (c *Client) send2Server(msgid uint32, data []byte) {
	msg := &pbc.CsMessageWrapper{
		Cmd:   msgid,
		SeqId: c.SeqId,
		Data:  data,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	c.SeqId++
	c.Conn.Write(data)
}

func (c *Client) EnterWorld(id uint64) {
	msg := &pbc.CsEnterWorld{
		Uid: id,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("proto.Marshal EnterWorld error", err)
	}
	c.send2Server(uint32(pbc.CMDDef_world_cs_enter_world), data)
}

func (c *Client) OnEnterWorld(data []byte) {
	msg := &pbc.ScEnterWorld{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal OnEnterWorld error = ", err)
	}
	c.Position[0] = msg.PosX
	c.Position[1] = msg.PosY
	c.Position[2] = msg.PosZ
	infolist := msg.GetOnlinePlayerList()
	infolistlen := len(infolist)
	c.LoginStatus = true
	fmt.Printf("UserID = %d, Enter World Success Pos = %f other player num = %d\n", msg.Uid, c.Position, infolistlen)
}

func (c *Client) HeartBeat() {
	msg := &pbc.CsReqUdpHeartbeat{
		Seqid: c.SeqId,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("proto.Marshal HeartBeat error", err)
	}
	c.send2Server(uint32(pbc.CMDDef_common_udp_heartbeat), data)
}

func (c *Client) OnHeartBeat(data []byte) {
	msg := &pbc.CsReqUdpHeartbeat{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal OnHeartBeat error = ", err)
	}
}

func (c *Client) PlayerMove(x float32, y float32, z float32) {
	msg := &pbc.CsForwardData{
		PosX: x,
		PosY: y,
		PosZ: z,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("proto.Marshal CsForwardData error", err)
	}
	c.send2Server(uint32(pbc.CMDDef_world_cs_forward_data), data)
}

func (c *Client) OnPlayerMove(data []byte) {
	msg := &pbc.ScForwardData{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal OnPlayerMove error = ", err)
	}
}

func (c *Client) OnOtherPlayerMove(data []byte) {
	msg := &pbc.PushForwardData{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal OnPlayerMove error = ", err)
	}

	//fmt.Printf("Userid = %d Move x = %f y = %f z = %f \n", msg.ForwardPlayer.GetUid(), msg.ForwardPlayer.PosX, msg.ForwardPlayer.PosY, msg.ForwardPlayer.PosZ)
}

func (c *Client) OnOtherPlayerEnterWorld(data []byte) {
	msg := &pbc.ScOtherEnterWorld{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Marshal OnOtherPlayerEnterWorld error", err)
	}

	fmt.Printf("Userid = %d Enter World\n", msg.OtherPlayer.Uid)
}

func (c *Client) OnOtherPlayerLeaveWorld(data []byte) {
	msg := &pbc.PushDisconnected{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal OnOtherPlayerLeaveWorld error = ", err)
	}
	fmt.Printf("Userid = %d Leave World\n", msg.Uid)
}

func (c *Client) Chat() {
	fmt.Println("Client Can Chat")
	var iscmd bool
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Client scan message error", err)
			break
		}

		newinput := strings.TrimSpace(string(input))
		fmt.Println("Client scan message is", newinput)
		msgArray := strings.Split(newinput, " ")
		if len(msgArray) <= 0 {
			continue
		}
		fmt.Println("Client scan message is", msgArray[0])
		iscmd = isCommand(msgArray[0])
		if !iscmd {
			if c.Conn != nil {
				var content string
				for _, val := range msgArray {
					content += " "
					content += val
				}

				c.Conn.Write([]byte(content))
				fmt.Println("Client send message is", content)
			}
		} else {
			Command(msgArray)
		}
	}
}
