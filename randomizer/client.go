package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "github.com/RaminCH_self/Gransoft/v2.1/randomizer/proto/consignment"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	fmt.Println("Client runs...")
	rn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer rn.Close()

	cl := pb.NewEncryptClient(rn)
	fmt.Printf("Created client: %f", cl)

	doUnary(cl)
}

func doUnary(cl pb.EncryptClient) {

	fmt.Println("Starting Unary RPC...")
	list := ListOfStrings(RandomString(4), 4)l
	// StrSlice := make([]string, 10)
	req := &pb.Request{
		StrSlice: []*pb.String{
			StrSlice: list,
		},
	}

	res, err := cl.Do(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Do RPC: %v", err)
	}
	log.Printf("Response from Hash: %v", res.Hash)

}

//RandomString function ...
func RandomString(n int) *pb.String {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	a, err := ptypes.MarshalAny(s)
	if err != nil {
		log.Fatal(err)
	}

	return a
}

//ListOfStrings func (receives random strings and generates a slice of strings)
func ListOfStrings(n *pb.String, i int) []*pb.String {
	var clsStr []*pb.String

	strng := RandomString(i)
	clsStr = append(clsStr, strng)
	return clsStr
}
