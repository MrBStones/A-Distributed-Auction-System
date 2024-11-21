package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	pb "auction/stc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	servers := flag.String("servers", "localhost:8001", "Comma-separated list of server addresses")
	action := flag.String("action", "", "Action to perform (bid/result)")
	amount := flag.Int("amount", 0, "Bid amount")
	bidderID := flag.String("bidder", "", "Bidder ID")
	flag.Parse()

	serverList := strings.Split(*servers, ",")
	for server := range serverList {
		conn, err := grpc.NewClient(serverList[server], grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to open conn: %v", err)
			continue // we try the other connections too
		}
		defer conn.Close()
		client := pb.NewAuctionServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		switch *action {
		case "bid":
			if *amount <= 0 || *bidderID == "" {
				log.Fatal("Bid amount and bidder ID are required for bidding")
			}
			response1, err1 := client.PlaceBid(ctx, &pb.BidRequest{Amount: int32(*amount), BidderId: *bidderID})
			if err1 != nil {
				// log.Printf("Error: %v", err1)
				continue
			}
			log.Printf("Response: Success = %s, error = %s", strconv.FormatBool(response1.Success), response1.Error)
			return
		case "result":
			response2, err2 := client.GetResult(ctx, &pb.ResultRequest{})
			if err2 != nil {
				// log.Printf("Error: %v", err2)
				continue
			}
			log.Printf("Response: Highest_bid = %d from %s; Is_Ended = %s; error = %s;", response2.HighestBid, response2.Winner, strconv.FormatBool(response2.IsEnded), response2.Error)
			return
		default:
			log.Fatal("Invalid action. Use 'bid' or 'result'")
		}
	}

	log.Fatal("No response from any of the listed addresses")
}
