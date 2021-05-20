package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	"net/smtp"
)

var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	pin = rpio.Pin(22)
)
// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
   }// Address URI to smtp server
   func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
   }

func main() {
	// Open and map memory to access gpio, check for errors
	send("The water monitoring system is now up and running.")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()
	pin.PullUp()
	send()
	for {
		if pin.Read() == rpio.High {
			fmt.Println("Sensor has been activated, take action!!")
			//do a thing
		}
		time.Sleep(1 * time.Second)
	}
}
//Sends a message to us 
func send(msg string) {    
	s// Sender data.
    from := "42depressedrobot@gmail.com"
    password := "NSQPnCr7oz%79%b"    
	// Receiver email address.
    to := []string{
        "taylorkendall@gmail.com",
		"JenniferAntos@gmai.com"
    }    
	// smtp server configuration.
    smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}    
    message := []byte(msg)  
    auth := smtp.PlainAuth("", from, password, smtpServer.host)    
    err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)    if err != nil {
        fmt.Println(err)
        return
    }    fmt.Println("Email Sent!")
}
