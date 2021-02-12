package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/RaminCH_self/Gransoft/v2.1/encryptor/proto/consignment"
	a "github.com/RaminCH_self/Gransoft/v2/randomizer"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes"
)

type server struct{}

func (*server) Do(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Printf("Hash function received %v", req)
	str := req.GetStrSlice()
	result := NewSHA256(str)
	res := &pb.Response{
		Hash: result,
	}

	return res, nil
}

// NewSHA256 ...
func NewSHA256(data []*pb.String) *pb.String {
	data, err := ptypes.UnmarshalAny(a.ListOfStrings(n *pb.String, i int)
	if err != nil {
		fmt.Println(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

//Process function
func Process(i int, b *pb.String, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	// go Combine(a.ListOfStrings(a.RandomString(i), i))
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	fmt.Println("Server runs...")

	//connection
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEncryptServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

	//pool
	var wg sync.WaitGroup
	num := 4
	var b = NewSHA256(data []*pb.String)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go Process(i, b, &wg)
	}
	wg.Wait()
	fmt.Println(string(b))
	fmt.Println("All go routines finished executing")
}
