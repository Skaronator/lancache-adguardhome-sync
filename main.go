package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Get sync interval from environment variable, default to 1440 minutes (24 hours)
	intervalStr := os.Getenv("SYNC_INTERVAL_MINUTES")
	if intervalStr == "" {
		intervalStr = "1440"
	}

	intervalMinutes, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.Fatalf("Invalid SYNC_INTERVAL_MINUTES value: %s", intervalStr)
	}

	if intervalMinutes <= 0 {
		log.Fatal("SYNC_INTERVAL_MINUTES must be greater than 0")
	}

	log.Printf("Starting lancache-adguardhome-sync with %d minute interval", intervalMinutes)

	// Run initial sync
	sync()

	// Set up ticker for periodic syncing
	ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
	defer ticker.Stop()

	// Keep running and sync on ticker
	for {
		select {
		case <-ticker.C:
			sync()
		}
	}
}

func sync() {
	log.Println("Running sync...")
	// TODO: Implement actual sync logic here
	// This is where you would implement the lancache to AdGuard Home sync functionality
	log.Println("Sync completed")
}