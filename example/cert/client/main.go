package main

import (
	"fmt"
	"github.com/TryRpc/pkg/utils"
)

func main() {
	conn, err := utils.CreateConnWithTLS("127.0.0.1:12345", "./config/certs/client.crt",
		"./config/certs/client.key", "./config/certs/ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("hello\n"))
}

//func main() {
//	certificate, err := tls.LoadX509KeyPair("./config/certs/client.crt", "./config/certs/client.key")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool := x509.NewCertPool()
//	ca, err := ioutil.ReadFile("./config/certs/ca.crt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	if ok := certPool.AppendCertsFromPEM(ca); !ok {
//		log.Fatal(err)
//	}
//	creds := credentials.NewTLS(&tls.Config{
//		Certificates: []tls.Certificate{certificate},
//		ServerName:   "server.io",
//		RootCAs:      certPool,
//	})
//
//	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	client := hello.NewHelloClient(conn)
//
//	resp, err := client.Hello(context.Background(), &hello.HelloRequest{Input: "xiaoming "})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(resp)
//}

//func main() {
//	// x509.Certificate
//	pool := x509.NewCertPool()
//
//	caCertPath := "./config/certs/ca.crt"
//	caCrt, err := ioutil.ReadFile(caCertPath)
//	if err != nil {
//		fmt.Println("ReadFile err:", err)
//		return
//	}
//	pool.AppendCertsFromPEM(caCrt)
//
//	cliCrt, err := tls.LoadX509KeyPair("./config/certs/client.crt", "./config/certs/client.key")
//	if err != nil {
//		fmt.Println("LoadX509keypair err: ", err)
//		return
//	}
//
//	//    tr := &http2.Transport{  // http2协议
//	tr := &http.Transport{ // http1.1协议
//		TLSClientConfig: &tls.Config{
//			RootCAs:      pool,
//			ServerName:   "server.io",
//			Certificates: []tls.Certificate{cliCrt},
//		},
//	}
//	client := &http.Client{Transport: tr}
//
//	//resp, err := client.Get("https://localhost:8088")
//	resp, err := client.Get("https://localhost:8088")
//	if err != nil {
//		fmt.Println("http get error: ", err)
//		panic(err)
//	}
//
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body))
//	fmt.Println(resp.Status)
//}
