package term2

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

type Term2 struct {
	w1  *net.Conn
	out chan string
}

func usage() {
	fmt.Fprintf(os.Stderr, "launch terminal ex: ncat -l 7070\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func New(port string) *Term2 {
	host := strings.Builder{}
	host.WriteString("127.0.0.1")
	host.WriteString(":")
	host.WriteString(port)

	t2conn, err := net.Dial("tcp", host.String())
	if err != nil {
		fmt.Println(err)
		usage()
	}
	//defer t2conn.Close()
	t2 := &Term2{
		w1:  &t2conn,
		out: make(chan string),
	}
	go func() {
		for elem := range t2.out {
			fmt.Fprint(*t2.w1, elem)
		}
	}()
	return t2
}

func (t2 *Term2) Prn(str string) *Term2 {
	t2.out <- str
	return t2
}

func (t2 *Term2) Prnl(str string) *Term2 {
	str += "\n"
	t2.Prn(str)
	return t2
}

//Clear screen, cursor on top
func (t2 *Term2) Cls() *Term2 {
	t2.Prn("\033[H\033[2J")
	return t2
}
