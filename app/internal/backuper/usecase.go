package backuper

import (
	"context"
)

type Backuper interface {
	PGBotDBBackup(ctx context.Context) error
}
