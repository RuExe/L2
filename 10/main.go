package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var timeout = flag.Int("timeout", 10, "Таймаут")
	log.SetFlags(0)
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("Недостаточно аргументов")
		return
	}

	host := args[0]
	port := args[1]

	var (
		conn net.Conn
		err  error
	)

	start := time.Now()
	to := time.Duration(*timeout) * time.Second
	for time.Since(start) < to {
		conn, err = net.Dial("tcp", host+":"+port)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("smerd: %v", to)
	}
	defer conn.Close()
	log.Printf("success connected to %s:%s", host, port)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Тебе написали: " + message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(conn, scanner.Text()+"\n"); err != nil {
			log.Fatal("Закрыто")
		}
	}
}
