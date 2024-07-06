package utilities

import (
	"flag"
	"os"
)

func GetPortFromFlagsOrEnv() string {
	// Define flags to accept command-line arguments
	portPtr := flag.String("p", "", "Port number for gRPC server (overrides PORT environment variable)")

	// Parse command-line arguments
	flag.Parse()

	// Check if a port was provided as a command-line argument
	var port string
	if *portPtr != "" {
		port = *portPtr
	} else {
		// Fetch the port number from an environment variable
		port = os.Getenv("PORT")
		if port == "" {
			port = "50051" // Default port if neither command-line argument nor PORT environment variable is set
		}
	}

	return port
}
