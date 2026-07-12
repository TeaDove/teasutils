package perfutils

import (
	"context"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/teadove/teasutils/utils/convutils"

	"github.com/rs/zerolog"
)

func LogMemUsage(ctx context.Context) {
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/total:bytes"

	metrics.Read(samples)
	totalBytes := samples[0].Value.Uint64()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	zerolog.Ctx(ctx).
		Info().
		Str("stop.the.world", time.Duration(m.PauseTotalNs).String()).
		Str("heap.alloc", convutils.Closest(m.HeapAlloc)).
		Str("cum.heap.alloc", convutils.Closest(m.TotalAlloc)).
		Str("heap.alloc.count", convutils.Closest(m.HeapObjects)).
		Str("stack.in.use", convutils.ClosestByte(m.StackInuse)).
		Str("total.sys.mb", convutils.ClosestByte(m.Sys)).
		//nolint: mnd // its percent and precision
		Float64("gc.cpu.percent", convutils.ToFixed(m.GCCPUFraction*100, 4)).
		Uint32("gc.cycles", m.NumGC).
		Str("total.mem", convutils.ClosestByte(totalBytes)).
		Int("goroutine.count", runtime.NumGoroutine()).
		Msg("perfstats")
}

func SpamLogMemUsage(ctx context.Context, d time.Duration) {
	t := time.NewTicker(d)

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			LogMemUsage(ctx)
		}
	}
}
