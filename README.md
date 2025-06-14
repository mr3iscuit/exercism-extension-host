# Exercism Extension Host

A native messaging host for browser extensions that enables communication between Chrome/Firefox extensions and the Exercism CLI.

## Features

- Native messaging host for Chrome/Firefox extensions
- Simple message routing system
- Easy installation process
- Command-line interface with help support

## Prerequisites

- Go 1.22 or later
- Chrome or Firefox browser
- Make (optional, for using Makefile commands)

## Installation

### Building from Source

1. Clone the repository:
   ```bash
   git clone git@github.com:mr3iscuit/exercism-extension-host.git
   cd exercism-extension-host
   ```

2. Build the application:
   ```bash
   make build
   ```

3. Install the Chrome native messaging host:
   ```bash
   make install
   ```

### Using Makefile

The project includes a Makefile with common commands:

```bash
make build    # Build the application
make run      # Build and run the application
make install  # Build and install the Chrome native messaging host
make clean    # Clean build files
make help     # Show help message
```

## Usage

### Running the Host

To run the native messaging host:

```bash
./extension-host
```

### Installing the Host

To install the Chrome native messaging host:

```bash
./extension-host install
```

This will create a manifest file at `~/.config/google-chrome/NativeMessagingHosts/com.google.chrome.example.echo.json`.

### Command Line Options

```bash
./extension-host --help
```

Available commands:
- `install`: Install the Chrome native messaging host
- `--help, -h`: Show help message

## Message Types

The host currently supports the following message types:

1. `text`: Echo back the received text
2. `ide`: Placeholder for IDE integration

## Development

### Project Structure

```
.
├── cmd/
│   ├── root.go      # Main command implementation
│   └── install.go   # Install command implementation
├── router/
│   └── router.go    # Message routing logic
├── types/
│   └── message.go   # Message type definitions
├── main.go          # Application entry point
├── Makefile         # Build and development commands
└── README.md        # This file
```

### Adding New Message Types

To add a new message type:

1. Add a new handler in `cmd/root.go`:
   ```go
   r.On("new-type", func(data *native.H) (*native.H, error) {
       // Handle the message
       return &native.H{}, nil
   })
   ```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 