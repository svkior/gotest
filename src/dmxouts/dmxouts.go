package dmxouts

import (
	"log"
	"net"
	"time"
)


type dmxOuts struct {
	name string
}


func Create() (dmxOuts) {
	var rv dmxOuts
	rv.name = "DMX Outputs"
	return rv
}

type ArtPollPacket struct {
	id 			[8]byte		// 00 .. 07 ID Of ArtNet
	OpCode 		int16		// 08, 09opPoll
	ProtVerHi	byte		// 10
	ProvVerLo	byte		// 11
	TalkToMe	byte		// 12
	Priority	byte		// 13
}


func myUDPClient(){

	addr := net.UDPAddr{
		Port: 9229,
		IP: net.ParseIP("127.0.0.1"),
	}

	broadcast := net.UDPAddr {
		Port: 9229,
		IP: net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &broadcast)
	defer conn.Close()
	if err != nil {
		log.Panic(err)
	}


	pingTicker := time.NewTicker(time.Second)
	defer pingTicker.Stop()

	// Do something with `conn`
	wb := []byte("              ")

	wb[0] = 'A'
	wb[1] = 'r'
	wb[2] = 't'
	wb[3] = '-'
	wb[4] = 'N'
	wb[5] = 'e'
	wb[6] = 't'
	wb[7] = 0x00
	wb[8] = 0x00
	wb[9] = 0x20
	wb[10] =0x03
	wb[11] = 14
	wb[12] = 0
	wb[13] = 0

	for {
		select {
		case <- pingTicker.C:
			conn.WriteMsgUDP([]byte(wb), nil, &addr)
		}
	}

}


func myUDPServer() {
	addr := net.UDPAddr{
		Port: 9229,
		IP: net.ParseIP("127.0.0.1"),
	}

	broadcast := net.UDPAddr {
		Port: 9229,
		IP: net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &broadcast)
	defer conn.Close()
	if err != nil {
		log.Panic(err)
	}


	pingTicker := time.NewTicker(time.Second)
	defer pingTicker.Stop()

	// Do something with `conn`
	wb := []byte("              ")

	wb[0] = 'A'
	wb[1] = 'r'
	wb[2] = 't'
	wb[3] = '-'
	wb[4] = 'N'
	wb[5] = 'e'
	wb[6] = 't'
	wb[7] = 0x00
	wb[8] = 0x00
	wb[9] = 0x20
	wb[10] =0x03
	wb[11] = 14
	wb[12] = 0
	wb[13] = 0

	for {
		select {
		case <- pingTicker.C:
			conn.WriteMsgUDP([]byte(wb), nil, &addr)
		}
	}


}

func (dmxOuts) Run(){
	log.Println("dmxOuts: GO!")
	go myUDPServer()
}
