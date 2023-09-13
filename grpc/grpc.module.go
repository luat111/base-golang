package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"practice/auth/core/config"
	. "practice/auth/core/constants"
)

var GRPCMap map[string]*grpc.ClientConn
var ServiceMap map[string]string

func InitService(address string) error {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", address, err)
	}

	GRPCMap[address] = conn
	return nil
}

func InitClientGRPC(config config.AppConfig) (map[string]*grpc.ClientConn, map[string]string) {

	blogAddress := config.GRPCSetting.BlogServiceHost + ":" + config.GRPCSetting.BlogServicePort

	address := []string{blogAddress}
	GRPCMap = make(map[string]*grpc.ClientConn)
	ServiceMap = make(map[string]string)

	ServiceMap[BlogService] = blogAddress
	for _, value := range address {
		InitService(value)
	}

	return GRPCMap, ServiceMap
}
