package main

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	handle *pcap.Handle
}

type TCPParams struct {
	SRC          string
	DST          string
	CountOptions int
}

type IPParams struct {
	SRC string
	DST string
}

func NewSniffer(device string) (*Sniffer, error) {
	handle, err := pcap.OpenLive(device, 1600, true, time.Millisecond)

	if err != nil {
		return nil, err
	}

	err = handle.SetBPFFilter("tcp and ip and dst port 443")

	if err != nil {
		return nil, err
	}

	sn := &Sniffer{
		handle: handle,
	}

	return sn, nil
}

func (s *Sniffer) Run() error {
	go func() {
		packetSource := gopacket.NewPacketSource(s.handle,
			s.handle.LinkType())

		for packet := range packetSource.Packets() {
			s.handlePacket(packet)
		}
	}()

	return nil
}

func (s *Sniffer) handlePacket(packet gopacket.Packet) {
	tcp := s.resolveTCPParams(packet)
	ip := s.resolveIPParams(packet)

	result := fmt.Sprintf("%s,%s,%s,%s,%d",
		ip.SRC,
		tcp.SRC,
		ip.DST,
		tcp.DST,
		tcp.CountOptions,
	)

	fmt.Println(result)
}

func (s *Sniffer) resolveTCPParams(packet gopacket.Packet) *TCPParams {
	tcp := packet.TransportLayer().(*layers.TCP)

	tcpParams := &TCPParams{
		SRC:          tcp.SrcPort.String(),
		DST:          tcp.DstPort.String(),
		CountOptions: len(tcp.Options),
	}

	return tcpParams
}

func (s *Sniffer) resolveIPParams(packet gopacket.Packet) *IPParams {
	ip := packet.NetworkLayer().(*layers.IPv4)

	ipParams := &IPParams{
		SRC: ip.SrcIP.String(),
		DST: ip.DstIP.String(),
	}

	return ipParams
}

func (s *Sniffer) Stop() {
	s.handle.Close()
}
