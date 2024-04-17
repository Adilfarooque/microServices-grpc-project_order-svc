package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Adilfarooque/microServices-grpc-project_order_svc/pkg/client"
	"github.com/Adilfarooque/microServices-grpc-project_order_svc/pkg/config"
	"github.com/Adilfarooque/microServices-grpc-project_order_svc/pkg/db"
	"github.com/Adilfarooque/microServices-grpc-project_order_svc/pkg/pb"
	"github.com/Adilfarooque/microServices-grpc-project_order_svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
	}
	h := db.Init(c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	list, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
