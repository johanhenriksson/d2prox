package d2prox

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Packet is a wrapper around a byte array with some utility features
type Packet []byte

// GsMsgID attempts to extract a game server message id from the packet. May panic
func (p Packet) GsMsgID() byte {
	return p[0]
}

// RealmMsgID attempts to extract a realm server message id from the packet. May panic
func (p Packet) RealmMsgID() byte {
	return p[2]
}

// BnetMsgID attempts to extract a battle.net message id from the packet. May panic
func (p Packet) BnetMsgID() byte {
	return p[1]
}

//
// battle.net packets
//

// LogonRealmExPacket wraps SID_LOGONREALMEX (S->C)
// https://redux.bnetdocs.org/?op=packet&pid=237
type LogonRealmExPacket Packet

// RealmIP returns the realm server ip and port as a string
func (p LogonRealmExPacket) RealmIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:6112", p[20], p[21], p[22], p[23])
}

// Token returns all the MCP chunk data required to authenticate with the realm server as a hex string
func (p LogonRealmExPacket) Token() string {
	token := make([]byte, 64)
	copy(token[0:16], p[4:20])
	copy(token[16:64], p[28:76])
	return hex.EncodeToString(token)
}

//
// realm server packets
//

// McpStartupPacket wraps MCP_STARTUP (C->S)
// https://redux.bnetdocs.org/?op=packet&pid=320
type McpStartupPacket Packet

// Token returns all the MCP chunk data required to authenticate with the realm server as a hex string
func (p McpStartupPacket) Token() string {
	return hex.EncodeToString(p[3:67])
}

// McpJoinGamePacket wraps MCP_JOINGAME (S->C)
// https://redux.bnetdocs.org/?op=packet&pid=107
type McpJoinGamePacket Packet

// Hash returns the game hash as a byte array
func (p McpJoinGamePacket) Hash() []byte {
	return p[13:17]
}

// Token returns the game token as a byte array
func (p McpJoinGamePacket) Token() []byte {
	return p[5:7]
}

// GameIP returns the game server ip & port as a string
func (p McpJoinGamePacket) GameIP() string {
	return fmt.Sprintf("%d.%d.%d.%d:4000", p[9], p[10], p[11], p[12])
}

// Status returns the join game status code
func (p McpJoinGamePacket) Status() int {
	return int(binary.LittleEndian.Uint32(p[17:21]))
}

//
// game server packets
//

// GsGameLogonPacket wraps D2GS_GAMELOGON (S->C)
// https://redux.bnetdocs.org/?op=packet&pid=131
type GsGameLogonPacket Packet

// Token returns game hash + game token as a hexadecimal string
func (p GsGameLogonPacket) Token() string {
	return hex.EncodeToString(p[1:7])
}
