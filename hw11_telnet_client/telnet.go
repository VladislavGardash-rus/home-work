package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"syscall"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	log.Println("...Connected to ", c.address)
	return nil
}

func (c *telnetClient) Close() error {
	return c.conn.Close()
}

func (c *telnetClient) Send() error {
	r := bufio.NewReader(c.in)
	for {
		message, err := r.ReadString('\n')
		if err != nil {
			printLogMessage(err)
			return err
		}

		_, err = c.conn.Write([]byte(message))
		if err != nil {
			return err
		}
	}
}

func (c *telnetClient) Receive() error {
	r := bufio.NewReader(c.conn)
	for {
		message, err := r.ReadString('\n')
		if err != nil {
			printLogMessage(err)
			return err
		}

		_, err = c.out.Write([]byte(message))
		if err != nil {
			return err
		}
	}
}

func printLogMessage(err error) {
	if errors.Is(err, io.EOF) {
		log.Println("...EOF")
	}

	if errors.Is(err, syscall.ECONNRESET) {
		log.Print("...Connection was closed by peer")
	}
}
