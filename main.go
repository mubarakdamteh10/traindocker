package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mubarakdamteh10/traindocker/customer_order"
	"github.com/mubarakdamteh10/traindocker/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// create method
	clientOptions := options.Client().ApplyURI("mongodb+srv://mubarakdamteh10:Mubarak12345@cluster0.yainj.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("coin")

	fmt.Println("Connected to MongoDB!")

	e := echo.New()
	e.GET("/users/:Name", users.GetUserByIdHandler(users.GetUserById(db)))
	e.GET("/getuser", users.GetAllUserHandler(users.GetAllUser(db)))
	e.POST("/user", users.CreateUserHandler(users.CreateUser(db)))
	e.PATCH("edit/:ID", users.UpdateUserByIdParamHandler(users.UpdateUserByIdParam(db)))
	e.PUT("editfield/:ID", users.UpdateUserByIdFieldHandler(users.UpdateUserByIdField(db)))
	e.DELETE("/delete/:ID", users.DeleteUserByIdHandler(users.DeleteUserById(db)))

	// order method
	e.POST("/createorder", customer_order.CreateOrderHandler(customer_order.CreateCustomerOrder(db)))
	e.GET("/getorder", customer_order.GetAllOrderHandler(customer_order.GetAllOrder(db)))
	e.Logger.Fatal(e.Start(":8000"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
