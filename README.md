# Distributed Auction System (gRPC)

A distributed auction system implemented in Go using gRPC for communication. The system provides fault tolerance through replication and can withstand node failure while maintaining consistency and availability.

## Features

- gRPC-based communication between nodes
- Leader-follower replication architecture
- Fault tolerance for one node failure
- Automatic leader election
- Heartbeat-based failure detection
- Support for multiple bidders

## Prerequisites

- Go 1.16 or higher
- Protocol Buffers compiler (`protoc`)
- Go gRPC tools

## Installation

1. Build the server and client:

```zsh
# Build server
go build -o auction-server server.go

# Build client
go build -o auction-client client.go
```

## Project Structure

```
.
├── auction.proto        # Protocol Buffers definition
├── server.go           # Server implementation
└── client.go           # Client implementation
```

## Running the System

### Starting the Nodes

1. Start the leader node:
```zsh
./auction-server -addr localhost:8001 -peers localhost:8002,localhost:8003
```

2. Start follower nodes:
```zsh
./auction-server -addr localhost:8002 -peers localhost:8001,localhost:8003
./auction-server -addr localhost:8003 -peers localhost:8001,localhost:8002
```

### Using the Client

1. Place a bid:
```zsh
./auction-client -servers "localhost:8001,localhost:8002,localhost:8003" -action bid -amount 100 -bidder user1
```

2. Query auction result:
```zsh
./auction-client -servers "localhost:8001,localhost:8002,localhost:8003" -action result
```

(you can also try only giving the client only one server address and see that it maintains consistency)

## Testing Fault Tolerance

You can test the system's fault tolerance by:

1. Starting multiple nodes
2. Placing some bids
3. Killing one of the nodes (including the leader)
4. Continuing to place bids and query results

The system should:
- Detect the node failure
- Elect a new leader if needed
- Continue processing bids
- Maintain consistency across remaining nodes

## System Architecture

### Components

1. **Server (server.go)**
   - Manages auction state
   - Handles replication
   - Implements failure detection
   - Manages leader election

2. **Client (client.go)**
   - Provides a command-line interface
   - Connects to any node in the cluster
   - Supports bidding and result queries

### Communication Flow

1. **Client Operations**
   - Bids and queries can be sent to any node
   - Nodes forward writes to the leader if necessary

2. **Replication**
   - Leader replicates all bids to followers
   - Uses synchronous replication for consistency

3. **Failure Detection**
   - Nodes exchange periodic heartbeats
   - Failed nodes are detected and removed from the cluster
   - Leader election is triggered if the leader fails

## Monitoring and Logging

The system logs important events including:
- Node startup
- Leader election
- Bid placement
- State replication
- Node failures
- Result requests

## Configuration

Key configurable parameters (in server.go):
```go
const (
    AuctionDuration   = 100 * time.Second
    HeartbeatInterval = 1 * time.Second
    NodeTimeout       = 3 * time.Second
)
```

## Error Handling

The system handles various error conditions:
- Network failures
- Node crashes
- Invalid bids
- Auction timeouts
- Replication failures

## Limitations

- Supports only fail-stop failures (not Byzantine failures)
- Requires at least two nodes for fault tolerance
- Simple leader election mechanism

## Future Improvements

Possible enhancements:
1. Implement more sophisticated leader election
2. Add support for multiple concurrent auctions
3. Implement bid history queries