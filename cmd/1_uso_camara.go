package main

import (
	"log"

	"gocv.io/x/gocv"
)

const DeviceID = 0

func main() {
	var camera *gocv.VideoCapture
	var key int

	camera, err := gocv.VideoCaptureDevice(DeviceID)
	if err != nil {
		log.Fatalf("*** Error abriendo cámara (device ID.:%d) ***\n", DeviceID)
	}
	defer camera.Close()

	frame := gocv.NewMat()
	defer frame.Close()

	window := gocv.NewWindow("Webcam")
	defer window.Close()

	for {
		if !camera.Read(&frame) {
			log.Println("*** No se pudo leer la cámara ***")
			break
		}
		if frame.Empty() {
			continue
		}

		window.IMShow(frame)
		key = window.WaitKey(1)
		if key == 27 { //Esc
			break
		}
	}
}
