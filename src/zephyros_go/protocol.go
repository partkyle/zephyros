package zephyros_go

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"os"
)


func connect() net.Conn {
	conn, err := net.Dial("unix", "/tmp/zephyros.sock")
	if err != nil {
		fmt.Println("Can't connect. Is Zephyros running?")
		os.Exit(1)
	}
	return conn
}

var c net.Conn = connect()

var respChans = make(map[float64]chan []byte)

func ListenForCallbacks() {
	reader := bufio.NewReader(c)
	for {
		buf, _ := reader.ReadBytes('\n')

		// fmt.Printf("%#v\n", string(buf))

		var msg []interface{}
		json.Unmarshal(buf, &msg)

		// fmt.Println(msg)

		id, obj := msg[0].(float64), msg[1]
		bytes, _ := json.Marshal(obj)
		respChans[id] <- bytes
	}
}


var msgidChan chan float64 = make(chan float64)

func init() {
	go func() {
		var i float64 = 0
		for {
			i++
			msgidChan <- i
		}
	}()
}

func guardError(obj []byte) {
	if string(obj) == "__api_exception__" {
		panic("API Exception. (See above.)")
	}
}

func send(recv float64, fn func([]byte), infinite bool, method string, args ...interface{}) []byte {
	msgid := <- msgidChan

	ch := make(chan []byte, 10) // probably enough
	respChans[msgid] = ch

	msg := []interface{}{msgid, recv, method}
	val, _ := json.Marshal(append(msg, args...))
	jsonstr := string(val)
	fmt.Fprintln(c, jsonstr)

	if fn == nil {
		resp := <-ch
		delete(respChans, msgid)
		guardError(resp)
		return resp
	}

 	go func() {
		<-ch // ignore

		wrappedFn := func(obj []byte) {
			guardError(obj)
			fn(obj)
		}

		if infinite {
			for { wrappedFn(<-ch) }
		} else {
			wrappedFn(<-ch)
		}

		delete(respChans, msgid)
	}()

	return nil
}
