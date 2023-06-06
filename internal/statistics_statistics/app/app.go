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
	"restaurant/internal/statistics_statistics/bootstrap"
	"restaurant/internal/statistics_statistics/config"
	"restaurant/internal/statistics_statistics/consumers"
	"restaurant/internal/statistics_statistics/repositories/orderrepository/orderrepositorysqlx"
	"restaurant/internal/statistics_statistics/repositories/productrepository/productrepositorygrpc"
	services "restaurant/internal/statistics_statistics/services"
	restaurant "restaurant/pkg/contracts/restaurant_menu"
	statistics "restaurant/pkg/contracts/statistics_statistics"
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
	db, err := bootstrap.InitSqlxDB(cfg)
	if err != nil {
		log.Fatalf("Sqlx init: %s", err.Error())
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
	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	restaurantMenuGRPCConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.RestaurantMenuHost, cfg.RestaurantMenuPort), grpcOpts...)
	if err != nil {
		log.Fatalf("Restaurant Menu grpc connection: %s", err.Error())
	}
	productServiceClient := restaurant.NewProductServiceClient(restaurantMenuGRPCConn)
	orderRepository := orderrepositorysqlx.New(db)
	productRepository := productrepositorygrpc.New(productServiceClient)

	orderConsumer := consumers.NewOrderConsumerAmqp(amqpChannel, cfg.AMQPOrderQueue, orderRepository, productRepository)
	go orderConsumer.Handle()

	statisticsService := services.NewStatisticsGRPCService(orderRepository)
	statistics.RegisterStatisticsServiceServer(grpcServer, statisticsService)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("GRPC Server serving: %s", err.Error())
	}
}

func runHTTPServer(
	ctx context.Context, cfg config.Config, mux *runtime.ServeMux,
) {
	err := statistics.RegisterStatisticsServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("HTTP Statistics service handling: %s", err)
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
