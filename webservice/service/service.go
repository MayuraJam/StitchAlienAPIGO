package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MayuraJam/StitchAlienAPIGO/webservice/database"
	"github.com/MayuraJam/StitchAlienAPIGO/webservice/model"
)

func GetCreatureList() ([]model.Creature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.Db.QueryContext(ctx, `SELECT *FROM creature_tb`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer result.Close()
	creatures := make([]model.Creature, 0) //เตรียมตัวแปรที่มาจาก model โดยเป็น ตวแปรประเภท list
	//การวนลูปข้อมูลทั้งหมด
	for result.Next() {
		var creature model.Creature
		//ทำการ map ข้อมูลจากฐานข้อมูลมาเก็บไว้ในตัวแปร
		result.Scan(
			&creature.CreatureID,
			&creature.CreatureName,
			&creature.NickName,
			&creature.Species,
			&creature.ImageUrl,
			&creature.Abilities,
		)
		creatures = append(creatures, creature)
	}
	return creatures, nil
}

func GetCreatureItem(id int) (*model.Creature, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.Db.QueryRowContext(ctx, `SELECT * FROM creature_tb WHERE creature_id = ?`, id)
	creature := &model.Creature{}
	err := row.Scan(
		&creature.CreatureID,
		&creature.CreatureName,
		&creature.NickName,
		&creature.Species,
		&creature.ImageUrl,
		&creature.Abilities,
	)
	if err == sql.ErrNoRows {
		// เมื่อไม่มีข้อมูลใน row ที่เราต้องการหา
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(creature, nil)
	return creature, nil
}

func InsertNewCreature(creatureData model.Creature) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.Db.ExecContext(ctx, `
	 INSERT INTO creature_tb
	  (creature_name, 
	  nickname, 
	  species, 
	  imageUrl, 
	  abilities) 
	  VALUES (?, ?, ?, ?, ?)
	`, creatureData.CreatureName,
		creatureData.NickName,
		creatureData.Species,
		creatureData.ImageUrl,
		creatureData.Abilities,
	)
	if err != nil {
		log.Panicln(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Panicln(err.Error())
		return 0, err
	}
	return int(insertID), nil
}
