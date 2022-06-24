package main
import("fmt";"os";"net";"log")

const(
	HOST = "localhost"
	PORT = "9001"
	TYPE = "tcp"

	BUF_SZ = 1024
)

func incomingSession(ses net.Conn) {
	// 存送過來的資料
	buffer := make([]byte, BUF_SZ)
	_, err := ses.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	msg := string(buffer[:])
	fmt.Println(msg)
	
	ses.Close()
}

func main() {
	fmt.Println("Hello!!");
	listen,err := net.Listen(TYPE, HOST + ":" + PORT)
	defer listen.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go incomingSession(conn)
	}
}
