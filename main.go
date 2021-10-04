package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

// Адрес сетевой службы
const address string = "127.0.0.1:54321"

// Протокол службы
const proto string = "tcp4"

// Интервал отправки поговорок клиенту в секундах
var interval time.Duration = 3 * time.Second

// Поговорки
var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	l, err := net.Listen(proto, address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	for {
		if len(proverbs) == 0 {
			continue
		}
		i := randProverbIndex(len(proverbs) - 1)
		conn.Write([]byte(proverbs[i] + "\n"))
		time.Sleep(interval)
	}
}

// Возвращает случайный индекс поговорки
func randProverbIndex(max int) int {
	if max <= 0 {
		return 0
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
}
