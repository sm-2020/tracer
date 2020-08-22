# tracer

> A simple imprementation of <b>traceroute</b> in golang.

The orginal traceroute works by continously incrementing the TTL variable in the header of a UDP packet. 
- When TTL = 1, the packet will be dropped after it reaches the first hop, if it was set to 2, the packet will be dropped after it reaches the second hop.
- When a server/router drops a packet, it returns a ICMP Time Exceeded message back. Parsing this message allows us to retrieve info on the particular hop. 
- Once the destination is reached (last hop) we are returned a ICMP Destination Unreachable message.
