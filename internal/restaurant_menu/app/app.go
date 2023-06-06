package app

import (
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"restaurant/internal/restaurant_menu/bootstrap"
	"restaurant/internal/restaurant_menu/config"
	"restaurant/internal/restaurant_menu/consumers"
	"restaurant/internal/restaurant_menu/repositories/menurepository/menurepositorygorm"
	"restaurant/internal/restaurant_menu/repositories/orderrepository/orderrepositorygrpc"
	"restaurant/internal/restaurant_menu/repositories/productrepository/productrepositorygorm"
	services "restaurant/internal/restaurant_menu/services"
	customer "restaurant/pkg/contracts/customer_office"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	"syscall"
	"time"
)

func Run(cfg config.Config) error {
	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go runGRPCServer(cfg, s)
	go runHTTPServer(ctx, cfg, mux)

	gracefulShutDown(s, cancel)

	return nil

}

func runGRPCServer(cfg config.Config, grpcServer *grpc.Server) {
	db, err := bootstrap.InitGormDB(cfg)
	if err != nil {
		log.Fatalf("Gorm init: %s", err.Error())
	}
	err = bootstrap.MigrateSchema(db)
	if err != nil {
		log.Fatalf("Migrate schema: %s", err.Error())
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("Net listening: %s", err.Error())
	}
	amqpConn, err := bootstrap.InitAMQPConnection(cfg)
	if err != nil {
		log.Fatalf("AMPQ connection: %s", err.Error())
	}
	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		log.Fatalf("AMPQ channel: %s", err.Error())
	}
	orderConsumer := consumers.NewOrderConsumerAmqp(amqpChannel, cfg.AMQPOrderQueue)
	go orderConsumer.Handle()
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	customerOfficeGRPCConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.CustomerOfficeHost, cfg.CustomerOfficePort), grpcOpts...)
	if err != nil {
		log.Fatalf("Customer Office GRPC connection: %s", err.Error())
	}
	orderServiceClient := customer.NewOrderServiceClient(customerOfficeGRPCConn)
	menuRepository := menurepositorygorm.New(db)
	orderRepository := orderrepositorygrpc.New(orderServiceClient)
	productRepository := productrepositorygorm.New(db)
	menuService := services.NewMenuServiceGRPCServer(menuRepository, productRepository)
	orderService := services.NewOrderServiceGRPCServer(orderRepository, productRepository)
	productService := services.NewProductServiceGRPCServer(productRepository)
	restaurant.RegisterMenuServiceServer(grpcServer, menuService)
	restaurant.RegisterOrderServiceServer(grpcServer, orderService)
	restaurant.RegisterProductServiceServer(grpcServer, productService)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("GRPC Server serving: %s", err.Error())
	}
}

func runHTTPServer(
	ctx context.Context, cfg config.Config, mux *runtime.ServeMux,
) {
	err := restaurant.RegisterMenuServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("HTTP Menu service handling: %s", err)
	}
	err = restaurant.RegisterOrderServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("HTTP Order service handling: %s", err)
	}
	err = restaurant.RegisterProductServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("HTTP Product service handling: %s", err)
	}
	log.Printf("starting listening http server at %s", fmt.Sprintf(":%d", cfg.HttpPort))
	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), mux); err != nil {
		log.Fatalf("error service http server %v", err)
	}
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	const waitTime = 5 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	log.Println(errorMessage)
	s.GracefulStop()
	cancel()
}
