package main

import (
	"context"
	"fmt"

	grpc "google.golang.org/grpc"
)

const (
	// какой адрес-порт слушать серверу
	listenAddr string = "127.0.0.1:1323"
)

// утилитарная функция для коннекта к серверу
func getGrpcConn() *grpc.ClientConn {
	grcpConn, err := grpc.Dial(
		listenAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Printf("getGrpcConn error: %v", err.Error())
		return nil
	}
	return grcpConn
}

func main() {

	conn := getGrpcConn()
	if conn == nil {
		return
	}
	defer conn.Close()

	catClient := NewCatsExplorerClient(conn)

	catId := new(Id)
	catId.Id = 2
	cat, err := catClient.GetCat(context.Background(), catId)
	if err != nil {
		fmt.Printf("GetCat error: %v", err.Error())
	}
	fmt.Println(cat)
}
