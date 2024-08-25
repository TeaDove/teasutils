package perf_utils

import (
	"context"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/teadove/teasutils/utils/converters_utils"

	"github.com/rs/zerolog"
)

func LogMemUsage(ctx context.Context) {
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/total:bytes"

	metrics.Read(samples)
	totalBytes := samples[0].Value.Uint64()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	zerolog.Ctx(ctx).Info().
		Float64("stop.the.world.ms", converters_utils.ToMega(m.PauseTotalNs)).
		Float64("heap.alloc.mb", converters_utils.ToMegaByte(m.HeapAlloc)).
		Float64("cum.heap.alloc.mb", converters_utils.ToMegaByte(m.TotalAlloc)).
		Float64("heap.alloc.count.k", converters_utils.ToKilo(m.HeapObjects)).
		Float64("stack.in.use.mb", converters_utils.ToMegaByte(m.StackInuse)).
		Float64("total.sys.mb", converters_utils.ToMegaByte(m.Sys)).
		Float64("gc.cpu.percent", converters_utils.ToFixed(m.GCCPUFraction*100, 4)).
		Uint32("gc.cycles", m.NumGC).
		Float64("total.mem.mb", converters_utils.ToMegaByte(totalBytes)).
		Int("goroutine.count", runtime.NumGoroutine()).
		Msg("perfstats")
}

func SpamLogMemUsage(ctx context.Context, rate time.Duration) {
	for {
		LogMemUsage(ctx)
		time.Sleep(rate)
	}
}
