package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	deviceName  = flag.String("device", "eth0", "capture device")
	file        = flag.String("file", "out.pcap", "output file")
	filter      = flag.String("filter", "not (tcp and port 22)", "capture filter")
	count       = flag.Int("count", 1000, "max number of packets to capture")
	snapshotLen int32  = 9000
	promiscuous bool   = false
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
	packetCount int = 0
)

func main() {
	// Open output pcap file and write header 
	f, _ := os.Create(*file)
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err = pcap.OpenLive(*deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		fmt.Printf("Error opening device %s: %v", deviceName, err)
		os.Exit(1)
	}
	defer handle.Close()

	if *filter != "" {
		if err := handle.SetBPFFilter(*filter); err != nil {
			fmt.Printf("Error setting filter: %v", err)
			os.Exit(1)
		}
	}

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++
		
		if packetCount > *count {
			break
		}
	}
}
