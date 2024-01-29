package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ip, port, timeout, err := getConnectionParams()
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)
	connection, err := createConnection(ip, port, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	go sendMessage(connection)
	go receiveMessage(connection)

	<-ctx.Done()
}

func getConnectionParams() (string, string, *time.Duration, error) {
	timeout := flag.Duration("timeout", 10*time.Second, "server connect timeout")
	flag.Parse()
	if flag.NArg() != 2 {
		return "", "", nil, errors.New("ip and port are required params")
	}

	return flag.Arg(0), flag.Arg(1), timeout, nil
}

func createConnection(ip, port string, timeout *time.Duration) (TelnetClient, error) {
	client := NewTelnetClient(fmt.Sprintf("%s:%s", ip, port), *timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func sendMessage(client TelnetClient) {
	err := client.Send()
	if err != nil {
		os.Exit(0)
	}
}

func receiveMessage(client TelnetClient) {
	err := client.Receive()
	if err != nil {
		os.Exit(0)
	}
}
