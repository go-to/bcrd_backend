package router

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/go-to/bcrd_backend/usecase"
	"github.com/go-to/bcrd_backend/usecase/input"
	"github.com/go-to/bcrd_protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedBcrdServiceServer
	Usecase Usecase
}

type Usecase struct {
	Shop  usecase.IShopUsecase
	Stamp usecase.IStampUsecase
}

func New(port int, u Usecase) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterBcrdServiceServer(s, NewServer(u))

	reflection.Register(s)

	go func() {
		log.Printf("grpc server listening on port %d", port)
		err := s.Serve(listener)
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutting down server...")
	s.GracefulStop()
}

func NewServer(u Usecase) *Server {
	return &Server{Usecase: u}
}

func checkApiKey(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing in request")
	}

	apiKey := md.Get("api-key")
	if len(apiKey) == 0 || apiKey[0] != os.Getenv("API_KEY") {
		return status.Errorf(codes.Unauthenticated, "invalid api key")
	}
	return nil
}

func (s *Server) GetShopsTotal(ctx context.Context, req *pb.ShopsTotalRequest) (*pb.ShopsTotalResponse, error) {
	if err := checkApiKey(ctx); err != nil {
		return nil, err
	}

	in := input.ShopsTotalInput{
		ShopsTotalRequest: req,
	}

	out, err := s.Usecase.Shop.GetShopsTotal(&in)
	if err != nil {
		return nil, err
	}

	return &out.ShopsTotalResponse, nil
}

func (s *Server) GetShops(ctx context.Context, req *pb.ShopsRequest) (*pb.ShopsResponse, error) {
	if err := checkApiKey(ctx); err != nil {
		return nil, err
	}

	in := input.ShopsInput{
		ShopsRequest: req,
	}

	t := time.Now()
	out, err := s.Usecase.Shop.GetShops(&in)
	if err != nil {
		return nil, err
	}
	fmt.Printf("process time: %s\n", time.Since(t))

	return &out.ShopsResponse, nil
}

func (s *Server) GetShop(ctx context.Context, req *pb.ShopRequest) (*pb.ShopResponse, error) {
	if err := checkApiKey(ctx); err != nil {
		return nil, err
	}

	in := input.ShopInput{
		ShopRequest: req,
	}

	out, err := s.Usecase.Shop.GetShop(&in)
	if err != nil {
		return nil, err
	}

	return &out.ShopResponse, nil
}

func (s *Server) AddStamp(ctx context.Context, req *pb.StampRequest) (*pb.StampResponse, error) {
	if err := checkApiKey(ctx); err != nil {
		return nil, err
	}

	in := input.StampInput{
		StampRequest: req,
	}

	out, err := s.Usecase.Stamp.AddStamp(&in)
	if err != nil {
		return nil, err
	}

	return &out.StampResponse, nil
}

func (s *Server) DeleteStamp(ctx context.Context, req *pb.StampRequest) (*pb.StampResponse, error) {
	if err := checkApiKey(ctx); err != nil {
		return nil, err
	}

	in := input.StampInput{
		StampRequest: req,
	}

	out, err := s.Usecase.Stamp.DeleteStamp(&in)
	if err != nil {
		return nil, err
	}

	return &out.StampResponse, nil
}
