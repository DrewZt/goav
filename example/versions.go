package main

import (
	"log"

	"github.com/DrewZt/goav/avcodec"
	"github.com/DrewZt/goav/avdevice"
	"github.com/DrewZt/goav/avfilter"
	"github.com/DrewZt/goav/avformat"
	"github.com/DrewZt/goav/avutil"
	"github.com/DrewZt/goav/swresample"
	"github.com/DrewZt/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
