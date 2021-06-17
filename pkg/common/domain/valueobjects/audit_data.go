package objects

import (
	"errors"
	"time"
)

type AuditData struct {
	createdBy string

	updatedBy string

	createdAt time.Time

	updateAt time.Time
}

func (data *AuditData) GetCreatedBy() string {
	return data.createdBy
}

func (data *AuditData) SetCreatedBy(name string) error {

	if name == "" {
		return errors.New("missing parameter: createdBy is required ")
	}

	data.createdBy = name

	return nil
}

func (data *AuditData) GetUpdatedBy() string {
	return data.updatedBy
}

func (data *AuditData) SetUpdatedBy(name string) error {

	if name == "" {
		return errors.New("missing parameter: updateName is required ")
	}

	data.updatedBy = name

	return nil
}

func (data *AuditData) GetCreatedAt() time.Time {
	return data.createdAt
}

func (data *AuditData) SetCreatedAt(timeCreated time.Time) error {

	if timeCreated.After((data.updateAt)) {
		return errors.New("data create time cannot be after date it was updated")
	}

	data.createdAt = timeCreated

	return nil
}

func (data *AuditData) GetUpdatedAt() time.Time {

	return data.updateAt
}

func (data *AuditData) SetUpdatedAt(timeUpdated time.Time) error {

	if timeUpdated.Before((data.createdAt)) {
		return errors.New("data updated time cannot be before the time it was created")
	}

	data.updateAt = timeUpdated

	return nil
}

func NewAuditData(creator string) AuditData {

	return AuditData{

		createdBy: creator,
		updatedBy: creator,

		createdAt: time.Now(),
		updateAt:  time.Now(),
	}

}
