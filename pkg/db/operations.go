package db

import (
	"context"
	"log"
	"time"

	"github.com/iliasgal/network-monitor/pkg/model"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
)

func WritePingMetricsToDB(stats *model.PingStats) {

	ctx := context.Background()

	// Initiate an append opperation
	app := tsdbClient.Appender(ctx)

	p := TSDBPoint{}
	p.Measurement = "ping_metrics"
	p.Timestamp = time.Now()
	p.Fields = map[string]float64{"avg_latency_ms": stats.AvgLatency, "jitter_ms": stats.Jitter, "packet_loss": stats.PacketLoss}

	err := WritePoint(app, p)
	if err != nil {
		app.Rollback()
		log.Fatal(err)
	}
	app.Commit()
}

type TSDBPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]float64
	Timestamp   time.Time
}

func WritePacketInfoToDB(info *model.PacketInfo) {
	writer.Write(TSDBPoint{
		Measurement: "network_traffic",
		Tags: map[string]string{
			"packet_type": info.PacketType,
			"src_ip":      info.SrcIP,
			"dst_ip":      info.DstIP,
			"src_port":    info.SrcPort,
			"dst_port":    info.DstPort,
		},
		Fields: map[string]float64{
			"packet_size": float64(info.Size),
		},
		Timestamp: time.Now(),
	})
}

func WritePoint(app storage.Appender, p TSDBPoint) error {
	ts := p.Timestamp.UnixMilli()

	for field, val := range p.Fields {
		// encode as "measurement_field" or just "measurement" if single field
		name := p.Measurement + "_" + field

		pairs := []string{"__name__", name}
		for k, v := range p.Tags {
			pairs = append(pairs, k, v)
		}

		lbls := labels.FromStrings(pairs...)
		if _, err := app.Append(0, lbls, ts, val); err != nil {
			return err
		}
	}

	return nil
}
