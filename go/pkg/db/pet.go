package db

import (
	"context"
	"fmt"
	"time"
)

func (service *Db) GetListPets(ctx context.Context, query string) ([]*Pet, error) {
	var pets []*Pet
	//id,user_id,name,type,age,photo,created_at
	rows, err := service.handler.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to create query %s, error: %v", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var userId string
		var typep string
		var createdAt time.Time
		var photo string
		var age string
		//request is SELECT id,user_id,name,type,age,photo,created_at FROM pets WHERE 1=1
		err = rows.Scan(&id, &userId, &name, &typep, &age, &photo, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("failed to get pet from pets table, error:%#v", err)
		}
		pets = append(pets, &Pet{ID: id, Name: name, Age: age, Type: typep, UserId: userId, CreatedAt: createdAt, Photo: photo})

	}
	if err != nil {
		return nil, fmt.Errorf("failed to Get Pet by id - error:%#v", err)
	}
	return pets, nil
}
