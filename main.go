package main

import (
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v3"
)

func main() {
	var err error
	outboundVideoTrack, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: "video/h264",
	}, "pion-rtsp", "pion-rtsp")
	if err != nil {
		panic(err)
	}

	go rtspConsumer()

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/doSignaling", doSignaling)

	fmt.Println("Open http://localhost:8080 to access this demo")
	panic(http.ListenAndServe(":8080", nil))
}