package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/bigtable"
	"github.com/golang/protobuf/ptypes/empty"
	hello "github.com/weathersource/test-hello/proto"
	context "golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

// server is used to implement hello.HelloService.
type server struct{}

type Clients struct {
	cbtClient *bigtable.Client
}

var clients *Clients

// SetOnpoints implements gRPC hello.HelloService
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

	time.Sleep(30 * time.Second)

	// get service account credentials for BigTable
	var cbtConfig *jwt.Config
	if 0 != len(os.Getenv("password")) {
		var err error
		jsonStr := os.Getenv("password")
		jsonKey := []byte(jsonStr)

		cbtConfig, err = google.JWTConfigFromJSON(jsonKey, bigtable.Scope)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		log.Println(errors.New("Failed to retrieve service account for BigTable client."))
		os.Exit(1)
	}

	// create BigTable client
	log.Println("Configuring cbtClient.")
	cbtClient, err := bigtable.NewClient(
		context.Background(),
		"ws-microservices-production",
		"legolas-production",
		option.WithTokenSource(cbtConfig.TokenSource(context.Background())),
	)
	log.Println("SUCCESS configuring cbtClient.")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// set global clients
	clients = &Clients{
		cbtClient: cbtClient,
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
