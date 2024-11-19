package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	pb "auction/stc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type AuctionClient struct {
	serverAddrs []string
	clients     map[string]*grpc.ClientConn
	mu          sync.RWMutex
	currentIdx  int
}

func NewAuctionClient(serverAddrs []string) (*AuctionClient, error) {
	if len(serverAddrs) == 0 {
		return nil, fmt.Errorf("at least one server address is required")
	}

	client := &AuctionClient{
		serverAddrs: serverAddrs,
		clients:     make(map[string]*grpc.ClientConn),
		currentIdx:  0,
	}

	// Initially connect to all servers
	for _, addr := range serverAddrs {
		if err := client.connectToServer(addr); err != nil {
			// log.Printf("Warning: Failed to connect to server %s: %v", addr, err)
		}
	}

	if len(client.clients) == 0 {
		return nil, fmt.Errorf("failed to connect to any server")
	}

	return client, nil
}

func (c *AuctionClient) connectToServer(addr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.clients[addr] = conn
	c.mu.Unlock()
	return nil
}

func (c *AuctionClient) getNextValidClient() (pb.AuctionServiceClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	startIdx := c.currentIdx
	for i := 0; i < len(c.serverAddrs); i++ {
		idx := (startIdx + i) % len(c.serverAddrs)
		addr := c.serverAddrs[idx]

		// Check if we have a connection to this server
		conn, exists := c.clients[addr]
		if !exists {
			// Try to establish connection if it doesn't exist
			if err := c.connectToServer(addr); err != nil {
				// log.Printf("Failed to connect to server %s: %v", addr, err)
				continue
			}
			conn = c.clients[addr]
		}

		// Check connection state
		state := conn.GetState()
		if state != connectivity.Ready {
			// Try to reconnect if not ready
			if err := c.connectToServer(addr); err != nil {
				// log.Printf("Failed to reconnect to server %s: %v", addr, err)
				delete(c.clients, addr)
				continue
			}
			conn = c.clients[addr]
		}

		// Update current index and return client
		c.currentIdx = (idx + 1) % len(c.serverAddrs)
		return pb.NewAuctionServiceClient(conn), nil
	}

	return nil, fmt.Errorf("no available servers")
}

func (c *AuctionClient) PlaceBid(amount int32, bidderID string) error {
	req := &pb.BidRequest{
		Amount:   amount,
		BidderId: bidderID,
	}

	// Try each server until successful or all fail
	var lastErr error
	for i := 0; i < len(c.serverAddrs); i++ {
		client, err := c.getNextValidClient()
		if err != nil {
			lastErr = err
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := client.PlaceBid(ctx, req)
		cancel()

		if err != nil {
			log.Printf("Failed to place bid with current server: %v", err)
			lastErr = err
			continue
		}

		if !resp.Success {
			return fmt.Errorf("bid failed: %s", resp.Error)
		}

		log.Printf("Bid placed successfully")
		return nil
	}

	return fmt.Errorf("failed to place bid with any server: %v", lastErr)
}

func (c *AuctionClient) GetResult() error {
	// Try each server until successful or all fail
	var lastErr error
	for i := 0; i < len(c.serverAddrs); i++ {
		client, err := c.getNextValidClient()
		if err != nil {
			lastErr = err
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := client.GetResult(ctx, &pb.ResultRequest{})
		cancel()

		if err != nil {
			log.Printf("Failed to get result from current server: %v", err)
			lastErr = err
			continue
		}

		if resp.IsEnded {
			log.Printf("Auction ended. Winner: %s, Winning bid: %d", resp.Winner, resp.HighestBid)
		} else {
			log.Printf("Auction in progress. Current highest bid: %d by %s", resp.HighestBid, resp.Winner)
		}
		return nil
	}

	return fmt.Errorf("failed to get result from any server: %v", lastErr)
}

func (c *AuctionClient) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for addr, conn := range c.clients {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection to %s: %v", addr, err)
		}
	}
}

func main() {
	servers := flag.String("servers", "localhost:8001", "Comma-separated list of server addresses")
	action := flag.String("action", "", "Action to perform (bid/result)")
	amount := flag.Int("amount", 0, "Bid amount")
	bidderID := flag.String("bidder", "", "Bidder ID")
	flag.Parse()

	serverList := strings.Split(*servers, ",")
	client, err := NewAuctionClient(serverList)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	switch *action {
	case "bid":
		if *amount <= 0 || *bidderID == "" {
			log.Fatal("Bid amount and bidder ID are required for bidding")
		}
		err = client.PlaceBid(int32(*amount), *bidderID)
	case "result":
		err = client.GetResult()
	default:
		log.Fatal("Invalid action. Use 'bid' or 'result'")
	}

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
