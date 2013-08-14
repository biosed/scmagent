package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"github.com/garyburd/redigo/redis"
)

func main() {




	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Printf("SCM Agent started...\n")
	Example_zpop()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.WriteString(c, "Hello")
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}


func Example_zpop() {
    c, err := redis.Dial("tcp","10.109.64.23:6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer c.Close()

    // Add test data using a pipeline.
       _,err =  c.Do("SELECT", "10")
      reply, err :=  redis.String(c.Do("GET", "plugins-httpreq"))
       fmt.Println(reply)
    if _, err := c.Do(""); err != nil {
        fmt.Println(err)
        return
    }

    // blue
}