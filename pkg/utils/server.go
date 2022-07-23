package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/TryRpc/component/pkg/cuszap"
	"io/ioutil"
	"log"
	"net"
)

func CreateListenWithTLS(addr string, cafile string, cakey string, CA string) (net.Listener, error) {
	fmt.Println(addr, cafile, cakey, CA)
	//tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	//if err != nil {
	//	return nil, err
	//}
	cert, err := tls.LoadX509KeyPair(cafile, cakey)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CA)
	if err != nil {
		return nil, err
	}
	cuszap.Debug("import ca")
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	cuszap.Info("listening on " + addr)
	lis, err := tls.Listen("tcp", addr, config)
	//lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	cuszap.Info("listening on " + addr)
	return lis, nil
}

func CreateListen(addr string) (net.Listener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	cuszap.Info("listening on " + addr)
	lis, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return lis, nil
}

func CreateConn(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func CreateConnWithTLS(addr string, cafile string, cakey string, CA string) (net.Conn, error) {
	fmt.Println(cafile, cakey, CA)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	certificate, err := tls.LoadX509KeyPair(cafile, cakey)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CA)
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "server.io",
		RootCAs:      certPool,
	}
	conn, err := tls.Dial("tcp", tcpAddr.String(), config)
	if err != nil {
		return nil, err
	}
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn, err
}
