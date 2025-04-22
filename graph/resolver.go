package graph

import (
	"github.com/himanshuholmes/bikerental/gen/go/proto/bikes"
	"github.com/himanshuholmes/bikerental/gen/go/proto/rentees"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	BikeClient bikes.BikesAPIClient
	RenteeClient rentees.RenteesAPIClient
}

func NewResolver()*Resolver{
	bikeConn,err := grpc.NewClient("localhost:6000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	renteeConn,err := grpc.NewClient("localhost:6001",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Resolver{BikeClient: bikes.NewBikesAPIClient(bikeConn),RenteeClient: rentees.NewRenteesAPIClient(renteeConn)}
}