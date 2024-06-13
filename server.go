package shardmaster

import (
	"sync/atomic"
)

func (sm *ShardMaster) Join(args *JoinArgs, reply *JoinReply) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if atomic.LoadInt32(&sm.dead) == 1 {
		return nil // Return if the server is dead
	}

	// Assume Paxos is used to agree on the new configuration
	config := sm.createNewConfig()
	config.Groups[args.GID] = args.Servers
	// Increment config number
	config.Num = len(sm.configs)
	sm.configs = append(sm.configs, config)

	// Actual implementation would involve a Paxos agreement to ensure all servers in the cluster
	// agree on the new configuration.

	return nil
}

func (sm *ShardMaster) Leave(args *LeaveArgs, reply *LeaveReply) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if atomic.LoadInt32(&sm.dead) == 1 {
		return nil // Return if the server is dead
	}

	// Assume Paxos is used to agree on the new configuration
	config := sm.createNewConfig()
	delete(config.Groups, args.GID)
	// Increment config number
	config.Num = len(sm.configs)
	sm.configs = append(sm.configs, config)

	// Actual implementation would involve a Paxos agreement to ensure all servers in the cluster
	// agree on the new configuration.

	return nil
}

func (sm *ShardMaster) Move(args *MoveArgs, reply *MoveReply) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if atomic.LoadInt32(&sm.dead) == 1 {
		return nil // Return if the server is dead
	}

	// Assume Paxos is used to agree on the new configuration
	config := sm.createNewConfig()
	config.Shards[args.Shard] = args.GID
	// Increment config number
	config.Num = len(sm.configs)
	sm.configs = append(sm.configs, config)

	// Actual implementation would involve a Paxos agreement to ensure all servers in the cluster
	// agree on the new configuration.

	return nil
}

func (sm *ShardMaster) Query(args *QueryArgs, reply *QueryReply) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if atomic.LoadInt32(&sm.dead) == 1 {
		return nil // Return if the server is dead
	}

	if args.Num == -1 || args.Num >= len(sm.configs) {
		reply.Config = sm.configs[len(sm.configs)-1]
	} else {
		reply.Config = sm.configs[args.Num]
	}

	return nil
}

// Helper function to create a new configuration by copying the latest one.
// This function assumes that the caller holds the lock on sm.mu.
func (sm *ShardMaster) createNewConfig() Config {
	latestConfig := sm.configs[len(sm.configs)-1]
	newConfig := Config{
		Num:    latestConfig.Num + 1,
		Shards: latestConfig.Shards, // Copy the shards assignment
		Groups: make(map[int64][]string),
	}
	for gid, servers := range latestConfig.Groups {
		newConfig.Groups[gid] = append([]string(nil), servers...)
	}
	return newConfig
}
