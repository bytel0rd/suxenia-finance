package mappers

import (
	"database/sql"
	"suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/dtos"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"
)

func BankingKycAggregateFromPersistence(kycEntity entities.BankingKycEntity) (*aggregates.BankingKYC, error) {

	var error error = nil

	kycAggregate := aggregates.BankingKYC{}

	if error = kycAggregate.SetId(&kycEntity.Id); error != nil {
		return nil, error
	}

	if error = kycAggregate.SetName(&kycEntity.Name); error != nil {
		return nil, error
	}

	if kycEntity.BankAccountName.Valid {

		if error = kycAggregate.SetBankAccountName(&kycEntity.BankAccountName.String); error != nil {
			return nil, error
		}

	}

	if kycEntity.BankAccountNumber.Valid {

		if error = kycAggregate.SetBankAccountNumber(&kycEntity.BankAccountNumber.String); error != nil {
			return nil, error
		}

	}

	if kycEntity.BVN.Valid {

		if error = kycAggregate.SetBVN(&kycEntity.BVN.String); error != nil {
			return nil, error
		}

	}

	if kycEntity.BankCode.Valid {

		if error = kycAggregate.SetBankCode(&kycEntity.BankCode.String); error != nil {
			return nil, error
		}

	}

	if error = kycAggregate.SetOwnerId(&kycEntity.OwnerId); error != nil {
		return nil, error
	}

	if error = kycAggregate.SetVerificationStatus(kycEntity.Verified); error != nil {
		return nil, error
	}

	kycAggregate.AuditData.SetCreatedAt(kycEntity.CreatedAt)
	kycAggregate.AuditData.SetUpdatedAt(kycEntity.UpdateAt)
	kycAggregate.AuditData.SetUpdatedBy(kycEntity.CreatedBy)
	kycAggregate.AuditData.SetUpdatedBy(kycEntity.UpdatedBy)

	return &kycAggregate, nil

}

func BankingKycAggregateToPersistence(kycAggregate aggregates.BankingKYC) (*entities.BankingKycEntity, error) {

	kycEntity := entities.BankingKycEntity{}

	if id, ok := kycAggregate.GetId(); ok {
		kycEntity.Id = *id
	}

	if name, ok := kycAggregate.GetName(); ok {
		kycEntity.Name = *name
	}

	if acctName, ok := kycAggregate.GetBankAccountName(); ok {
		kycEntity.BankAccountName = sql.NullString{String: *acctName, Valid: ok}
	}

	if acctNo, ok := kycAggregate.GetBankAccountNumber(); ok {
		kycEntity.BankAccountNumber = sql.NullString{String: *acctNo, Valid: ok}
	}

	if bvn, ok := kycAggregate.GetBVN(); ok {
		kycEntity.BVN = sql.NullString{String: *bvn, Valid: ok}
	}

	if bankCode, ok := kycAggregate.GetBankCode(); ok {
		kycEntity.BankCode = sql.NullString{String: *bankCode, Valid: ok}
	}

	if ownerId, ok := kycAggregate.GetOwnerId(); ok {
		kycEntity.OwnerId = *ownerId
	}

	kycEntity.Verified = kycAggregate.IsVerified()

	kycEntity.CreatedAt = kycAggregate.GetCreatedAt()
	kycEntity.UpdateAt = kycAggregate.GetUpdatedAt()
	kycEntity.CreatedBy = kycAggregate.GetCreatedBy()
	kycEntity.UpdatedBy = kycAggregate.GetUpdatedBy()

	return &kycEntity, nil
}

func BankingKycAggregateToViewModel(kycAggregate aggregates.BankingKYC) (*dtos.BankingKycViewModel, error) {

	kycEntity := dtos.BankingKycViewModel{}

	if id, ok := kycAggregate.GetId(); ok {
		kycEntity.Id = *id
	}

	if name, ok := kycAggregate.GetName(); ok {
		kycEntity.Name = *name
	}

	if acctName, ok := kycAggregate.GetBankAccountName(); ok {
		kycEntity.BankAccountName = acctName
	}

	if acctNo, ok := kycAggregate.GetBankAccountNumber(); ok {
		kycEntity.BankAccountNumber = acctNo
	}

	if bvn, ok := kycAggregate.GetBVN(); ok {
		kycEntity.BVN = bvn
	}

	if bankCode, ok := kycAggregate.GetBankCode(); ok {
		kycEntity.BankCode = bankCode
	}

	if ownerId, ok := kycAggregate.GetOwnerId(); ok {
		kycEntity.OwnerId = *ownerId
	}

	kycEntity.Verified = kycAggregate.IsVerified()

	kycEntity.CreatedAt = kycAggregate.GetCreatedAt()
	kycEntity.UpdateAt = kycAggregate.GetUpdatedAt()
	kycEntity.CreatedBy = kycAggregate.GetCreatedBy()
	kycEntity.UpdatedBy = kycAggregate.GetUpdatedBy()

	return &kycEntity, nil
}
