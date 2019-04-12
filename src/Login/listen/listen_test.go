package listen

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"proto/login"
	"testing"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial("192.168.5.14:10001", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := login.NewGameLoginServiceClient(conn)
	in := new(login.Request)
	in.Login = "root"
	in.Password = "password"
	out, err := client.GameLogin(context.Background(), in)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
func Test1(t *testing.T){
	conn,err:=net.Dial("tcp","localhost:10002")
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	for{
		var by = make([]byte,1024)
		conn.Read(by)
		fmt.Println(string(by))
		input:=bufio.NewScanner(os.Stdin)
		input.Scan()
		conn.Write([]byte(input.Text()))
	}

}