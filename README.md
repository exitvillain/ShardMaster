ShardMaster: A Key-Value Store for Distributed Systems

Overview

ShardMaster is a robust key-value store designed for distributed systems, emphasizing fault tolerance and scalability. This project was developed as part of my Distributed Systems course at NYU, where I explored the complexities of maintaining consistency and availability in a networked environment. The repository, named shardmaster, includes essential components for managing a dynamic set of replica groups in a distributed database system.

Technical Implementation

The core functionality of ShardMaster revolves around managing shard allocations across different server groups to ensure efficient data distribution and access. The system utilizes the Paxos consensus algorithm to achieve agreement across the network, ensuring that all changes to the system's state are consistent, even in the face of network failures or concurrent updates.

Server Management: The server.go file contains methods like Join, Leave, Move, and Query that handle server group configurations and shard assignments. These operations allow for dynamic adjustments in the distributed system's structure, responding to various conditions such as load balancing and server failures.

Client Interaction: The client.go file demonstrates how clients interact with the ShardMaster through remote procedure calls (RPCs). This includes joining new server groups to the system, leaving existing groups, moving shards between groups, and querying the current configuration.

Configuration and Utilities: The commons.go file defines the data structures and constants used across the system, including the configuration of shards, replica groups, and the sequence in which changes are applied.

Potential Applications

ShardMaster is particularly useful in environments where high availability and reliable data access are crucial. It can serve as the backbone for applications requiring resilient data storage and retrieval across a distributed set of servers, such as large-scale web applications, real-time data processing systems, and cloud storage services.
