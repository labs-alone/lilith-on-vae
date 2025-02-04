# Lilith - First Agent Built on VAE

<div align="center">
  <img src="https://github.com/labs-alone/.github/blob/main/images/vae.png?raw=true" alt="VAE Framework" width="100%" />
</div>

Lilith is the pioneering autonomous agent built on the VAE (Versatile AI Engine) platform, demonstrating the platform's capabilities while providing a foundation for next-generation AI agents.

## Features

- **Advanced Memory System**: Multi-tiered memory management with intelligent cleanup
- **Task Processing**: Priority-based task queue with concurrent execution
- **VAE Integration**: Seamless connection to VAE's visual processing capabilities
- **State Management**: Robust state tracking and persistence
- **Monitoring**: Comprehensive metrics and performance tracking
- **Security**: Built-in security features and encryption support

## Quick Start

### Prerequisites
- Go 1.21+
- VAE Platform
- Docker (optional)

### Installation

# Clone the repository
git clone https://github.com/labs-alone/lilith.git
cd lilith

# Install dependencies
go mod download

# Run Lilith
go run cmd/lilith/main.go

### Basic Usage

package main

import (
    "github.com/labs-alone/lilith/internal/agent"
    "github.com/labs-alone/lilith/pkg/logger"
)

func main() {
    // Initialize agent with default configuration
    config := agent.NewDefaultConfig()
    log := logger.New()
    
    lilith, err := agent.NewAgent(config, log)
    if err != nil {
        log.Fatal("Failed to initialize agent:", err)
    }
    
    // Start the agent
    if err := lilith.Start(); err != nil {
        log.Fatal("Failed to start agent:", err)
    }
}

## Architecture

Lilith is designed with a modular architecture:

- **Agent Core**: Main agent lifecycle and coordination
- **Processor**: Task handling and execution
- **Memory**: Multi-tiered memory management
- **State**: Agent state and persistence
- **VAE Client**: Integration with VAE platform

## Development

# Run tests
go test ./...

# Build binary
go build -o lilith cmd/lilith/main.go

# Run with custom config
./lilith -config=/path/to/config.json

## Configuration

Lilith can be configured through environment variables or a configuration file:

{
  "name": "lilith",
  "version": "1.0.0",
  "environment": "development",
  "process_interval": "100ms",
  "max_concurrent_tasks": 10,
  "memory": {
    "max_size": 1000000,
    "cleanup_interval": "1h"
  }
}

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

- Twitter: [@alone_labs](https://x.com/alone_labs)

## Acknowledgments

Special thanks to the VAE team and the Alone Labs community for making this project possible.
