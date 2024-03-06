package repository

const (
	queryIsBotDBUpdated = `
	SELECT EXISTS
	(SELECT id FROM infos
	WHERE EXTRACT(EPOCH FROM created_at) >= $1 OR EXTRACT(EPOCH FROM updated_at) >= $1)
	`
)
