package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "auction/stc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	AuctionDuration   = 100 * time.Second
	HeartbeatInterval = 1 * time.Second
	NodeTimeout       = 3 * time.Second
)

type Bid struct {
	Amount    int32
	BidderID  string
	Timestamp time.Time
}

type AuctionServer struct {
	pb.UnimplementedAuctionServiceServer

	address  string
	isLeader bool
	mu       sync.RWMutex
	mu2      sync.Mutex

	// Auction state
	currentBid   Bid
	bidHistory   []Bid
	auctionStart time.Time
	auctionEnded bool

	// Node management
	peers         map[string]pb.AuctionServiceClient
	lastHeartbeat map[string]time.Time
	peerAddresses map[pb.AuctionServiceClient]string
}

func NewAuctionServer(address string) *AuctionServer {
	return &AuctionServer{
		address:       address,
		peers:         make(map[string]pb.AuctionServiceClient),
		lastHeartbeat: make(map[string]time.Time),
		peerAddresses: make(map[pb.AuctionServiceClient]string),
		auctionStart:  time.Now(),
		bidHistory:    make([]Bid, 0),
	}
}

func (s *AuctionServer) PlaceBid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Recived bid %d from %s", req.Amount, req.BidderId)

	if time.Since(s.auctionStart) > AuctionDuration {
		s.auctionEnded = true
		log.Println("Auction has ended")
		return &pb.BidResponse{
			Success: false,
			Error:   "auction has ended",
		}, nil
	}

	if s.currentBid.Amount >= req.Amount {
		log.Println("Bid too low")
		return &pb.BidResponse{
			Success: false,
			Error:   "bid too low",
		}, nil
	}

	newBid := Bid{
		Amount:    req.Amount,
		BidderID:  req.BidderId,
		Timestamp: time.Now(),
	}

	// If we're the leader, replicate to followers
	if s.isLeader {
		if err := s.replicateBidToFollowers(ctx, newBid); err != nil {
			log.Println("failed to replicate bid", err)
			return &pb.BidResponse{
				Success: false,
				Error:   "failed to replicate bid",
			}, nil
		}
		s.currentBid = newBid
		s.mu2.Lock()
		s.bidHistory = append(s.bidHistory, newBid)
		s.mu2.Unlock()
		return &pb.BidResponse{Success: true}, nil

	} else {
		lowest := strings.Split(s.address, ":")[1]
		for addr := range s.peers {
			port := strings.Split(addr, ":")[1]
			if port < lowest {
				lowest = port
			}
		}

		leader := s.peers[s.peerAddresses[s.peers["localhost:"+lowest]]]
		response, err := leader.PlaceBid(ctx, req)
		if err != nil {
			log.Println("failed to send bid to leader")
			return &pb.BidResponse{
				Success: false,
				Error:   "failed to send bid to leader",
			}, nil
		}

		log.Printf("Sent bid %d to leader", newBid.Amount)
		return response, nil
	}

}

func (s *AuctionServer) GetResult(ctx context.Context, req *pb.ResultRequest) (*pb.ResultResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	isEnded := s.auctionEnded || time.Since(s.auctionStart) > AuctionDuration

	log.Printf("Responded with %d %s %s", s.currentBid.Amount, s.currentBid.BidderID, strconv.FormatBool(isEnded))

	return &pb.ResultResponse{
		HighestBid: s.currentBid.Amount,
		Winner:     s.currentBid.BidderID,
		IsEnded:    isEnded,
	}, nil
}

func (s *AuctionServer) ReplicateBid(ctx context.Context, req *pb.ReplicationRequest) (*pb.ReplicationResponse, error) {
	s.mu2.Lock()
	defer s.mu2.Unlock()

	if !s.isLeader {
		s.currentBid = Bid{
			Amount:    req.Amount,
			BidderID:  req.BidderId,
			Timestamp: time.Unix(req.Timestamp, 0),
		}
		s.bidHistory = append(s.bidHistory, s.currentBid)
		log.Printf("Recived Replicated bid %d from user %s", s.currentBid.Amount, s.currentBid.BidderID)
	}

	return &pb.ReplicationResponse{Success: true}, nil
}

func (s *AuctionServer) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastHeartbeat[req.Addr] = time.Unix(req.Timestamp, 0)

	return &pb.HeartbeatResponse{Acknowledged: true}, nil
}

func (s *AuctionServer) replicateBidToFollowers(ctx context.Context, bid Bid) error {
	req := &pb.ReplicationRequest{
		Amount:    bid.Amount,
		BidderId:  bid.BidderID,
		Timestamp: bid.Timestamp.Unix(),
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(s.peers))

	for _, peer := range s.peers {
		wg.Add(1)
		go func(p pb.AuctionServiceClient) {
			defer wg.Done()
			_, err := p.ReplicateBid(ctx, req)
			if err != nil {
				errors <- err
			}
		}(peer)
	}
	wg.Wait()
	close(errors)

	// Return first error if any
	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

func (s *AuctionServer) startHeartbeat() {
	ticker := time.NewTicker(HeartbeatInterval)
	go func() {
		for range ticker.C {
			s.sendHeartbeats()
		}
	}()
}

func (s *AuctionServer) sendHeartbeats() {
	ctx := context.Background()
	req := &pb.HeartbeatRequest{
		Addr:      s.address,
		IsLeader:  s.isLeader,
		Timestamp: time.Now().Unix(),
	}

	for _, peer := range s.peers {
		_, err := peer.Heartbeat(ctx, req)
		if err != nil {
			log.Printf("Failed to send heartbeat: %v", err)
		}
	}
}

func (s *AuctionServer) detectFailures() {
	ticker := time.NewTicker(HeartbeatInterval)
	for range ticker.C {
		s.mu.Lock()
		now := time.Now()

		for addr, lastBeat := range s.lastHeartbeat {
			if now.Sub(lastBeat) > NodeTimeout {
				log.Printf("Node %s appears to have failed", addr)
				delete(s.lastHeartbeat, addr)
				delete(s.peerAddresses, s.peers[addr])
				delete(s.peers, addr)

				if !s.isLeader {
					s.initiateLeaderElection()
				}
			}
		}
		s.mu.Unlock()
	}
}

func (s *AuctionServer) initiateLeaderElection() {
	// Simple leader election: node with lowest ID becomes leader
	thisPort := strings.Split(s.address, ":")[1]
	lowestPort := thisPort
	for addr := range s.peers {
		port := strings.Split(addr, ":")[1]
		if port < lowestPort {
			lowestPort = port
		}
	}

	if lowestPort == thisPort {
		s.isLeader = true
		log.Printf("Node %s became the new leader", s.address)
	}
}

func (s *AuctionServer) connectToPeers(peerAddresses []string) error {
	for _, addr := range peerAddresses {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return fmt.Errorf("failed to connect to peer %s: %v", addr, err)
		}

		client := pb.NewAuctionServiceClient(conn)
		s.peers[addr] = client
		s.peerAddresses[client] = addr
		log.Printf("Connected to peer %s", addr)
	}
	return nil
}

func main() {
	addr := flag.String("addr", "", "Node address (host:port)")
	peers := flag.String("peers", "", "Comma-separated list of peer addresses")
	flag.Parse()

	if *addr == "" {
		log.Fatal("Address is required")
	}

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := NewAuctionServer(*addr)
	grpcServer := grpc.NewServer()
	pb.RegisterAuctionServiceServer(grpcServer, server)

	// Start server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Connect to peers if specified
	if *peers != "" {
		for _, peer := range strings.Split(*peers, ",") {
			if err := server.connectToPeers([]string{peer}); err != nil {
				log.Printf("Failed to connect to peer %s: %v", peer, err)
			}
		}
	}

	// Start heartbeat and failure detection
	server.initiateLeaderElection()
	server.startHeartbeat()
	go server.detectFailures()

	// Keep the main thread alive
	select {}
}
