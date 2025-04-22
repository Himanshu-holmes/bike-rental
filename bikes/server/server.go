package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	bikesV1 "github.com/himanshuholmes/bikerental/gen/go/proto/bikes"
	"github.com/himanshuholmes/bikerental/models"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	bikeTableName = "bikes"
	defaultPort   = "6000"
)

type Server struct {
	bikesV1.UnimplementedBikesAPIServer
	db *gorm.DB
}

func NewServer(ctx context.Context, db *gorm.DB) (*Server, error) {
	return &Server{
		db: db,
	}, nil
}

func (s *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
		return
	}
	grpcServer := grpc.NewServer()
	bikesV1.RegisterBikesAPIServer(grpcServer, s)
	reflection.Register(grpcServer)
	log.Printf("Starting Rental Bikes server on port %s", port)
	go func() {
		grpcServer.Serve(listener)
	}()
}

func (s *Server) ListBikes(ctx context.Context, req *bikesV1.ListBikesRequest) (*bikesV1.ListBikesResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Find(&bikes).Error; err != nil {
		return nil, err
	}
	var bikesV1Bikes []*bikesV1.Bike
	for _, bike := range bikes {
		bikesV1Bikes = append(bikesV1Bikes, &bikesV1.Bike{
			Id:        bike.ID,
			Type:      bike.Type,
			Make:      bike.Make,
			Serial:    bike.Serial,
			OwnerName: bike.OwnerName,
		})
	}
	return &bikesV1.ListBikesResponse{Bikes: bikesV1Bikes}, nil
}

func (s *Server) GetBike(ctx context.Context, req *bikesV1.GetBikeRequest) (*bikesV1.GetBikeResponse, error) {
	if req == nil || req.Id == "" {
		return nil, fmt.Errorf("request is empty")
	}
	bike := new(models.BikeModel)
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).First(&bike).Error; err != nil {
		return nil, err
	}
	bikeProto := &bikesV1.Bike{
	Id:        bike.ID,
	Type:      bike.Type,
	Make:      bike.Make,
	OwnerName: bike.OwnerName,
	Serial:    bike.Serial,
	}
	
	return &bikesV1.GetBikeResponse{Bike:bikeProto }, nil
}

func (s *Server) GetBikes(ctx context.Context, req *bikesV1.GetBikesRequest) (*bikesV1.GetBikesResponse, error) {
	if req == nil || len(req.Ids) == 0 {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("id in (?)", req.Ids).Find(&bikes).Error; err != nil {
		return nil, err
	}
	bikesProto := make([]*bikesV1.Bike, len(bikes))
	for i, bike:= range bikes {
		bikesProto[i]=&bikesV1.Bike{
			Id:        bike.ID,
			Type:      bike.Type,
			Make:      bike.Make,
			OwnerName: bike.OwnerName,
			Serial:    bike.Serial,
		}
	}
	return &bikesV1.GetBikesResponse{Bikes: bikesProto}, nil
}

func (s *Server) GetBikesByTYPE(ctx context.Context, req *bikesV1.GetBikesByTYPERequest) (*bikesV1.GetBikeByTYPEResponse, error) {
	if req == nil || req.Type == "" {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("type = ?", req.Type).Find(&bikes).Error; err != nil {
		return nil, err
	}
	bikesProto := make([]*bikesV1.Bike, len(bikes))
	for i, bike := range bikes {
		bikesProto[i] = &bikesV1.Bike{
			Id:        bike.ID,
			Type:      bike.Type,
			Make:      bike.Make,
			OwnerName: bike.OwnerName,
			Serial:    bike.Serial,
		}
	}
	return &bikesV1.GetBikeByTYPEResponse{Bikes: bikesProto}, nil
}

func (s *Server)GetBikesByOWNER(ctx context.Context, req *bikesV1.GetBikesByOWNERRequest)(*bikesV1.GetBikesByOWNERResponse, error){
	if req == nil || req.OwnerName == ""{
		return nil, fmt.Errorf("bike owner is not provided")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("owner_name = ?", req.OwnerName).Find(&bikes).Error; err != nil{
		return nil, err
	}
	protoBikes := make([]*bikesV1.Bike, len(bikes))
	for i, bike := range bikes {
		protoBikes[i] = &bikesV1.Bike{
			Id:        bike.ID,
			Type:      bike.Type,
			Make:      bike.Make,
			OwnerName: bike.OwnerName,
			Serial:    bike.Serial,
		}
	}
	return &bikesV1.GetBikesByOWNERResponse{Bikes: protoBikes}, nil
	
}

func(s *Server)GetBikesByMAKE(ctx context.Context, req *bikesV1.GetBikesByMAKERequest)(*bikesV1.GetBikesByMAKEResponse, error){
	if req == nil || req.Make == ""{
		return nil, fmt.Errorf("bike make is not provided")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("make = ?", req.Make).Find(&bikes).Error; err != nil{
		return nil, err
	}
	protoBikes := make([]*bikesV1.Bike, len(bikes))
	for i, bike := range bikes {
		protoBikes[i] = &bikesV1.Bike{
			Id:        bike.ID,
			Type:      bike.Type,
			Make:      bike.Make,
			OwnerName: bike.OwnerName,
			Serial:    bike.Serial,
		}
	}
	return &bikesV1.GetBikesByMAKEResponse{Bikes: protoBikes}, nil
}

func (s *Server) AddBike(ctx context.Context, req *bikesV1.AddBikeRequest) (*bikesV1.AddBikeResponse, error) {
	if req == nil || req.Bike == nil {
		return nil, fmt.Errorf("bike is not provided")
	}
	// print readable bike request
	readBike := req.Bike
	fmt.Printf("Adding bike: %+v\n", readBike)
	bike := new(models.BikeModel)
	bike.ID = readBike.Id
	bike.Type = readBike.Type
	bike.Make = readBike.Make
	bike.OwnerName = readBike.OwnerName
	bike.Serial = readBike.Serial
	if err := s.db.WithContext(ctx).Create(&bike).Error; err != nil {
		return nil, err
	}
	protoBike := &bikesV1.Bike{
		Id:        bike.ID,
		Type:      bike.Type,
		Make:      bike.Make,
		OwnerName: bike.OwnerName,
		Serial:    bike.Serial,
	}
	return &bikesV1.AddBikeResponse{Bike: protoBike}, nil
}
func (s *Server) DeleteBike(ctx context.Context, req *bikesV1.DeleteBikeRequest)(*bikesV1.DeleteBikeResponse, error){
	if req == nil || req.Id == ""{
		return nil, fmt.Errorf("bike id is not provided")
	}
	bike := new(models.BikeModel)
	if err := s.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&bike).Error; err != nil{
		return nil, err
	}
	
	return &bikesV1.DeleteBikeResponse{}, nil
}