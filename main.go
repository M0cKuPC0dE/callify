package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usb Device not define.")
		return
	}

	c := &serial.Config{Name: os.Args[1], Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(s)
	for {
		reply, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(reply))
		if strings.HasPrefix(string(reply), "+CLIP") {
			_, err := s.Write([]byte("ath"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
