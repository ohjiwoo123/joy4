package main

import (
	"fmt"

	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
)

func init() {
	format.RegisterAll()
}

func main() {
	fmt.Println("start")
	file, err := avutil.Open("projectindex.flv")
	if err != nil {
		fmt.Println("Open err")
	}

	streams, err2 := file.Streams()
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(streams)
	for _, stream := range streams {
		if stream.Type().IsAudio() {
			astream := stream.(av.AudioCodecData)
			fmt.Println(astream.Type(), astream.SampleRate(), astream.SampleFormat(), astream.ChannelLayout())
		} else if stream.Type().IsVideo() {
			vstream := stream.(av.VideoCodecData)
			fmt.Println(vstream.Type(), vstream.Width(), vstream.Height())
		} else {
			fmt.Println("NOT AUDIO VIDEO")
		}
	}

	for i := 0; i < 10; i++ {
		var pkt av.Packet
		var err error
		if pkt, err = file.ReadPacket(); err != nil {
			break
		}
		fmt.Println("pkt", i, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
	}

	file.Close()
}
