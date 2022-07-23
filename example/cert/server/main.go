package main

import (
	"bufio"
	"context"
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"github.com/TryRpc/pkg/utils"
	"net/http"
)

type Hello struct {
}

func (h *Hello) Hello(context.Context, *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Output: "hello"}, nil
}

//func main() {
//	cred, err := tls.LoadX509KeyPair("./config/certs/server.crt", "./config/certs/server.key")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool := x509.NewCertPool()
//	ca, err := ioutil.ReadFile("./config/certs/ca.crt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if ok := certPool.AppendCertsFromPEM(ca); !ok {
//		log.Fatal("Fail to append certs")
//	}
//
//	creds := credentials.NewTLS(&tls.Config{
//		Certificates: []tls.Certificate{cred},
//		ClientAuth:   tls.RequireAndVerifyClientCert,
//		ClientCAs:    certPool,
//	})
//
//	server := grpc.NewServer(grpc.Creds(creds))
//	hello.RegisterHelloServer(server, new(Hello))
//	lis, err := net.Listen("tcp", ":12345")
//	if err != nil {
//		log.Fatal(err)
//	}
//	server.Serve(lis)
//}

func main() {
	lis, err := utils.CreateListenWithTLS("127.0.0.1:12345", "./config/certs/server.crt",
		"./config/certs/server.key", "./config/certs/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}
		buf, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println(string(buf))
	}
}

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		" Hi, This is an example of https service in golang!\n")
}

//func main() {
//	pool := x509.NewCertPool()
//	caCertPath := "./config/certs/ca.crt"
//
//	caCrt, err := ioutil.ReadFile(caCertPath)
//	if err != nil {
//		fmt.Println("ReadFile err: ", err)
//		return
//	}
//	pool.AppendCertsFromPEM(caCrt)
//
//	s := &http.Server{
//		Addr:    ":8088",
//		Handler: &myhandler{},
//		TLSConfig: &tls.Config{
//			ClientCAs:  pool,
//			ClientAuth: tls.RequireAndVerifyClientCert,
//		},
//	}
//
//	fmt.Println("listen...")
//	err = s.ListenAndServeTLS("./config/certs/server.crt", "./config/certs/server.key")
//	if err != nil {
//		fmt.Println(err)
//	}
//}
