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

	ctx, contextCancelFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	connection, err := createConnection(ip, port, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	go sendMessage(ctx, connection, contextCancelFunc)
	go receiveMessage(ctx, connection, contextCancelFunc)

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

func sendMessage(ctx context.Context, client TelnetClient, contextCancelFunc context.CancelFunc) {
	select {
	case <-ctx.Done():
		return
	default:
		client.Send()
		contextCancelFunc()
		return
	}
}

func receiveMessage(ctx context.Context, client TelnetClient, contextCancelFunc context.CancelFunc) {
	select {
	case <-ctx.Done():
		return
	default:
		client.Receive()
		contextCancelFunc()
		return
	}
}
