package timeserver
import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		func() {
			_, err := io.WriteString(conn, time.Now().Format("15:05:05"))
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}()
	}

}