package d2prox

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Packet []byte

func (p Packet) GsMsgID() byte    { return p[0] }
func (p Packet) RealmMsgID() byte { return p[2] }
func (p Packet) BnetMsgID() byte  { return p[1] }

//
// SID_LOGONREALMEX
//

type LogonRealmExPacket Packet

func (p LogonRealmExPacket) RealmIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:6112", p[20], p[21], p[22], p[23])
}

func (p LogonRealmExPacket) Token() string {
	token := make([]byte, 64)
	copy(token[0:16], p[4:20])
	copy(token[16:64], p[28:76])
	return hex.EncodeToString(token)
}

//
// MCP_STARTUP
//

type McpStartupPacket Packet

func (p McpStartupPacket) Token() string {
	return hex.EncodeToString(p[3:67])
}

//
// MCP_JOINGAME
//

type McpJoinGamePacket Packet

func (p McpJoinGamePacket) Hash() []byte {
	return p[13:17]
}

func (p McpJoinGamePacket) Token() []byte {
	return p[5:7]
}

func (p McpJoinGamePacket) GameIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:4000", p[9], p[10], p[11], p[12])
}

func (p McpJoinGamePacket) Status() int {
	return int(binary.LittleEndian.Uint32(p[17:21]))
}

//
// D2GS_GAMELOGON
// https://redux.bnetdocs.org/?op=packet&pid=131
//

type GsGameLogonPacket Packet

// Token returns game hash + game token as a hexadecimal string
func (p GsGameLogonPacket) Token() string {
	return hex.EncodeToString(p[1:7])
}
