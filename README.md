# tracer

> A simple imprementation of <b>traceroute</b> in golang.

The orginal traceroute works by continously incrementing the TTL variable in the header of a UDP packet. 
- When TTL = 1, the packet will be dropped after it reaches the first hop, if it was set to 2, the packet will be dropped after it reaches the second hop.
- When a server/router drops a packet, it returns a ICMP Time Exceeded message back. Parsing this message allows us to retrieve info on the particular hop. 
- Once the destination is reached (last hop) we are returned a ICMP Destination Unreachable message.


If a router along the way decides to not send the ICMP time exceeded (i.e. TTL reached en-route) or destination 
unreachable message (i.e. UDP-packet reached final host but port closed, proper behaviour though), you will get a timeout at 
that point in the tracer.

This simple implementation of traceroute uses UDP packets,that is sending UDP-packets with a low TTL, starting from 1, 
and increasing by 1 per step. If a packet dies at a router, i.e. TTL becomes 0, that router should, according to RFC 792 
and some others, send an ICMP "time exceeded" message, ergo saying that we could not deliver the package within the timeframe, 
but at least we tell you that your package died. However some sites BLOCK UDP-based  packets therefore tracer will run until it 
times out.

There are two other methods for doing a traceroute, such as the following. But in short we can also send ICMP Echo packets 
or TCP SYN packets. To summarise, there are three methods all based on an ever increasing TTL to map the "hosts" along the route:

    - UDP to random port (usually 33434 + 100) at host with low TTL
        - In my experience the default for all command line tools, such as traceroute and tracert
    - ICMP Echo to host with low TTL
        - I've encountered this in a couple of graphical tools, also as an option for most command line tools.
    - TCP SYN, often to port 80, that way the traffic is "kinda" masked as http-traffic is passes by many routers which normally drop ICMP Echoes and UDP-packets to weird ports.
       -  Neat trick and "new" method, although unorthodox, for finding a route to a host. Unorthodox in that you are in a way missusing Internet-standards. Exists as an option for most command line tools.

The router may pass on normal traffic, thus allowing your TCP-based http request to complete, but it may silently drop UDP to weird ports, half open TCP to weird ports or ICMP pings with low TTL, leaving your local traceroute process waiting and then timing out on that stop.
share improve this answer

If a server between you and cnn.com decides not to send a response for your traceroute, it can truncate the remaining responses.
share improve this answer 
