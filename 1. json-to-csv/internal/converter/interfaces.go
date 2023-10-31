package converter

import "learn/internal/converter/models"

type Writable interface {
	WriteRow(logRaw models.LogRow) error
	WriteHead(logs []models.LogRow) error
}

type Readable interface {
	GetLogs() *[]models.LogRow
}
