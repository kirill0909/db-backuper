package backuper

import (
	"context"
	"time"
)

type Backuper interface {
	PGBotDBBackup(ctx context.Context, now time.Time) error
}
