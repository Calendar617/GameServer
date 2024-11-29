package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goserver/protocol/pbc"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

type Client struct {
	NetId           uint64
	LogicId         uint64
	Name            string
	Conn            net.Conn
	LogicProcessMap map[uint32]func(uint32, []byte)
	Position        [3]float32
}

func NewClient(id uint64, name string, conn net.Conn) *Client {
	client := Client{NetId: id, Name: name, Conn: conn}
	client.Init()
	//wg.Add(1)
	go client.HandleConn()
	return &client
}

func (c *Client) Init() {
	c.LogicProcessMap = make(map[uint32]func(uint32, []byte))
	c.LogicProcessMap[uint32(pbc.CMDDef_world_cs_enter_world)] = c.OnEnterWorld
	c.LogicProcessMap[uint32(pbc.CMDDef_common_udp_heartbeat)] = c.OnHeartBeat
	c.LogicProcessMap[uint32(pbc.CMDDef_world_cs_forward_data)] = c.OnPlayerMove
}

func (c *Client) HandleConn() {
	buffer := make([]byte, 2048)
	for {
		n, err := c.Conn.Read(buffer)
		if err != nil {
			//db 保存属性
			dbp.UpdateRole(c.LogicId, c.Position)
			fmt.Println("Client disconnect", c.Conn.RemoteAddr())
			removeClient(c.Conn)
			//wg.Done()
			return
		}

		msg := &pbc.CsMessageWrapper{}
		temp := buffer[4:n] //unity 客户端把msglen写在包头前4个字节，需要过滤掉
		err1 := proto.Unmarshal(temp, msg)
		if err1 != nil {
			fmt.Println("proto.Unmarshal CsMessageWrapper error =", err1)
		}

		processFun := c.LogicProcessMap[msg.Cmd]
		if processFun != nil {
			processFun(msg.SeqId, msg.Data)
		} else {
			fmt.Println("processFun is null")
		}
	}
}

func (c *Client) send2Client(msgid uint32, seqId uint32, result int32, data []byte) {
	msg := &pbc.ScMessageWrapper{
		Cmd:    msgid,
		SeqId:  seqId,
		Result: result,
		Data:   data,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	//unity 客户端需要把msglen写在包头前4个字节
	msglen := int32(len(data))
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, msglen)
	newbuff := append(bytesBuffer.Bytes()[:], data[:]...)
	c.Conn.Write(newbuff)
}

func (c *Client) OnEnterWorld(seqId uint32, data []byte) {
	msgContent := &pbc.CsEnterWorld{}
	err := proto.Unmarshal(data, msgContent)
	if err != nil {
		fmt.Println("proto.Unmarshal CsEnterWorld error = ", err)
	}
	c.LogicId = msgContent.Uid
	result, Position := dbp.QueryRole(msgContent.Uid)
	if !result {
		pos := [3]float32{0, 0, 0}
		dbp.CreateRole(msgContent.Uid, c.Name, pos)
		fmt.Printf("userid = %d, createrole success:\n", c.LogicId)
	} else {
		fmt.Printf("userid = %d, querydata success:\n", c.LogicId)
		c.Position = Position
	}

	str := msgContent.String()
	fmt.Println(c.Conn.RemoteAddr().String(), "receive data string:\n", str)

	otherplayerinfo := GetOtherPlayerInfo(c.NetId)
	fmt.Println("otherplayerinfo len = ", len(otherplayerinfo))

	msg2client := &pbc.ScEnterWorld{
		Uid:              c.LogicId,
		PosX:             c.Position[0],
		PosY:             c.Position[1],
		PosZ:             c.Position[2],
		OnlinePlayerList: otherplayerinfo,
	}
	data, err = proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal ScEnterWorld error = ", err)
	}
	c.send2Client(uint32(pbc.CMDDef_world_sc_enter_world), seqId, 0, data)

	posInfo := [3]float32{msg2client.PosX, msg2client.PosY, msg2client.PosZ}
	BroadcastPlayerEnterWorldInfo(uint64(c.NetId), posInfo)
}

func (c *Client) OnHeartBeat(seqId uint32, data []byte) {
	msgContent := &pbc.CsReqUdpHeartbeat{}
	err := proto.Unmarshal(data, msgContent)
	if err != nil {
		fmt.Println("proto.Unmarshal CsReqUdpHeartbeat error = ", err)
	}

	//fmt.Printf("Userid = %d HeartBeat\n", c.LogicId)

	msg2client := &pbc.CsReqUdpHeartbeat{}
	data, err = proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal CsReqUdpHeartbeat error = ", err)
	}
	c.send2Client(uint32(pbc.CMDDef_common_udp_heartbeat), seqId, 0, data)
}

func (c *Client) OnPlayerMove(seqId uint32, data []byte) {
	msg := &pbc.CsForwardData{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("proto.Unmarshal CsReqUdpHeartbeat error = ", err)
	}

	fmt.Printf("Userid = %d Move x = %f y = %f z = %f \n", c.LogicId, msg.PosX, msg.PosY, msg.PosZ)

	msg2client := &pbc.CsReqUdpHeartbeat{}
	data, err = proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal CsReqUdpHeartbeat error = ", err)
	}
	c.send2Client(uint32(pbc.CMDDef_world_sc_forward_data), seqId, 0, data)

	posInfo := [3]float32{msg.PosX, msg.PosY, msg.PosZ}
	c.Position[0] = msg.PosX
	c.Position[1] = msg.PosY
	c.Position[2] = msg.PosZ
	BroadcastPlayerMoveInfo(uint64(c.NetId), posInfo)
}

func (c *Client) OtherPlayerMove(id uint64, pos [3]float32) {
	playerinfo := &pbc.PlayerInfo{
		Uid:  id,
		PosX: pos[0],
		PosY: pos[1],
		PosZ: pos[2],
	}
	msg2client := &pbc.PushForwardData{
		ForwardPlayer: playerinfo,
	}
	data, err := proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal CsReqUdpHeartbeat error = ", err)
	}
	//转发给其他玩家，需要增加过滤
	c.send2Client(uint32(pbc.CMDDef_world_push_forward_data), 0, 0, data)
}

func (c *Client) OtherEnterWorld(id uint64, pos [3]float32) {
	playerinfo := &pbc.PlayerInfo{
		Uid:  id,
		PosX: pos[0],
		PosY: pos[1],
		PosZ: pos[2],
	}
	msg2client := &pbc.ScOtherEnterWorld{
		OtherPlayer: playerinfo,
	}
	data, err := proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal ScOtherEnterWorld error = ", err)
	}
	//转发给其他玩家，需要增加过滤
	c.send2Client(uint32(pbc.CMDDef_world_sc_other_enter_world), 0, 0, data)
}

func (c *Client) OtherLeaveWorld(id uint64) {
	msg2client := &pbc.PushDisconnected{
		Uid: id,
	}
	data, err := proto.Marshal(msg2client)
	if err != nil {
		fmt.Println("proto.Marshal PushDisconnected error = ", err)
	}
	//转发给其他玩家，需要增加过滤
	c.send2Client(uint32(pbc.CMDDef_world_push_disconnected), 0, 0, data)
}

func (c *Client) showInfo() {
	fmt.Printf(" client id = %d  name is %s\n", c.LogicId, c.Name)
}
