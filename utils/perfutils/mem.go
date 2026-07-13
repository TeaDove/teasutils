package perfutils

import (
	"context"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/teadove/teasutils/utils/convutils"

	"github.com/rs/zerolog"
)

// LogMemUsage logs a one-off snapshot of runtime memory and GC statistics
// at info level using the ctx logger.
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

// SpamLogMemUsage calls LogMemUsage every d until ctx is cancelled.
// It blocks, so run it in its own goroutine.
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
