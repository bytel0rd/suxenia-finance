package mappers

import (
	"database/sql"
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/kyc/domain/aggregates"
	"suxenia-finance/pkg/kyc/dtos"
	"suxenia-finance/pkg/kyc/infrastructure/persistence/entities"
)

func BankingKycAggregateFromPersistence(kycEntity entities.BankingKycEntity) aggregates.BankingKYC {

	kycAggregate := aggregates.BankingKYC{}

	kycAggregate.SetId(&kycEntity.Id)
	kycAggregate.SetName(&kycEntity.Name)
	kycAggregate.SetOwnerId(&kycEntity.OwnerId)
	kycAggregate.SetVerificationStatus(kycEntity.Verified)

	if kycEntity.BankAccountName.Valid {
		kycAggregate.SetBankAccountName(&kycEntity.BankAccountName.String)
	}

	if kycEntity.BankAccountNumber.Valid {
		kycAggregate.SetBankAccountNumber(&kycEntity.BankAccountNumber.String)
	}

	if kycEntity.BVN.Valid {
		kycAggregate.SetBVN(&kycEntity.BVN.String)
	}

	if kycEntity.BankCode.Valid {
		kycAggregate.SetBankCode(&kycEntity.BankCode.String)
	}

	kycAggregate.AuditData.SetCreatedAt(kycEntity.CreatedAt)
	kycAggregate.AuditData.SetUpdatedAt(kycEntity.UpdateAt)
	kycAggregate.AuditData.SetUpdatedBy(kycEntity.CreatedBy)
	kycAggregate.AuditData.SetUpdatedBy(kycEntity.UpdatedBy)

	return kycAggregate

}

func BankingKycAggregateToPersistence(kycAggregate aggregates.BankingKYC) entities.BankingKycEntity {

	kycEntity := entities.BankingKycEntity{}

	id, _ := kycAggregate.GetId()

	if id != nil {
		kycEntity.Id = *id
	}

	name, _ := kycAggregate.GetName()

	if name != nil {
		kycEntity.Name = *name
	}

	acctName, _ := kycAggregate.GetBankAccountName()

	if acctName != nil {
		kycEntity.BankAccountName = sql.NullString{String: *acctName, Valid: utils.IsValidStringPointer(acctName)}
	}

	acctNo, _ := kycAggregate.GetBankAccountNumber()

	if acctNo != nil {
		kycEntity.BankAccountNumber = sql.NullString{String: *acctNo, Valid: utils.IsValidStringPointer(acctNo)}
	}

	bvn, _ := kycAggregate.GetBVN()

	if bvn != nil {
		kycEntity.BVN = sql.NullString{String: *bvn, Valid: utils.IsValidStringPointer(bvn)}
	}

	bankCode, _ := kycAggregate.GetBankCode()
	{

		if bankCode != nil {
			kycEntity.BankCode = sql.NullString{String: *bankCode, Valid: utils.IsValidStringPointer(bankCode)}
		}

	}

	ownerId, _ := kycAggregate.GetOwnerId()

	if ownerId != nil {
		kycEntity.OwnerId = *ownerId
	}

	kycEntity.Verified = kycAggregate.IsVerified()

	kycEntity.CreatedAt = kycAggregate.GetCreatedAt()
	kycEntity.UpdateAt = kycAggregate.GetUpdatedAt()
	kycEntity.CreatedBy = kycAggregate.GetCreatedBy()
	kycEntity.UpdatedBy = kycAggregate.GetUpdatedBy()

	return kycEntity
}

func BankingKycAggregateToViewModel(kycAggregate aggregates.BankingKYC) dtos.BankingKycViewModel {

	kycEntity := dtos.BankingKycViewModel{}

	if id, error := kycAggregate.GetId(); error == nil {
		kycEntity.Id = *id
	}

	if name, error := kycAggregate.GetName(); error == nil {
		kycEntity.Name = *name
	}

	if acctName, error := kycAggregate.GetBankAccountName(); error == nil {
		kycEntity.BankAccountName = acctName
	}

	if acctNo, error := kycAggregate.GetBankAccountNumber(); error == nil {
		kycEntity.BankAccountNumber = acctNo
	}

	if bvn, error := kycAggregate.GetBVN(); error == nil {
		kycEntity.BVN = bvn
	}

	if bankCode, error := kycAggregate.GetBankCode(); error == nil {
		kycEntity.BankCode = bankCode
	}

	if ownerId, error := kycAggregate.GetOwnerId(); error == nil {
		kycEntity.OwnerId = *ownerId
	}

	kycEntity.Verified = kycAggregate.IsVerified()

	kycEntity.CreatedAt = kycAggregate.GetCreatedAt()
	kycEntity.UpdateAt = kycAggregate.GetUpdatedAt()
	kycEntity.CreatedBy = kycAggregate.GetCreatedBy()
	kycEntity.UpdatedBy = kycAggregate.GetUpdatedBy()

	return kycEntity
}

// func BankingKycEntityToViewModel(kycEntity entities.BankingKycEntity) aggregates.BankingKYC {

// 	kycAggregate := aggregates.BankingKYC{}

// 	kycAggregate.SetId(&kycEntity.Id)
// 	kycAggregate.SetName(&kycEntity.Name)
// 	kycAggregate.SetOwnerId(&kycEntity.OwnerId)
// 	kycAggregate.SetVerificationStatus(kycEntity.Verified)

// 	if kycEntity.BankAccountName.Valid {
// 		kycAggregate.SetBankAccountName(&kycEntity.BankAccountName.String)
// 	}

// 	if kycEntity.BankAccountNumber.Valid {
// 		kycAggregate.SetBankAccountNumber(&kycEntity.BankAccountNumber.String)
// 	}

// 	if kycEntity.BVN.Valid {
// 		kycAggregate.SetBVN(&kycEntity.BVN.String)
// 	}

// 	if kycEntity.BankCode.Valid {
// 		kycAggregate.SetBankCode(&kycEntity.BankCode.String)
// 	}

// 	kycAggregate.AuditData.SetCreatedAt(kycEntity.CreatedAt)
// 	kycAggregate.AuditData.SetUpdatedAt(kycEntity.UpdateAt)
// 	kycAggregate.AuditData.SetUpdatedBy(kycEntity.CreatedBy)
// 	kycAggregate.AuditData.SetUpdatedBy(kycEntity.UpdatedBy)

// 	return kycAggregate

// }
