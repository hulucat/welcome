package main

import (
	"fmt"
	"github.com/hulucat/conf"
	"github.com/hulucat/gofaceplus"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome!")
	faceClient := &gofaceplus.FaceClient{
		conf.Get("face_server"),
		conf.Get("face_api_key"),
		conf.Get("face_secret"),
	}
	sessionId, faces, img, err := faceClient.DetectImg("/Users/wangzheng/Downloads/4.png")
	if err != nil {
		fmt.Printf("Error detect img: %s \n", err.Error())
		os.Exit(0)
	}
	fmt.Printf("SessionId: %s \n", sessionId)
	fmt.Printf("Img: %v \n", img)
	fmt.Printf("Faces: %v \n", faces)
}

//timer
func regTasks() {
	timer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer.C:
			continue
			fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
		}
	}

}
