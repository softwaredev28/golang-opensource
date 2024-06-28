package models

import (
	"context"
	"database/sql"
	"unsia/pb/cities"
)

type City struct{
	Pb cities.City
}

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	// u.Pb.Id = 1
	// u.Pb.Name = "Malang"
	query := `select id,name from cities where id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		return err
	}
	return nil
}

func (u *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `insert into cities (name) values($1) returning id`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	 err = stmt.QueryRowContext(ctx, in.Name).Scan(&u.Pb.Id)

	if err != nil {
		return err
	}
	
	// lastId,err := rs.LastInsertId()
	// if err != nil {
	// 	return err
	// }

	u.Pb.Name = in.Name
	// u.Pb.Id = int32(lastId)
	
	
	return nil
}