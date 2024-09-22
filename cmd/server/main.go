package main

import (
	"fmt"
	"xyz-consumer-service/common/config"
	gormConn "xyz-consumer-service/common/gorm"
	commonJwt "xyz-consumer-service/common/jwt"
	"xyz-consumer-service/common/mysql"
	"xyz-consumer-service/server"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	splash(cfg)

	dsn, derr := mysql.NewPool(&cfg.MySQL)
	checkError(derr)

	db, gerr := gormConn.NewMySQLGormDB(dsn)
	checkError(gerr)

	jwtManager := commonJwt.NewJWT(cfg.JWT.JwtSecretKey, cfg.JWT.TokenDuration)

	grpcServer := server.NewGrpcServer(cfg.Port.GRPC, jwtManager)
	grpcConn := server.InitGRPCConn(fmt.Sprintf("127.0.0.1:%v", cfg.Port.GRPC), false, "")

	registerGrpcHandlers(grpcServer.Server, *cfg, db, grpcConn)

	_ = grpcServer.Run()
	_ = grpcServer.AwaitTermination()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func registerGrpcHandlers(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	// postModule.InitGrpc(server, cfg, db, grpcConn)
	// commentModule.InitGrpc(server, cfg, db, grpcConn)
}

func splash(cfg *config.Config) {
	colorReset := "\033[0m"
	// colorBlue := "\033[34m"
	colorCyan := "\033[36m"

	// fmt.Println(colorBlue, fmt.Sprintf(`⇨ REST server started on port :%s`, cfg.Port.REST))
	fmt.Println(colorCyan, fmt.Sprintf(`⇨ GRPC xyz-consumer service server started on port :%s`, cfg.Port.GRPC))
	fmt.Println(colorReset, "")
}