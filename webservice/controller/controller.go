package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/MayuraJam/StitchAlienAPIGO/webservice/model"
	"github.com/MayuraJam/StitchAlienAPIGO/webservice/service"
)

const path = "stitch"

// obj หลายตัว
func HandlerCreatures(w http.ResponseWriter, r *http.Request) {
	// สร้าง switch case เพื่อทำการเลือก ตาม http method ที่ส่งเข้ามาจากทาง request
	switch r.Method {
	case http.MethodGet: //กรณี Get
		creatureList, err := service.GetCreatureList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //ผิดผลาดที่เกิดจาก server
			return
		}
		j, err := json.Marshal(creatureList) //แปลงให้เป็น json
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //ผิดผลาดที่เกิดจาก server
			return
		}
		_, err = w.Write(j) //ทำการแสดงผล request ที่อยู่ในรูปแบบ Json
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //ผิดผลาดที่เกิดจาก server
			return
		}
	case http.MethodPost:
		var creature model.Creature
		err := json.NewDecoder(r.Body).Decode(&creature) //ข้อมูลจากการ request
		if err != nil {
			log.Print(err)
			return
		}
		CreatureId, err := service.InsertNewCreature(creature)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"creature_id:%d"}`, CreatureId)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func HandlerCreature(w http.ResponseWriter, r *http.Request) {
	urlPathment := strings.Split(r.URL.Path, fmt.Sprintf("%s/", path))
	if len(urlPathment[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	creatureId, err := strconv.Atoi(urlPathment[len(urlPathment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		creature, err := service.GetCreatureItem(creatureId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if creature == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(creature)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}
