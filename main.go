package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	hello "github.com/weathersource/test-hello/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

// server is used to implement hello.HelloService.
type server struct{}

var clients *Clients

// SetOnpoints implements hello.HelloService
func (s *server) SayHello(ctx context.Context, in *empty.Empty) (*hello.SayHelloResponse, error) {
	// TODO: use clients to access BigTable
	return &hello.SayHelloResponse{Msg: "hello"}, nil
}

// sayHealthy demonstrates health of the server by returning "Healthy" to http requests
func sayHealthy(w http.ResponseWriter, r *http.Request) {
	message := "Healthy"
	w.Write([]byte(message))
}

func main() {
	// instantiate clients object
	var e error
	clients, e = NewClients()
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// start GRPC server
	go func() {
		defer wg.Done()
		log.Println("Serve gRPC.")
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Println(err)
			return
		}

		s := grpc.NewServer(grpc.MaxRecvMsgSize(100 * 1024 * 1024))
		hello.RegisterHelloServiceServer(s, &server{})
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Println(err)
			return
		}
	}()

	// start health check HTTP server
	go func() {
		defer wg.Done()
		log.Println("Serve heath check")
		http.HandleFunc("/", sayHealthy)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println(err)
			return
		}
	}()

	wg.Wait()
	os.Exit(1)
}
