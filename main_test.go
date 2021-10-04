package main

import (
	"bufio"
	"log"
	"net"
	"testing"
	"time"
)

func Test_randProverbIndex(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Кейс 1",
			args: args{
				max: 19,
			},
		}, {
			name: "Кейс 2",
			args: args{
				max: 0,
			},
		}, {
			name: "Кейс 3",
			args: args{
				max: 5,
			},
		}, {
			name: "Кейс 4",
			args: args{
				max: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randProverbIndex(tt.args.max); got > tt.args.max && tt.args.max >= 0 {
				t.Errorf("Случайное число randInt() = %v больше верхнего диапозона max: %v", got, tt.args.max)
			}
		})
	}
}

func Test_handleConn(t *testing.T) {
	srv, cl := net.Pipe()

	go func() {
		handleConn(srv)
	}()

	// Выполняем чтение сообщений от сервера в течение 20 секунд, чтобы не сработал
	// таймаут и отработал тест. Для более длительного времени выполнения тест
	// можно запустить с большим значением параметра -timeout

	ticker := time.Tick(20 * time.Second)
	for {
		select {
		case <-ticker:
			cl.Close()
			srv.Close()
			return
		default:
			reader := bufio.NewReader(cl)
			b, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			// Проверка ответа.
			if len(b) == 0 {
				t.Error("Поговорка не получена")
			}
			t.Logf("Получена поговорка: %v", string(b))
		}
	}
}
