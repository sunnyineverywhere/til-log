# 1. Data Communications, Data Networking, and the Internet

최종 수정: April 16, 2023

### 인터넷이란?

---

- 네트워크들의 네트워크
- 서로 연결되어있는 ISP들을 의미
- ISP: Internet Service Provider
- 접속된 컴퓨터(== 호스트)들의 집합
- LAN을 통하여 호스트들이 하나의 그룹으로 묶임 → 개별 호스트와 LAN이 ISP에 연결됨
- Infrastructure that provides services to applications

### 프로토콜

---

- Format, order of messages, and actions taken on message transmission and receipt
- 통신 규약: 통신하기 위한 절차를 규정한 것
- 방법, 절차, 언어 역할마다 계층 구조로 구분되어 있음
- 네트워크를 통해 컴퓨터끼리 통신하기 위한 규약

### Network edge

---

**Hosts**: 클라이언트와 서버(요청자와 수행자)

- 서버는 보통 데이터 센터에 있으며, 클라이언트의 요청을 수행한다

### **Access networks, physical media**

---

- 유선 / 무선 communication links
- Host(Edge Systems) ↔ Edge Router(Core)
    - Residential access nets 필요
    - access network가 중개
    - 모바일 access network 존재(WiFi, 4G/5G)
- Transmission rate(bits / sec)에 주의
- shared와 dedicated의 차이에 유의

**Cable-based Access**

- FDM(Frequency Division Multiplexing): 각 채널들이 각기 다른 frequency band를 가짐
- HFC(Hybrid Fiber Coax): 비대칭
    - down → 40 Mbps ~ 1.2 Gbps
    - up → 30-100 Mbps
    - Cable이나 Fiber는 가정의 ISP 라우터와 연결되어 있음
    - 가정(엔드 호스트)는 access network를 share
- DSL(Digital Subscriber Line)
    - 존재하는 전화유선을 DSLAM으로서 사용하여 연결
    - down → 24-52 Mbps
    - up → 3.5-16 Mbps

**Wireless access Network**

- Wireless local area networks(WLANs)
    - 빌딩 하나 안쪽
    - 802.11b/g/n/ac/ad (WiFi)
- Wide-area cellular access networks
    - mobile, cellular network 기기에 의해 제공됨(10km 안쪽)
    - 4G/5G cellular networks

**Enterprise Networks**

- Mixed of wired / wireless link technologies
- 다수의 스위치 / 라우터와 연결되어있다
    - Ethernet: wired access at 100Mbps, 1Gbps, 10Gbps (802.11)
    - WiFi: wireless access points at 11 / 54 / 450 Mbps (801.06)

### Host

---

호스트가 Send하는 과정

1. Application Message를 만듬
2. Message → packets 쪼갬 (L bits의 packet)
3. packet을 transmission (이때의 transmission rate = Link Bandwidth = R)

⇒ packet transmission delay == L비트의 패킷을 링크로 transmit하는데 걸리는 시간 == L(bits) / R(bits/sec)

### Link: Physical Media

---

Bit: tx와 rx 사이에서 propagates

Physical Link: tx와 rx 사이 이어진 길

Guided Media: 전선에서 signal(신호)가 propates → copper, fiber, coax

Unguided media: signals propagate freely

TP(Twisted Pair): Two insulated copper wires → Ethernet

- Coaxial cable
    - 양방향
    - Two concentric copper conductors
    - broadband: 케이블에 여러개의 frequency channel이 존재 → 채널당 100 Mbps
- Fiber optic cable
    - high-speed operation
    - low error rate
- Wireless radio
    - 전기적인 스펙트럼에서 데이터가 이동함
    - physical wire가 존재하지 않음
    - BroadCast → 단방향임(half-duplex)
- Radio link types
    - terrestiral microwave
    - wireless LAN(WiFi)
    - wide-area
    - satellite

### Network Core

---

- 내부에서 서로 연결된 router들의 집합
- packet-switching : 호스트들은 application layer에서 메세지를 packet으로 쪼갬
- 각 패킷은 라우터에서 다음 라우터로 forward됨(목표 host에 도착하기 전까지)
- 각 패킷은 full link capacity만큼 transmit됨

**Store-And-Forward**

![Untitled](https://file.notion.so/f/s/fccca2a6-208d-4e72-aef6-747485ab89c6/Untitled.png?id=7b65e586-5333-4750-987e-92b74e1af230&table=block&spaceId=c1344d5b-67e4-40b7-92d0-952b632930a9&expirationTimestamp=1684893759812&signature=gZuN6lmY96C_Zsw5I66Z-MH7XlhKDFNyH13XfhFFCsI&downloadName=Untitled.png)

- **Transmission Delay**: L/R sec → 패킷이 링크로 전송될 때
- Store & Forward: 모든 패킷은 다음 링크로 전송되기 전에 라우터에 도착해야 함
- E2E Delay: 2 * L / R ← assume! zero propagation delay

**Queueing Delay** 

![Untitled](1%20Data%20Communications,%20Data%20Networking,%20and%20the%20In%2071e5300900384a5cbf1d703101b9f26c/Untitled%201.png)

IF (arrival rate > transmission rate):

- 패킷이 queue에 들어가서 아웃풋 링크로 transmit 되기 전까지 대기
- 라우터의 메모리 버퍼가 다 차있으면 → 패킷들이 drop(lost)됨

**핵심 기능**

- Forwarding
    - local action
    - 도착한 패킷들은 라우터의 인풋 링크에서 아웃풋 링크로 옮김
- Routing
    - global action
    - 패킷들의 출발점(source)와 도착점(destination) path를 결정
    - routing algorithm이 필요

**Packet Switching vs. Circuit Switching**

| Packet Switching | Circuit Switching |
| --- | --- |
| No Call set-up (simple) | Requires Call set-up |
| Queuing Delay | Call set-up delay |
| No resource reservation
— Allocating link use on demand | Requires Resource reservation
regardless of demand |
| Possible to use full link rate |  |
| Resource sharing among multiple users | Resource dedicated to a particular user even while idle |
| Packet delay or loss
due to network congestion | Guaranteed service |
| More users can be served
depending on the network situation | # of users with connections at the same time is limited |
| Good: Intermittent, bursty data | Good: application which requires the guaranteed bandwidth |
| Dynamic use of bandwidth | Fixed use of bandwidth |

** *Circuit Switching*

- FDM(Frequency Division Multiplexing
    
    ![Untitled](1%20Data%20Communications,%20Data%20Networking,%20and%20the%20In%2071e5300900384a5cbf1d703101b9f26c/Untitled%202.png)
    

- TDM(Time Division Multiplexing)

![Untitled](1%20Data%20Communications,%20Data%20Networking,%20and%20the%20In%2071e5300900384a5cbf1d703101b9f26c/Untitled%203.png)
