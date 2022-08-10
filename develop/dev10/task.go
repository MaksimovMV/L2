package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeOut := flag.Duration("timeout", 10*time.Second, "таймаут на подключение к серверу")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("need host and port")
	}

	host := args[0]
	port := args[1]

	dial, err := net.DialTimeout("tcp", host+":"+port, *timeOut)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			// Чтение входных данных от stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			// Отправляем в socket
			fmt.Fprintf(dial, text+"\n")
			// Прослушиваем ответ
			message, err := bufio.NewReader(dial).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Print("Message from server: " + message)
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGHUP,
	)

	go func() {
		<-sign
		fmt.Println("Программа завершена вручную")
		dial.Close()
		wg.Done()
	}()
	wg.Wait()
}
