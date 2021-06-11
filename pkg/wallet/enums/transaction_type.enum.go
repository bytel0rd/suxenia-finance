package enums

type TransactionType int

const (
	PAYMENT TransactionType = iota + 1
	WITHDRAWAL
)

func (t *TransactionType) names() []string {
	return []string{"PAYMENT", "WITHDRAWAL"}
}

func (t TransactionType) String() string {

	return t.names()[int(t)]
}

func (t TransactionType) Value() int {
	return int(t)
}
