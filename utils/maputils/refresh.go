package maputils

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/teadove/teasutils/utils/reflectutils"
	"github.com/teadove/teasutils/utils/timeutils"
)

// Refresher produces the full map contents for one refresh cycle.
type Refresher[K comparable, V any] func() (map[K]V, error)

// RefreshingAtomicMap wraps an AtomicMap that is periodically rebuilt from a
// Refresher in the background. Reads are lock-free via Get.
type RefreshingAtomicMap[K comparable, V any] struct {
	atomicMap     AtomicMap[K, V]
	refresher     Refresher[K, V]
	refresherName string
}

// NewRefreshingAtomicMap builds the map once synchronously (returning an error
// if that first refresh fails) and then spawns a goroutine that refreshes it
// every period. The goroutine runs until ctx is cancelled; failed background
// refreshes are logged and keep the previous contents.
func NewRefreshingAtomicMap[K comparable, V any](
	ctx context.Context,
	period time.Duration,
	refresher Refresher[K, V],
) (*RefreshingAtomicMap[K, V], error) {
	refresherName := reflectutils.GetFunctionName(refresher)

	atomicMap := RefreshingAtomicMap[K, V]{refresher: refresher, refresherName: refresherName}

	err := atomicMap.Refresh(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "initialization refresh")
	}

	go atomicMap.autoRefresh(ctx, period)

	return &atomicMap, nil
}

// Refresh runs the refresher once and atomically swaps in the new contents,
// leaving the current map untouched if the refresher returns an error.
func (r *RefreshingAtomicMap[K, V]) Refresh(ctx context.Context) error {
	t0 := time.Now()

	newMap, err := r.refresher()
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	atomicMapLen := len(newMap)
	r.atomicMap.Store(newMap)

	elapsed := time.Since(t0)

	zerolog.Ctx(ctx).Info().
		Int("len", atomicMapLen).
		Str("elapsed", timeutils.RoundDuration(elapsed)).
		Str("refresher", r.refresherName).
		Msg("atomic.map.refreshed")

	return nil
}

// Get returns the value for key from the most recent successful refresh.
func (r *RefreshingAtomicMap[K, V]) Get(key K) (V, bool) {
	return r.atomicMap.Get(key)
}

func (r *RefreshingAtomicMap[K, V]) autoRefresh(ctx context.Context, period time.Duration) {
	ticker := time.NewTicker(period)

	for {
		select {
		case <-ctx.Done():
			zerolog.Ctx(ctx).Info().
				Str("refresher", r.refresherName).
				Msg("atomic.map.refreshing.cancelled")

			return
		case <-ticker.C:
			err := r.Refresh(ctx)
			if err != nil {
				zerolog.Ctx(ctx).Error().
					Err(err).
					Str("refresher", r.refresherName).
					Msg("failed.to.refresh.map")
			}
		}
	}
}
