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
	u.Pb.Id = 1
	u.Pb.Name = "Malang"
	query := `select id,name from cities where id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		return err
	}
	return nil
}