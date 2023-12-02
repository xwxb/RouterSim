package consts

type ArpRequestPacket struct {
}

type ArpResponsePacket struct {
	MACAddress
}

type ArpTable map[IPAddress]MACAddress
