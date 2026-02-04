# TinyKV - Distributed Key-Value Storage System

> A learning implementation of a distributed key-value storage system based on the [TinyKV Course](https://github.com/tidb-incubator/tinykv)

## About

This project is an implementation of the TinyKV course, which builds a key-value storage system with the Raft consensus algorithm. The course is inspired by [MIT 6.824](https://pdos.csail.mit.edu/6.824/) and the [TiKV Project](https://github.com/tikv/tikv).

TinyKV helps you understand how to implement a horizontally scalable, highly available, key-value storage service with distributed transaction support, along with a deeper understanding of TiKV architecture and implementation.

## Project Status

- [x] **Project 1: Standalone KV** - ✅ Completed
  - Standalone storage engine implementation
  - Raw key-value service handlers
  
- [ ] **Project 2: Raft KV** - 🚧 In Progress
  - Basic Raft algorithm
  - Fault-tolerant KV server on top of Raft
  - Raft log garbage collection and snapshot support
  
- [ ] **Project 3: Multi-raft KV** - 📋 Planned
  - Membership change and leadership change
  - Conf change and region split on Raft store
  - Basic scheduler implementation
  
- [ ] **Project 4: Transaction** - 📋 Planned
  - Multi-version concurrency control (MVCC) layer
  - Transaction handlers implementation

## Course Architecture & Projects

The project consists of four main stages, each building upon the previous one:

### Project 1: Standalone KV
[Documentation](doc/project1-StandaloneKV.md) | ✅ **Completed**
- Implement a standalone storage engine
- Implement raw key-value service handlers

### Project 2: Raft KV
[Documentation](doc/project2-RaftKV.md) | 🚧 **In Progress**
- Implement the basic Raft algorithm
- Build a fault-tolerant KV server on top of Raft
- Add support for Raft log garbage collection and snapshot

### Project 3: Multi-raft KV
[Documentation](doc/project3-MultiRaftKV.md) | 📋 **Planned**
- Implement membership change and leadership change in Raft
- Implement conf change and region split on Raft store
- Implement a basic scheduler

### Project 4: Transaction
[Documentation](doc/project4-Transaction.md) | 📋 **Planned**
- Implement the multi-version concurrency control (MVCC) layer
- Implement handlers for `KvGet`, `KvPrewrite`, and `KvCommit` requests
- Implement handlers for `KvScan`, `KvCheckTxnStatus`, `KvBatchRollback`, and `KvResolveLock` requests

## Code Structure

![overview](doc/imgs/overview.png)

Similar to the architecture of TiDB + TiKV + PD that separates storage and computation, TinyKV focuses on the storage layer of a distributed database system. If you're interested in the SQL layer, check out [TinySQL](https://github.com/tidb-incubator/tinysql). TinyScheduler acts as the central control of the TinyKV cluster, collecting information from heartbeats and generating scheduling tasks. All instances communicate via gRPC.

### Directory Structure

```
├── kv/                  # Key-value store implementation
│   ├── server/          # gRPC server and service handlers
│   ├── storage/         # Storage interface and implementations
│   ├── raftstore/       # Raft-based storage engine
│   ├── transaction/     # Transaction layer (MVCC)
│   └── coprocessor/     # Coprocessor for data processing
├── raft/                # Raft consensus algorithm implementation
├── scheduler/           # TinyScheduler for cluster management
│   ├── server/          # Scheduler server and coordinator
│   └── client/          # Scheduler client
├── proto/               # Protocol Buffers definitions
└── log/                 # Logging utilities
```

## Learning Resources

### Prerequisites

- **Git**: [Install Git](https://git-scm.com/downloads)
- **Go**: Version ≥ 1.13 ([Installation Guide](https://golang.org/doc/install))

### Build

Build TinyKV from source:

```bash
cd go-distributed-system
make
```

This builds the `tinykv-server` and `tinyscheduler-server` binaries in the `bin/` directory.

### Running the Server

Run the standalone server:

```bash
make
./bin/tinykv-server
```
### Testing

Run tests for specific components:

```bash
# Test standalone storage
make project1

# Test Raft implementation
make project2

# Test Multi-Raft
make project3

# Test transactions
make project4
```

## Integration with TinySQL

You can run TinyKV with TinySQL to get a complete distributed database experience:

1. Get `tinysql-server` following [TinySQL documentation](https://github.com/tidb-incubator/tinysql#deploy)
2. Place `tinyscheduler-server`, `tinykv-server`, and `tinysql-server` binaries in the same directory
3. Run the following commands:

```bash
mkdir -p data
./tinyscheduler-server
./tinykv-server -path=data
./tinysql-server --store=tikv --path="127.0.0.1:2379"
```

4. Connect with MySQL client:

```bash
mysql -u root -h 127.0.0.1 -P 4000
```

## Learning Resources

### Reading List

Check out the [reading list](doc/reading_list.md) for resources on distributed storage systems.

### Recommended Reading

- **TiKV Design**: Data storage architecture
  - [English](https://en.pingcap.com/blog/tidb-internal-data-storage)
  - [Chinese](https://pingcap.com/zh/blog/tidb-internal-1)
- **PD Design**: Scheduling system
  - [English](https://en.pingcap.com/blog/tidb-internal-scheduling)
  - [Chinese](https://pingcap.com/zh/blog/tidb-internal-3)
- **Raft Paper**: [The Raft Consensus Algorithm](https://raft.github.io/raft.pdf)
- **Raft Visualization**: [Interactive Raft](https://raft.github.io/)

## Acknowledgments

This project is based on the [TinyKV Course](https://github.com/tidb-incubator/tinykv) by PingCAP, part of the [Talent Plan](https://github.com/pingcap/talent-plan) educational initiative.

## License

See [LICENSE](LICENSE) file for details.