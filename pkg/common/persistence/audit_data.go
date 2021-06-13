package persistence

import "time"

type AuditInfo struct {
	CreatedBy string `db:"created_by" json:"createdBy"`

	UpdatedBy string `db:"updated_by" json:"updatedBy"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`

	UpdateAt time.Time `db:"updated_at" json:"updateAt"`
}

func NewAuditInfo(creator string) AuditInfo {
	return AuditInfo{
		CreatedBy: creator,
		UpdatedBy: creator,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
