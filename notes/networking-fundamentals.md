# Networking fundamentals
---

## Basics

### TCP
TCP is one of the protocols used in the **transport layer** of the **TCP/IP** model.

![TCP/IP Model Image](https://www.computernetworkingnotes.org/images/intro/ccna-study-guide/similarities-and-differences-between-osi-and-tcp-ip-model.png)

The **transport layer** is the **forth layer** in the updated TCP/IP model. It is an **end-to-end layer** used to **deliver messages** to a **host**. It is termed an **end-to-end layer** because it **provides a point-to-point connection** rather than hop-to-hop, between the **source host** and **destination host** to deliver the services reliably. The **unit of data** encapsulation in the **Transport Layer** is a **segment**. 

**TCP (Transmission Control Protocol)** is a **communications standard** that enables application programs and computing devices to **exchange messages over a network**. It is designed to **send packets across the internet** and **ensure** the **successful delivery** of data and messages over networks.

A **segment** of **TCP** has the following structure:

![TCP Segment Structure](https://media.geeksforgeeks.org/wp-content/uploads/TCPSegmentHeader-1.png)

where:
- **Source Port Address:** a **16-bit field** that holds the **port address** of the application that is **<ins>sending</ins> the data segment**;
 
- **Destination Port Address:** a **16-bit field** that holds the **port address** of the application in the host that is **<ins>receiving</ins> the data segment**;
 
- **Sequence Number:** a **32-bit field** that holds the **sequence number**, i.e, the **byte number** of the **first byte** that is **sent** in **that particular segment**. It is **used to reassemble the message** at the **receiving end of the segments** that are received **out of order**; 
 
- **Acknowledgement Number:** a **32-bit field** that holds the **acknowledgement number**, i.e, the **byte number** that the **receiver expects to receive next**. It is an **acknowledgement** for the **previous bytes** being **received successfully**; 
 
- **Header Length (HLEN):** a **4-bit field** that indicates the **length** of the **TCP header** by a **number of 4-byte words in the header**.
  
   For example:
   - if the header is **20 bytes(min length of TCP header)**, then **this field** will hold the value **5 (because 5 x 4 = 20)**;
   - if the header is **60 bytes(max length of TCP header)**, then **this field** will hold the value **15(because 15 x 4 = 60)**. 

    Hence, the **value** of this field is always **between 5 and 15**. 
 
- **Control flags:** are **6 1-bit control bits** that **control**:
  -  **connection establishment**;
  -  **connection termination**;
  -  **connection abortion****, 
  -  **flow control**
  -  **mode of transfer** 
  etc. 
  
    The flags are the following:
    - **URG:** Urgent pointer is valid;
    - **ACK:** Acknowledgement number is valid( used in case of cumulative acknowledgement);
    - **PSH:** Request for push;
    - **RST:** Reset the connection;
    - **SYN:** Synchronize sequence numbers;
    - **FIN**: Terminate the connection;

- **Window size:** a **16-bit field** that holds the value of the **receive window**, i.e, the **number of window size units** that the **sender** of this segment is **currently willing to receive**;
 
- **Checksum:** a **16-bit field** that is used for **error-checking** of the **TCP header**, the **payload** and an **IP pseudo-header**. The **pseudo-header** consists of the **source IP address**, the **destination IP address**, the **protocol number for the TCP protocol (6)** and the **length of the TCP headers** and **payload (in bytes)**;
 
- **Urgent pointer:** a **16-bit field** (valid **only if the <inc>URG control</inc> flag is set**) is used to **point to data** that is **urgently required** (needs to **reach the receiving process at the earliest**). The **value** of this field is **added to the sequence number** to get the **byte number** of the **last urgent byte**;



### Network namespace
**References:** 
- [Scott's Webblog - Introducing Linux Network Namespaces](https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/)
- [ip-netns(8) — Linux manual page](https://man7.org/linux/man-pages/man8/ip-netns.8.html)

A **network namespace** is logically **another copy** of the **network stack**, with its *own routes, firewall rules, and network devices*.

By **default** a **process inherits** its **network namespace** from its **parent**. Initially **all the processes share the same default network namespace** from the init process.

#### Common commands:
- **Creating** a new network namespace:
  `ip netns add <new-namespace-name>`
- **Listing** all network namespace:
  `ip netns list`    
- **List** routing table entries:
  `ip route list`
- **Assigning interaces** to network namespace:
    > **NOTE:** **Virtual Ethernet interfaces (`veth`) always come in pairs** and they are **connected like a tube** — whatever comes in one veth interface will come out the other peer veth interface. As a result, **you can use** veth interfaces **to connect a network namespace to the outside world** via the **“default”** or **“global”** namespace where physical interfaces exist.
    - **Create** `veth` interfaces:
      `ip link add <pair1-name> netns <pair1-namespace> type veth peer <pair2-name> netns <pair2-namespace>`
    - **List** `veth` interfaces: from the **global namespace**:
      `ip link list`
    - **Move** a `veth` to from the **global namespace** to **another namespace**:
      `ip link set <pair2-name> netns <pair2-namespace>`
    - **List** `veth`s from **another namespace**:
      `ip netns exec <namespace-name> ip link list`
    - **Add/Delete** an **IP address** to a `veth`:
      `ip addr { add | del } <ip-address> dev <veth-name>`
      where
      `<ip-address>` $\rightarrow$ `<network-address>.<host-address>/<mask>`
    - **Bring up/down** an interface from **another namespace**:W
      `ip link set dev <veth-name> { up | down }` 

### NICs
A **Network Interface Card (NIC)** is a **hardware component** without which a computer cannot be **connected over a network**. It is a **circuit board** installed in a computer that provides a **dedicated network connection** to the computer. It is also called **Network Interface Controller**, **Network Adapter** or **LAN adapter**.

The NIC **allows computers** to **communicate** over a computer network, either by **using cables or wirelessly**. The NIC is **both** a **Physical Layer** and **Data Link Layer** device, as it **provides physical access** to a **networking medium** and, for **IEEE 802** and **similar networks**, provides a **low-level addressing system** through the use of **MAC addresses** that are **uniquely assigned** to network interfaces.

### IP Tables

**Reference:** [The Beginner’s Guide to iptables, the Linux Firewall](https://www.howtogeek.com/177621/the-beginners-guide-to-iptables-the-linux-firewall/)

**iptables** is a **user-space utility program** that **allows a system administrator** to **configure the IP packet filter rules** of the **Linux kernel firewall**, implemented as different Netfilter modules. The **filters** are **organized** in **different tables**, which contain **chains of rules** for how to **treat network traffic packets**. Different **kernel modules and programs** are currently used for **different protocols**; *iptables applies to IPv4, ip6tables to IPv6, arptables to ARP, and ebtables to Ethernet frames*;

#### Chains
**iptables** uses **three different chains**: **input**, **forward**, and **output**:

- **Input:** used to **control** the behavior for **incoming connections**. For example, if a *user attempts to SSH into your PC/server*, iptables will attempt to *match the IP address and port* to a *rule* in the *input chain*;

- **Forward:** used for **incoming connections** that **aren’t** actually being **delivered locally**. Think of a **router** – data is always being sent to it but rarely actually destined for the router itself; **the data is just forwarded to its target**. Unless you’re doing some kind of *routing, NATing, or something else* on your system that requires forwarding, you won’t even use this chain;

- **Output:** used for **outgoing connections**. For example, if you *try to ping howtogeek.com*, iptables will *check its output chain* to see what the *rules* are regarding ping and howtogeek.com before making a decision to allow or deny the connection attempt.

Check iptables chains:
`iptables -L -v`

#### How to accept/drop iptables chains
`iptables --policy { INPUT | FORWARD | OUTPUT } { ACCEPT | DROP | REJECT }`

where

- **Accept:** **allows** the **connection**.
- **Drop:** **does not allow** the **connection**, **does not send** back an **error**;
- **Reject:** **does not allow** the **connection**, **sends back** an **error**;

### IP failover
**Failover IPs** are **IPs** that can be **easily moved** from device to device. This makes them optimal for:
- networking VMs on a *Dedicated Server;
- load balancing;
- IP aliasing;
etc

![IP Failover Example](https://www.soyoustart.com/de/images/schema-ip6.png)

### TLS
**Transport Layer Security (TLS)** **encrypts data sent over the Internet** to ensure that eavesdroppers and hackers are unable to see what you transmit which is particularly useful for private and sensitive information such as *passwords, credit card numbers, and personal correspondence*.

### http2
**HTTP (Hypertext Transfer Protocol)** is the **protocol** that powers most of the web today. It is a **client-server** protocol.

**HTTP/2** is the **second version** of the **HTTP protocol** aiming to *make applications faster, simpler, and more robust* by improving many of the drawbacks of the first HTTP version.

The **primary goals** of **HTTP/2** are:
- **Enable request** and **response multiplexing**;
- **Header compression**;
- **Compatibility** with the *methods, status codes, URIs, and header fields* defined by the **HTTP/1.1 standard**;
- **Optimized prioritization of requests**, making sure that loading for optimal user experience is as fast as possible;
- **Support for server-side push**;
- **Server-side backwards compatibility**, making sure **servers** can still **serve clients** only supporting HTTP/1.1 without any changes;
- Transforming to a **binary protocol** from the **text-based HTTP/1.1**;

---

## Load Balancer
**Load balancing** refers to **efficiently distributing incoming network traffic** across a **group of backend servers**, also known as a server farm or server pool.

![Load Balancer Diagram](https://www.nginx.com/wp-content/uploads/2014/07/what-is-load-balancing-diagram-NGINX-640x324.png)

A load balancer:

- **Distributes client requests** or **network load** efficiently across **multiple servers**;
- **Ensures high availability** and **reliability** by **sending requests** only to **servers** that are **online**;
- Provides the **flexibility to add or subtract servers** as demand dictates;

### Load Balancing Algorithms

- **Round Robin:** **Requests** are **distributed** across the group of servers **sequentially**;
- **Least Connections:** A **new request** is **sent** to the **server** with the **fewest current connections** to clients;
- **Least Time:** – **Sends requests** to the **server** selected by a formula that **combines** the **fastest response time** and **fewest active connections**;
- **Hash:** **Distributes requests based on a key** you define, such as the *client IP address* or *the request URL*;
- **IP Hash:** The IP address of the client is used to determine which server receives the request.
- **Random with Two Choices:** **Picks two servers at random** and **sends the request** to the one that is **selected by** then applying the **Least Connections algorithm**;

---

## NAT Gateway
