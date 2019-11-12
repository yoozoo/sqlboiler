package boil

// Shard finds the shard by shard key
func Shard(sk int64) Executor {
	shardIndex := sk % int64(len(currentShards))
	return currentShards[shardIndex]
}

// ContextShard finds the context shard by shard key
func ContextShard(sk int64) Executor {
	shardIndex := sk % int64(len(currentContextShards))
	return currentContextShards[shardIndex]
}
