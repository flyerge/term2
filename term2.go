package term2

import (
	"fmt"
	"log"
	"net"
)

type Term2 struct {
	w1  *net.Conn
	out chan string
}

func New() *Term2 {
	t2conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	//defer t2conn.Close()
	return &Term2{
		w1:  &t2conn,
		out: make(chan string),
	}
}

func (t2 *Term2) Start() {
	for elem := range t2.out {
		fmt.Fprint(*t2.w1, elem)
	}
}

func (t2 *Term2) Prn(str string) {
	t2.out <- str
}
