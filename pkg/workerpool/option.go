package workerpool

type Option func( *Pool)

func WithBlock(block bool) Option  {
	return func(pool *Pool) {
		pool.block=block
	}

}
func WithPreAllocWorkers(preAlloc bool) Option  {
	return func(pool *Pool) {
		pool.preAlloc=preAlloc
	}
}