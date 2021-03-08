package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	shape "github.com/ryanbeau/grpc-go-server/api/shape/v1"
	"google.golang.org/grpc"
)

const (
	address      = "localhost:50051"
	defaultSideA = uint32(2)
	defaultSideB = uint32(3)
	defaultSideC = uint32(4)
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := shape.NewShapeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get sides
	sideA := defaultSideA
	sideB := defaultSideB
	sideC := defaultSideC
	if len(os.Args) > 3 {
		if a, err := strconv.Atoi(os.Args[1]); err == nil {
			sideA = uint32(a)
		}
		if b, err := strconv.Atoi(os.Args[2]); err == nil {
			sideB = uint32(b)
		}
		if c, err := strconv.Atoi(os.Args[3]); err == nil {
			sideC = uint32(c)
		}
	}

	triangle := &shape.TrianglePerimeterRequest{
		SideA: sideA,
		SideB: sideB,
		SideC: sideC,
	}

	r, err := c.TrianglePerimeter(ctx, triangle)
	if err != nil {
		log.Fatalf("could not calculate perimeter: %v", err)
	}
	log.Printf("Perimeter: %d", r.Value)
}
