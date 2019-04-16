package d2prox

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
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

// RealmIP returns the realm server ip
func (p LogonRealmExPacket) RealmIP() net.IP {
	return net.IP(p[20:24])
}

// RealmPort returns the realm port number
func (p LogonRealmExPacket) RealmPort() int {
	return int(binary.BigEndian.Uint16(p[24:26]))
}

// RealmTarget returns a connection string combining the realm ip and the realm port
func (p LogonRealmExPacket) RealmTarget() string {
	return fmt.Sprintf("%s:%d", p.RealmIP(), p.RealmPort())
}

// Token returns all the MCP chunk data required to authenticate with the realm server as a hex string
func (p LogonRealmExPacket) Token() string {
	token := make([]byte, 64)
	copy(token[0:16], p[4:20])
	copy(token[16:64], p[28:76])
	return hex.EncodeToString(token)
}

// SetRealmPort modifies the realm port
func (p LogonRealmExPacket) SetRealmPort(port int) {
	binary.BigEndian.PutUint16(p[24:26], uint16(port))
}

// SetRealmIP modifies the realm ip
func (p LogonRealmExPacket) SetRealmIP(ip net.IP) {
	copy(p[20:24], ip)
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

// McpJoinedGamePacket wraps MCP_JOINGAME (S->C)
// https://redux.bnetdocs.org/?op=packet&pid=107
type McpJoinedGamePacket Packet

// Hash returns the game hash as a byte array
func (p McpJoinedGamePacket) Hash() []byte {
	return p[13:17]
}

// Token returns the game token as a byte array
func (p McpJoinedGamePacket) Token() []byte {
	return p[5:7]
}

// GameIP returns the game server ip & port as a string
func (p McpJoinedGamePacket) GameIP() net.IP {
	return net.IP(p[9:13])
}

// SetGameIP modifies the game server ip
func (p McpJoinedGamePacket) SetGameIP(ip net.IP) {
	copy(p[9:13], ip)
}

// Status returns the join game status code
func (p McpJoinedGamePacket) Status() int {
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

// GsChatMessagePacket represents a chat message sent from the client
type GsChatMessagePacket Packet

// Message returns the chat message as a string
func (p GsChatMessagePacket) Message() string {
	pb := PacketBuffer(p)
	return pb.NullString(3)
}

// Target returns the message target as a string
func (p GsChatMessagePacket) Target() string {
	pb := PacketBuffer(p)
	start := pb.IndexOf(0x00, 3) + 1
	return pb.NullString(start)
}
