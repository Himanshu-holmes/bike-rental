package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	renteesV1 "github.com/himanshuholmes/bikerental/gen/go/proto/rentees"
	"github.com/himanshuholmes/bikerental/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

const (
	renteeTableName = "rentees"
	defaultPort     = "6001"
)

type Server struct {
	renteesV1.UnimplementedRenteesAPIServer
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
	listener,err := net.Listen("tcp",fmt.Sprintf(":%s",port))
	if err != nil {
		log.Printf("net.Listen error: %v", err)
		return
	}
	grpcServer := grpc.NewServer();
	renteesV1.RegisterRenteesAPIServer(grpcServer,s);
	reflection.Register(grpcServer);

	log.Printf("Starting Rental Rentees server on port %s",port)
	go func(){
		grpcServer.Serve(listener);
	}()

}

func (s *Server)ListRentees(ctx context.Context, req *renteesV1.ListRenteesRequest)(*renteesV1.ListRenteesResponse, error){
	if req == nil {
		return nil, fmt.Errorf("request is empty")
	}
	var rentees [] models.RenteeModel
	if err := s.db.WithContext(ctx).Find(&rentees).Error; err != nil {
		return nil, err
	}
	
	protoRentees := make([]*renteesV1.Rentee, len(rentees))
	for i, rentee := range rentees {
		protoHeldBikes := make([]string,len(rentee.HeldBikes))
		for i, bike := range rentee.HeldBikes {
			protoHeldBikes[i] = bike.ID
		}

		protoRentees[i] = &renteesV1.Rentee{
			Id:   rentee.ID,
			FirstName: rentee.FirstName,
			LastName: rentee.LastName,
			National_Id_Number: rentee.NationalIdNumber,
			Phone: rentee.Phone,
			Email: rentee.Email,
			HeldBikes: protoHeldBikes,
		}
	}
	return &renteesV1.ListRenteesResponse{Rentees: protoRentees}, nil

}

func (s *Server)GetRentee(ctx context.Context, req *renteesV1.GetRenteeRequest)(*renteesV1.GetRenteeResponse, error){
	if req == nil || req.Id == "" {
		return nil, fmt.Errorf("request is empty")
	}
	rentee := new(models.RenteeModel)
	
	if err := s.db.WithContext(ctx).Where("id = ?",req.Id).First(&rentee).Error; err != nil {
		return nil, err
	}
	protoHeldBikes := make([]string, len(rentee.HeldBikes))
	for i, bike := range rentee.HeldBikes {
		protoHeldBikes[i] = bike.ID
	}
	protoRentee := &renteesV1.Rentee{
		Id:   rentee.ID,
		FirstName:  rentee.FirstName,
		LastName: rentee.LastName,
		National_Id_Number: rentee.NationalIdNumber,
		Phone: rentee.Phone,
		Email: rentee.Email,
		HeldBikes: protoHeldBikes,
	}
	return &renteesV1.GetRenteeResponse{Rentee: protoRentee}, nil
}

func (s *Server)GetRenteeByBikeId(ctx context.Context, req *renteesV1.GetRenteeByBikeIdRequest)(*renteesV1.GetRenteeByBikeIdResponse, error){
	if req == nil || req.Id == "" {
      return nil, fmt.Errorf("bike id is not provided")
	}
	bike := new(models.BikeModel)
	if err := s.db.WithContext(ctx).Where("id = ?",req.Id).First(&bike).Error; err!=nil{
		return nil, err
	}
	rentee := new(models.RenteeModel)
	if err := s.db.WithContext(ctx).Where("id = ?",bike.ID).First(&rentee).Error; err!=nil{
		return nil, err
	}

	protoHeldBikes := make([]string, len(rentee.HeldBikes))
	for i, bike := range rentee.HeldBikes {
		protoHeldBikes[i] = bike.ID
	}
	protoRentee := &renteesV1.Rentee{
		Id:   rentee.ID,
		FirstName:  rentee.FirstName,
		LastName: rentee.LastName,
		National_Id_Number: rentee.NationalIdNumber,
		Phone: rentee.Phone,
		Email: rentee.Email,
		HeldBikes: protoHeldBikes,
	}
	return &renteesV1.GetRenteeByBikeIdResponse{Rentee: protoRentee}, nil

}

func (s *Server) GetRenteesByBikeTYPE(ctx context.Context, req *renteesV1.GetRenteesByBikeTYPERequest)(*renteesV1.GetRenteeByBikeTYPEResponse, error){
	if req == nil || req.Type == "" {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("type = ?", req.Type).Find(&bikes).Error; err != nil {
		return nil, err
	}
	var rentees [] models.RenteeModel
	if err := s.db.WithContext(ctx).Where("id in (?)", bikes).Find(&rentees).Error; err != nil {
		return nil, err
	}
	
	protoRentees := make([]*renteesV1.Rentee, len(rentees))
	for i, rentee := range rentees {
		protoHeldBikes := make([]string,len(rentee.HeldBikes))
		for i, bike := range rentee.HeldBikes {
			protoHeldBikes[i] = bike.ID
		}

		protoRentees[i] = &renteesV1.Rentee{
			Id:   rentee.ID,
			FirstName: rentee.FirstName,
			LastName: rentee.LastName,
			National_Id_Number: rentee.NationalIdNumber,
			Phone: rentee.Phone,
			Email: rentee.Email,
			HeldBikes: protoHeldBikes,
		}
	}
	return &renteesV1.GetRenteeByBikeTYPEResponse{Rentees: protoRentees}, nil
}

func (s *Server)GetRenteesByBikeMAKE(ctx context.Context, req *renteesV1.GetRenteeByBikeMAKERequest)(*renteesV1.GetRenteeByBikeMAKEResponse, error){
	if req == nil || req.Make == "" {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("make = ?", req.Make).Find(&bikes).Error; err != nil {
		return nil, err
	}
	var rentees [] models.RenteeModel
	if err := s.db.WithContext(ctx).Where("id in (?)", bikes).Find(&rentees).Error; err != nil {
		return nil, err
	}
	
	protoRentees := make([]*renteesV1.Rentee, len(rentees))
	for i, rentee := range rentees {
		protoHeldBikes := make([]string,len(rentee.HeldBikes))
		for i, bike := range rentee.HeldBikes {
			protoHeldBikes[i] = bike.ID
		}

		protoRentees[i] = &renteesV1.Rentee{
			Id:   rentee.ID,
			FirstName: rentee.FirstName,
			LastName: rentee.LastName,
			National_Id_Number: rentee.NationalIdNumber,
			Phone: rentee.Phone,
			Email: rentee.Email,
			HeldBikes: protoHeldBikes,
		}
	}
	return &renteesV1.GetRenteeByBikeMAKEResponse{Rentees: protoRentees}, nil
}

func (s *Server) GetRenteesByBikeOWNER(ctx context.Context, req *renteesV1.GetRenteeByBikeOWNERRequest)(*renteesV1.GetRenteeByBikeOWNERResponse, error){
	if req == nil || req.OwnerName == "" {
		return nil, fmt.Errorf("request is empty")
	}
	var bikes [] models.BikeModel
	if err := s.db.WithContext(ctx).Where("owner_name = ?", req.OwnerName).Find(&bikes).Error; err != nil {
		return nil, err
	}
	var rentees [] models.RenteeModel
	if err := s.db.WithContext(ctx).Where("id in (?)", bikes).Find(&rentees).Error; err != nil {
		return nil, err
	}
	
	protoRentees := make([]*renteesV1.Rentee, len(rentees))
	for i, rentee := range rentees {
		protoHeldBikes := make([]string,len(rentee.HeldBikes))
		for i, bike := range rentee.HeldBikes {
			protoHeldBikes[i] = bike.ID
		}

		protoRentees[i] = &renteesV1.Rentee{
			Id:   rentee.ID,
			FirstName: rentee.FirstName,
			LastName: rentee.LastName,
			National_Id_Number: rentee.NationalIdNumber,
			Phone: rentee.Phone,
			Email: rentee.Email,
			HeldBikes: protoHeldBikes,
		}
	}
	return &renteesV1.GetRenteeByBikeOWNERResponse{Rentees: protoRentees}, nil
}

func (s *Server)AddRentee(ctx context.Context, req *renteesV1.AddRenteeRequest)(*renteesV1.AddRenteeResponse, error){
	if req == nil || req.Rentee == nil {
		return nil, fmt.Errorf("request is empty")
	}
	rentee := &models.RenteeModel{
		ID: req.Rentee.Id,
		FirstName: req.Rentee.FirstName,
		LastName: req.Rentee.LastName,
		NationalIdNumber: req.Rentee.National_Id_Number,
		Phone: req.Rentee.Phone,
		Email: req.Rentee.Email,
	}
	if err := s.db.WithContext(ctx).Create(rentee).Error; err != nil {
		return nil, err
	}
	protoRentee := &renteesV1.Rentee{
		Id:                rentee.ID,
		FirstName:         rentee.FirstName,
		LastName:          rentee.LastName,
		National_Id_Number: rentee.NationalIdNumber,
		Phone:             rentee.Phone,
		Email:             rentee.Email,
	}
	return &renteesV1.AddRenteeResponse{Rentee: protoRentee}, nil
}
func (s *Server)UpdateRentee(ctx context.Context, req *renteesV1.UpdateRenteeRequest)(*renteesV1.UpdateRenteeResponse, error){
	if req == nil || req.Rentee == nil {
		return nil, fmt.Errorf("request is empty")
	}
	rentee := &models.RenteeModel{
		ID: req.Rentee.Id,
		FirstName: req.Rentee.FirstName,
		LastName: req.Rentee.LastName,
		NationalIdNumber: req.Rentee.National_Id_Number,
		Phone: req.Rentee.Phone,
		Email: req.Rentee.Email,
	}
	if err := s.db.WithContext(ctx).Save(rentee).Error; err != nil {
		return nil, err
	}
	protoRentee := &renteesV1.Rentee{
		Id:                rentee.ID,
		FirstName:         rentee.FirstName,
		LastName:          rentee.LastName,
		National_Id_Number: rentee.NationalIdNumber,
		Phone:             rentee.Phone,
		Email:             rentee.Email,
	}
	return &renteesV1.UpdateRenteeResponse{Rentee: protoRentee}, nil
}

