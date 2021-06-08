package common

import "time"

type AuditInfo struct {
	CreatedBy string

	UpdatedBy string

	CreatedAt time.Time

	UpdateAt time.Time
}
