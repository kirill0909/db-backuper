package backuper

import (
	"context"
	"time"
)

type Usecase interface {
	PGBotDBBackup(ctx context.Context, now time.Time) error
}
