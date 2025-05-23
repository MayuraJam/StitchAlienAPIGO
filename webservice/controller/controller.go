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

// controller สำหรับดึงข้อมูลที่ไม่ต้องใช้ Id

func HandlerCreatures(w http.ResponseWriter, r *http.Request) {
	//parameter ที่รับเข้ามาใน func จะเป็น web server ที่ใช้ในการ handler โดยตัวแรกเป็น w ตัวที่ทำการ response ผลไปยัง client อีกอัน r เป็นตัว request ที่มีการรับค่าจาก client ส่งเข้ามาทำงานภายใน function
	// request มีได้หลายแบบ เช่น http method , endpoint url , body , header
	// สร้าง switch case เพื่อทำการเลือก ตาม http method ที่ส่งเข้ามาจากทาง request
	switch r.Method {
	case http.MethodGet: //กรณี Get
		creatureList, err := service.GetCreatureList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) //สามารถ reponse ในส่วนของ HTTP status code และ message ได้
			return
		}
		j, err := json.Marshal(creatureList) //แปลงให้เป็น json
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(j) //ทำการแสดงผล response
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var creature model.Creature
		err := json.NewDecoder(r.Body).Decode(&creature) //อ่านค่า Json จาก req body และทำการ map กับ modelข้อมูล
		if err != nil {
			log.Print(err)
			return
		}
		CreatureId, err := service.InsertNewCreature(creature) //return ค่าออกเป็น id ที่พึงสร้างใหม่
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)                              //201 แจ้งเตือนสร้างสำเร็จ
		w.Write([]byte(fmt.Sprintf(`{"creature_id:%d"}`, CreatureId))) //ทำการแสดง creature_id จากตัวแปร CreatureId และแปลงออกมาในรูปแบบ JSON ด้วย []byte
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

// controller สำหรับดึงข้อมูลที่ต้องใช้ Id
func HandlerCreature(w http.ResponseWriter, r *http.Request) {
	urlPathment := strings.Split(r.URL.Path, fmt.Sprintf("%s/", path)) //ทำการแยกส่วนของ URL เพื่อดึงเอาตัวเลข id
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

//Marshal = struc data -> JSON format
//Unmarshal = JSON format -> struc data
