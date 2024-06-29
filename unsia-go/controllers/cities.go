package controllers

import (
	"context"
	"database/sql"
	"log"
	"unsia/models"
	"unsia/pb/cities"
)

// City struct
type City struct{
	DB *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel  models.City
	cityModel.Log = s.Log
	err := cityModel.Get(ctx,s.DB, in)
	return &cityModel.Pb, err
}

// here s

func (s *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, s.DB, in)
	return &cityModel.Pb, err
}

func (s *City) Delete(ctx context.Context, in *cities.Id) (*cities.MyBoolean, error) {
	var cityModel models.City
	err := cityModel.Delete(ctx, s.DB, in)
	if err != nil {
		return &cities.MyBoolean{Boolean:false}, err
	}
	return &cities.MyBoolean{Boolean:true}, err
}

func (s *City) Update(ctx context.Context, in *cities.City) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = s.Log
	err := cityModel.Update(ctx, s.DB, in)
	return &cityModel.Pb, err
}