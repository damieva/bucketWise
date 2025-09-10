package transaction

import "bucketWise/pkg/ports"

type Service struct {
	Repo ports.TransactionRepository
}
