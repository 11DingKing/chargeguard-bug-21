package charging

import "errors"

var ErrEmptyClosureBatch = errors.New("empty closure batch")

type BatchSummary struct {
	Accepted int
	Rejected int
}

func CloseBatch(ids []string) (*BatchSummary, error) {
	if len(ids) == 0 {
		return nil, ErrEmptyClosureBatch
	}
	return &BatchSummary{Accepted: len(ids)}, nil
}
