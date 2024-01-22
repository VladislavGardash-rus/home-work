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

	ctx, contextCancelFunc := createContext(ip, port, timeout)

	connection, err := createConnection(ip, port, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	go sendMessage(connection, contextCancelFunc)
	go receiveMessage(connection, contextCancelFunc)

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

func createContext(ip, port string, timeout *time.Duration) (context.Context, context.CancelFunc) {
	ctx, ctxCancelFunc := context.WithCancel(context.Background())
	ctx, ctxCancelFunc = signal.NotifyContext(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return ctx, func() {
		fmt.Println("...Connection was closed by peer")
		ctxCancelFunc()
	}
}

func createConnection(ip, port string, timeout *time.Duration) (TelnetClient, error) {
	client := NewTelnetClient(fmt.Sprintf("%s:%s", ip, port), *timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func sendMessage(client TelnetClient, contextCancelFunc context.CancelFunc) {
	defer contextCancelFunc()
	err := client.Send()
	if err != nil {
		log.Println(err)
	}
}

func receiveMessage(client TelnetClient, contextCancelFunc context.CancelFunc) {
	defer contextCancelFunc()
	err := client.Receive()
	if err != nil {
		log.Println(err)
	}
}
