package main

import (
	"fmt"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pktque"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

func init() {
	format.RegisterAll()
}

// as same as: ffmpeg -re -i projectindex.flv -c copy -f flv rtmp://localhost:1936/app/publish

func main() {
	file, err := avutil.Open("projectindex.flv")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := rtmp.Dial("rtmp://localhost:1935/app/publish")
	// conn, _ := avutil.Create("rtmp://localhost:1936/app/publish")
	if err != nil {
		fmt.Println(err)
	}
	demuxer := &pktque.FilterDemuxer{Demuxer: file, Filter: &pktque.Walltime{}}
	avutil.CopyFile(conn, demuxer)

	file.Close()
	conn.Close()
}
