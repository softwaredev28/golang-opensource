package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"unsia/pb/cities"
)

type City struct{
	Pb cities.City
	Log *log.Logger
}

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `select id,name from cities where id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		u.Log.Println("error any query", err)
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
	
	u.Pb.Name = in.Name
	
	
	return nil
}

func (u *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `delete from cities where id = $1;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	//  _,err = stmt.ExecContext(ctx, in.Id)
	 rs,err := stmt.ExecContext(ctx, in.Id)

	if err != nil {
		return err
	}

	affected,err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if affected== 0 {
		return fmt.Errorf("data not found!")
	}

	return nil
}

func (u *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
	query := `UPDATE cities SET name = $1 WHERE id = $2`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, in.Name, in.Id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("data not found!")
	}

	u.Pb.Id = in.Id
	u.Pb.Name = in.Name

	return nil
}