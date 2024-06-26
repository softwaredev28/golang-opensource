package main

import (
	"client-unsia/pb/cities"
	"fmt"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %s", err)
		return
	}
	defer conn.Close()

	ctx := context.Background()

	city := cities.NewCitiesServiceClient(conn)

	// response, err := city.GetCity(ctx, &cities.City{
	// 	Id: 1,
	// })

	err = callStream(ctx, city)

	if err != nil {
		fmt.Printf("Error when calling grpc service: %s", err)
		return
	}
	// fmt.Printf("Resp received: %v", response)
}

func callStream(ctx context.Context, city cities.CitiesServiceClient) error {
	stream, err := city.GetCities(ctx, &cities.EmptyMessage{})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end stream")
			break
		}
		if err != nil {
			fmt.Printf("cannot receive %v", err)
			return err
		}
		fmt.Printf("Resp : %v", resp.GetCity())
		println()
	}

	return nil
}