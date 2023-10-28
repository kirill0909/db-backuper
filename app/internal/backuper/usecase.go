package backuper

import (
	"context"
)

type Backuper interface {
	PGBackup(ctx context.Context)
}
