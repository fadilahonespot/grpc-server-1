package main

import (
	"fmt"
	"grpc-server-1/model"
	userHandler "grpc-server-1/user/handler"
	userRepo "grpc-server-1/user/repo"
	userUsecase "grpc-server-1/user/usecase"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func main() {
	var port = "8957"
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_grpc?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	
	db.Debug().AutoMigrate(
		model.UserDB{},
	)

	server := grpc.NewServer()

	userRepo := userRepo.CreateUserRepoImpl(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)

	userHandler.CreateUserHandler(server, userUsecase)

	conn, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Starting at Port: ", port)
	log.Fatal(server.Serve(conn))
}
