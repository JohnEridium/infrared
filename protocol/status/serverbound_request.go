package status

import (
	"github.com/JohnEridium/infrared/protocol"
)

const ServerBoundRequestPacketID byte = 0x00

type ServerBoundRequest struct{}

func (pk ServerBoundRequest) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ServerBoundRequestPacketID,
	)
}
