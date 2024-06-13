// The master shard server is responsible for allocating shards to different replication groups.
//
// RPC interface summary:
// Join(gid, servers) - Adds a new replica group identified by 'gid' and allocates some shards to it.
// Leave(gid) - Removes the replica group 'gid' and redistributes its shards among the remaining groups.
// Move(shard, gid) - Transfers ownership of a specific 'shard' to the group identified by 'gid'.
// Query(num) - Retrieves the configuration identified by 'num', or the most recent configuration if 'num' is -1.
//
// 'Config' structures represent the shard allocation across different replica groups. Each 'Config' is sequentially numbered.
// The initial configuration, numbered as 0, assigns all shards to an invalid group (group 0).
//
// 'GID' represents the unique identifier for a replica group, and it must be greater than 0.
// A 'GID' should not be reused once its corresponding group has joined and then left the system.


const NShards = 10  // Number of shards

type Config struct {
    Num    int                // Config number
    Shards [NShards]int64     // Map of shard to gid
    Groups map[int64][]string // Map of gid to server list
}

type JoinArgs struct {
    GID     int64    // Unique replica group ID
    Servers []string // List of server ports in the group
}

type JoinReply struct {}

type LeaveArgs struct {
    GID int64 // Replica group ID to remove
}

type LeaveReply struct {}

type MoveArgs struct {
    Shard int   // Shard number
    GID   int64 // Target group ID for the shard
}

type MoveReply struct {}

type QueryArgs struct {
    Num int // Config number to fetch
}

type QueryReply struct {
    Config Config // The requested configuration
}