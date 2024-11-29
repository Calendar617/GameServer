package main

import (
	"fmt"
	"goserver/protocol/pbc"
	"net"
	"strconv"
)

var clientMap map[uint64]*Client = make(map[uint64]*Client)

var netConnMap map[net.Conn]uint64 = make(map[net.Conn]uint64)

var clientNetId = uint64(1)

func addClient(conn net.Conn) {
	clientName := "player" + strconv.Itoa(int(clientNetId))
	clientMap[clientNetId] = NewClient(clientNetId, clientName, conn)
	netConnMap[conn] = clientNetId
	clientNetId++
	fmt.Printf("Client connect Server ip = %s, clientid = %d clientName = %s\n", conn.RemoteAddr().String(), clientNetId, clientName)
}

func removeClient(conn net.Conn) {
	id, exists1 := netConnMap[conn]
	if exists1 {
		_, exists2 := clientMap[id]
		if exists2 {
			BroadcastPlayerLeaveWorld(id)
			delete(clientMap, id)
			fmt.Printf("delete %d from Client map\n", id)
		}
	}
}

func showAllClient() {
	fmt.Println("client map length is ", len(clientMap))
	for _, val := range clientMap {
		c := val
		c.showInfo()
	}
}

func BroadcastPlayerMoveInfo(id uint64, posInfo [3]float32) {
	//fmt.Println("broadcast player move info")
	for _, val := range clientMap {
		c := val
		if c.NetId != id {
			c.OtherPlayerMove(id, posInfo)
		}
	}
}

func BroadcastPlayerEnterWorldInfo(id uint64, posInfo [3]float32) {
	fmt.Println("broadcast player enterworld info")
	for _, val := range clientMap {
		c := val
		if c.NetId != id {
			c.OtherEnterWorld(id, posInfo)
		}
	}
}

func BroadcastPlayerLeaveWorld(id uint64) {
	fmt.Println("broadcast player leave world")
	for _, val := range clientMap {
		c := val
		if c.NetId != id {
			c.OtherLeaveWorld(id)
		}
	}
}

// 有没有内存泄漏，这个要弄懂才行
func GetOtherPlayerInfo(id uint64) []*pbc.PlayerInfo {
	var infolist []*pbc.PlayerInfo
	for _, val := range clientMap {
		c := val
		if c.NetId != id {
			info := pbc.PlayerInfo{
				Uid:  c.LogicId,
				PosX: c.Position[0],
				PosY: c.Position[1],
				PosZ: c.Position[2],
			}
			infolist = append(infolist, &info)
		}
	}

	return infolist
}
