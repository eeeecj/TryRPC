package consumer

import (
	"fmt"
	"github.com/TryRpc/component/pkg/cuszap"
	"google.golang.org/grpc"
	"sync"
)

var (
	grpcClient *grpc.ClientConn
	once       sync.Once
)

func GetGRPC(config *ConsumerConfig) (*grpc.ClientConn, error) {
	once.Do(func() {
		var (
			err  error
			conn *grpc.ClientConn
		)
		//创建证书
		//creds, err = credentials.NewClientTLSFromFile(clentCA, "")
		//if err != nil {
		//	cuszap.Panicf("credentials.NewClientTLSFromFile err: %v", err)
		//}
		cuszap.Infof("connect to grpc:%s", config.Address())
		//连接grpc
		//conn, err = grpc.Dial("127.0.0.1:"+address, grpc.WithBlock(), grpc.WithTransportCredentials(creds))
		conn, err = grpc.Dial(config.Address(), grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			cuszap.Panicf("Connect to grpc server failed, error: %s", err.Error())
		}
		fmt.Println(conn)
		//注册服务
		cuszap.Infof("Connected to grpc server, address: %s", config.Address())
		grpcClient = conn
	})
	if grpcClient == nil {
		cuszap.Panicf("failed to get apiserver store fatory")
	}

	return grpcClient, nil
}
