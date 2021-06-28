package mappers

import (
	"suxenia-finance/pkg/common/utils"
	"suxenia-finance/pkg/wallet/dtos"
	"suxenia-finance/pkg/wallet/infrastructure/persistence/entities"
)

func PaymentEntityToView(payment entities.Payment) dtos.PaymentViewModel {

	view := dtos.PaymentViewModel{}

	view.Id = payment.Id
	view.ProcessedBy = payment.ProcessedBy

	view.OpeningBalance = &payment.OpeningBalance.Int32

	view.FormatedOpeningBalance = utils.IntToDecimal(int(payment.OpeningBalance.Int32)).RoundBank(2).String()

	view.Amount = payment.Amount
	view.FormatedAmount = utils.IntToDecimal(view.Amount).RoundBank(2).String()

	view.TransactionReference = payment.TransactionReference
	view.TransactionSource = payment.TransactionSource

	view.SourceReference = &payment.SourceReference.String

	view.Platform = payment.Platform
	view.Status = payment.Status

	view.Comments = payment.Comments
	view.OwnerId = payment.Id

	view.AuditInfo = payment.AuditInfo

	return view
}
