package backuper

import (
	"context"
)

type PGRepo interface {
	IsBotDBUpdated(ctx context.Context, now int64) (bool, error)
}
