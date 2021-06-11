package enums

type TransactionStatus int

const (
	PENDING TransactionStatus = iota + 1
	SUCCESS
	FAILED
	REJECTED
)

func (t *TransactionStatus) names() []string {
	return []string{"PENDING", "SUCCESS", "FAILED", "REJECTED"}
}

func (t TransactionStatus) String() string {

	return t.names()[int(t)]
}

func (t TransactionStatus) Value() int {
	return int(t)
}
