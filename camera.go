package main

import (
	"fmt"
	"github.com/hulucat/conf"
	"github.com/hulucat/gofaceplus"
	"github.com/hulucat/utils"
	"os"
	"time"
)

type Person struct {
	Age    int64  `json:"age"`
	Gender string `json:"gender"`
	Glass  string `json:"glass"`
}

var faceClient *gofaceplus.FaceClient

func StartCamera() {
	faceClient = &gofaceplus.FaceClient{
		conf.Get("face_server"),
		conf.Get("face_api_key"),
		conf.Get("face_secret"),
	}
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			detect("./1.jpg")
		}
	}
}

func detect(filePath string) {
	sessionId, faces, _, err := faceClient.DetectImg(filePath)
	if err != nil {
		fmt.Printf("Error detect img: %s \n", err.Error())
		os.Exit(0)
	}
	utils.Debugf("Detect result: sessionId=%s, faces: %d", sessionId, len(faces))
	if len(faces) < 1 {
		return
	}
	people := make([]*Person, len(faces))
	for i, face := range faces {
		people[i] = &Person{
			Age:    face.Attrs.Age.Value,
			Gender: face.Attrs.Gender.Value,
			Glass:  face.Attrs.Glass.Value,
		}
	}
	if WebSocketConn != nil {
		if err = WebSocketConn.WriteJSON(people); err != nil {
			utils.Errorf("Error write json to websocket: %s", err.Error())
		}
	} else {
		utils.Debugf("Websocket is nil")
	}
}
