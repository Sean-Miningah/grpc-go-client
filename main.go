package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"time"

	pb "github./protobuffers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Load the certificate for the server
	cert, err := os.ReadFile("path/to/your/cert.pem")
	if err != nil {
		log.Print("Cannot Read Server Certificate")
	}

	// Create a certificate pool and add the server's certificate
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPem(cert) {
		log.Print("Could not find the certificates")
	}

	// Create the TLS credentials for the client
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})

	// Dial the server with TLS credentials
	conn, err := grpc.Dial("dev-analytics.dialafrika.com:80", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Print("Could not connect to server")
	}
	defer conn.Close()

	// Create gRPC client client :=
	client := pb.NewCuratedAnalyticsClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Prepare the request
	request := &pb.TotalAccLeadContactDealRequest{}

	// Make the gRPC call
	response, err := client.GetTotalAccLeadContactDeal(ctx, request)
	if err != nil {
		log.Fatal("RPC failed: %v", err)
	}

	fmt.Println(response)
}
