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
		Float64("stop.the.world.ms", conv_utils.ToMega(m.PauseTotalNs)).
		Float64("heap.alloc.mb", conv_utils.ToMegaByte(m.HeapAlloc)).
		Float64("cum.heap.alloc.mb", conv_utils.ToMegaByte(m.TotalAlloc)).
		Float64("heap.alloc.count.k", conv_utils.ToKilo(m.HeapObjects)).
		Float64("stack.in.use.mb", conv_utils.ToMegaByte(m.StackInuse)).
		Float64("total.sys.mb", conv_utils.ToMegaByte(m.Sys)).
		//nolint: mnd // its percent and precision
		Float64("gc.cpu.percent", conv_utils.ToFixed(m.GCCPUFraction*100, 4)).
		Uint32("gc.cycles", m.NumGC).
		Float64("total.mem.mb", conv_utils.ToMegaByte(totalBytes)).
		Int("goroutine.count", runtime.NumGoroutine()).
		Msg("perfstats")
}

func SpamLogMemUsage(ctx context.Context, rate time.Duration) {
	for {
		LogMemUsage(ctx)
		time.Sleep(rate)
	}
}
