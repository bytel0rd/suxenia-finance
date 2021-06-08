package objects

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var auditData AuditData = NewAuditData("Abiodun Oyegoke")

func TestAuditCreatedByName(t *testing.T) {
	assert.Error(t, auditData.SetCreatedBy(""))
}

func TestAuditUpdatedByName(t *testing.T) {
	assert.Error(t, auditData.SetCreatedBy(""))
}

func TestAuditCreatedAt(t *testing.T) {
	assert.Error(t, auditData.SetCreatedAt(time.Now()))
}

func TestAuditUpdatedAt(t *testing.T) {
	assert.Error(t, auditData.SetUpdatedAt(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)))
}
