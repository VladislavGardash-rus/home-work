package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	scanner := bufio.NewScanner(c.in)
	for scanner.Scan() {
		err := writeMessage(c.conn, scanner.Text())
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *telnetClient) Receive() error {
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		err := writeMessage(c.out, scanner.Text())
		if err != nil {
			return err
		}
	}
	return nil
}

func writeMessage(w io.Writer, message string) error {
	if message == "" || []byte(message)[0] == 4 {
		fmt.Println("...EOF")
		return nil
	}

	_, err := w.Write([]byte(fmt.Sprintf("%s\n", message)))
	if err != nil {
		return err
	}

	return nil
}
