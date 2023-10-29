package backuper

import (
	"context"
	"time"
)

type PGRepo interface {
	IsBotDBUpdated(ctx context.Context, now time.Time) (bool, error)
}
