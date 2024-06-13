package shardmaster

import (
	"fmt"
	"net/rpc"
	"time"
)

type MasterClerk struct {
	replicaAddresses []string // Addresses of shardmaster replicas
}

func NewMasterClerk(replicaAddresses []string) *MasterClerk {
	masterClerk := &MasterClerk{
		replicaAddresses: replicaAddresses,
	}
	return masterClerk
}
// remotecall() sends an RPC to the rpcname handler on server srv
func remoteCall(serverAddress, methodName string, args interface{}, reply interface{}) bool {
	connection, err := rpc.Dial("unix", serverAddress)
	if err != nil {
		return false
	}
	defer connection.Close()

	callErr := connection.Call(methodName, args, reply)
	if callErr == nil {
		return true
	}

	fmt.Println(callErr)
	return false
}

func (mc *MasterClerk) QueryConfig(shardNumber int) Config {
	for {
		for _, serverAddress := range mc.replicaAddresses {
			queryArgs := &QueryArgs{Num: shardNumber}
			var queryReply QueryReply

			if success := remoteCall(serverAddress, "ShardMaster.Query", queryArgs, &queryReply); success {
				return queryReply.Config
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (mc *MasterClerk) JoinGroup(groupID int64, servers []string) {
	for {
		for _, serverAddress := range mc.replicaAddresses {
			joinArgs := &JoinArgs{GID: groupID, Servers: servers}
			var joinReply JoinReply

			if success := remoteCall(serverAddress, "ShardMaster.Join", joinArgs, &joinReply); success {
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (mc *MasterClerk) LeaveGroup(groupID int64) {
	for {
		for _, serverAddress := range mc.replicaAddresses {
			leaveArgs := &LeaveArgs{GID: groupID}
			var leaveReply LeaveReply

			if success := remoteCall(serverAddress, "ShardMaster.Leave", leaveArgs, &leaveReply); success {
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (mc *MasterClerk) MoveShard(shardID int, groupID int64) {
	for {
		for _, serverAddress := range mc.replicaAddresses {
			moveArgs := &MoveArgs{Shard: shardID, GID: groupID}
			var moveReply MoveReply

			if success := remoteCall(serverAddress, "ShardMaster.Move", moveArgs, &moveReply); success {
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
