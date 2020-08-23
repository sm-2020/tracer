package tracer

import (
    "fmt"
    "net"
    "syscall"
    "time"
)

type Hop struct {
    Status  bool
    Addr  syscall.Sockaddr
    TTL   int
    Num   int
    Duration time.Duration
}

func (h  *Hop) Domain() string {
    ip := h.Addr.(*syscall.SockaddrInet4).Addr
    ipString := fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
    host, err := net.LookupAddr(ipString)
    if err != nil {
        return ipString
    } else {
        return host[0]
    }
}

func (h *Hop) GetIP() string {
    ip := h.Addr.(*syscall.SockaddrInet4).Addr
    ipString := fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])
    return ipString
}
