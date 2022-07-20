package net

import (
	"net"
	"sync"
	"testing"

	pk "github.com/Tnze/go-mc/net/packet"
)

func BenchmarkConnTCPCompress(b *testing.B) {
	cs, ss, err := getTCPConnectionPair()
	if err != nil {
		b.Fatal(err)
	}
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompress(b, 0, cs, ss)
}

func BenchmarkConnTCPCompressLevelDefault(b *testing.B) {
	cs, ss, err := getTCPConnectionPair()
	if err != nil {
		b.Fatal(err)
	}
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompressLevel(b, -1, cs, ss)
}

func BenchmarkConnTCPCompressLevelBestSpeed(b *testing.B) {
	cs, ss, err := getTCPConnectionPair()
	if err != nil {
		b.Fatal(err)
	}
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompressLevel(b, 1, cs, ss)
}

func BenchmarkConnTCPNoCompress(b *testing.B) {
	cs, ss, err := getTCPConnectionPair()
	if err != nil {
		b.Fatal(err)
	}
	defer cs.Close()
	defer ss.Close()

	// no compress
	benchConnCompress(b, -1, cs, ss)
}

func BenchmarkConnPipeCompress(b *testing.B) {
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompress(b, 0, cs, ss)
}

func BenchmarkConnPipeCompressLevelDefault(b *testing.B) {
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompressLevel(b, -1, cs, ss)
}

func BenchmarkConnPipeCompressLevelBestSpeed(b *testing.B) {
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	// force compress
	benchConnCompressLevel(b, 1, cs, ss)
}

func BenchmarkConnPipeNoCompress(b *testing.B) {
	cs, ss := net.Pipe()
	defer cs.Close()
	defer ss.Close()

	// no compress
	benchConnCompress(b, -1, cs, ss)
}

func benchConnCompress(b *testing.B, threshold int, cs net.Conn, ss net.Conn) {
	rd := WrapConn(cs)
	rd.SetThreshold(threshold)

	wr := WrapConn(ss)
	wr.SetThreshold(threshold)

	bench(b, rd, wr)
}

func benchConnCompressLevel(b *testing.B, level int, cs net.Conn, ss net.Conn) {
	rd := WrapConn(cs)
	rd.SetThresholdLevel(0, level)

	wr := WrapConn(ss)
	wr.SetThresholdLevel(0, level)

	bench(b, rd, wr)
}

func getTCPConnectionPair() (net.Conn, net.Conn, error) {
	lst, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, nil, err
	}
	defer lst.Close()

	var conn0 net.Conn
	var err0 error
	done := make(chan struct{})
	go func() {
		conn0, err0 = lst.Accept()
		close(done)
	}()

	conn1, err := net.Dial("tcp", lst.Addr().String())
	if err != nil {
		return nil, nil, err
	}

	<-done
	if err0 != nil {
		return nil, nil, err0
	}
	return conn0, conn1, nil
}

func bench(b *testing.B, rd *Conn, wr *Conn) {
	buf := make([]byte, 128*1024)
	buf2 := make([]byte, 128*1024)
	pkt := pk.Packet{
		ID:   0x01,
		Data: buf,
	}
	pkt2 := pk.Packet{
		ID:   0x01,
		Data: buf2,
	}

	b.SetBytes(128 * 1024)
	b.ResetTimer()
	b.ReportAllocs()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := 0
		for {
			rd.ReadPacket(&pkt2)
			count += len(pkt2.Data)
			if count == 128*1024*b.N {
				return
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		wr.WritePacket(pkt)
	}
	wg.Wait()
}
