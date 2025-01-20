package perf_utils

import (
	"context"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/teadove/teasutils/utils/conv_utils"

	"github.com/rs/zerolog"
)

func LogMemUsage(ctx context.Context) {
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/total:bytes"

	metrics.Read(samples)
	totalBytes := samples[0].Value.Uint64()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	zerolog.Ctx(ctx).Error().
		Str("stop.the.world", time.Duration(m.PauseTotalNs).String()).
		Str("heap.alloc", conv_utils.Closest(m.HeapAlloc)).
		Str("cum.heap.alloc", conv_utils.Closest(m.TotalAlloc)).
		Str("heap.alloc.count", conv_utils.Closest(m.HeapObjects)).
		Str("stack.in.use", conv_utils.ClosestByte(m.StackInuse)).
		Str("total.sys.mb", conv_utils.ClosestByte(m.Sys)).
		//nolint: mnd // its percent and precision
		Float64("gc.cpu.percent", conv_utils.ToFixed(m.GCCPUFraction*100, 4)).
		Uint32("gc.cycles", m.NumGC).
		Str("total.mem", conv_utils.ClosestByte(totalBytes)).
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
