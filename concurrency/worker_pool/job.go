package worker_pool

import "context"

type (
	job struct {
		executeFn func(ctx context.Context) (interface{}, error)
	}
	result struct {
		value interface{}
		err   error
	}
)

func (j job) execute(ctx context.Context) result {
	v, err := j.executeFn(ctx)
	if err != nil {
		return result{
			value: nil,
			err:   err,
		}
	}
	return result{
		value: v,
		err:   nil,
	}
}
