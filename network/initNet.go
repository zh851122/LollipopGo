package impl

// LollipopGo 支持的网络类型
const (
	WebSocket = "websocket"
	RPC = "rpc"
	TCP = "tcp"
	UDP = "udp"
	KCP = "kcp"
	NCN = "ncn"
)

func InitNet( netty string ) interface{} {
	switch netty {
	case WebSocket:
		return IMsg
	case RPC:
	case KCP:
	case TCP:
	case UDP:
	case NCN:
	default:
	}
	return nil
}