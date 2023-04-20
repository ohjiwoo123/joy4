package main

import (
	"fmt"
	"reflect"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

func init() {
	format.RegisterAll()
}

// as same as: ffmpeg -re -i projectindex.flv -c copy -f flv rtmp://localhost:1936/app/publish

func main() {
	file, err := avutil.Open("sample-3.flv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(file))
	conn, err := rtmp.Dial("rtmp://localhost:1935/app/publish")
	// conn, _ := avutil.Create("rtmp://localhost:1936/app/publish")
	if err != nil {
		fmt.Println(err)
	}

	// for {
	// 	demuxer := &pktque.FilterDemuxer{Demuxer: file, Filter: &pktque.Walltime{}}
	// 	avutil.CopyFile(conn, demuxer)

	// 	// Reset file to the beginning for subsequent loops
	// 	file.Seek(0, 0)
	// }

	file.Close()
	conn.Close()
}
