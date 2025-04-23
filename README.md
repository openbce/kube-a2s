# Kube-a2s

This is a Go project that [describe your project purpose here].

## Requirements

- Go 1.21 or higher
- Make

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/openbce/kube-a2s.git
cd kube-a2s
```

2. Build the project:
```bash
make build
```

3. Run the application:
```bash
./bin/kube-a2s
```

## Development

The project includes a Makefile with common development tasks:

- `make build`: Build the binary
- `make test`: Run unit tests
- `make coverage`: Generate test coverage report
- `make fmt`: Format code
- `make vet`: Run go vet
- `make lint`: Run golangci-lint
- `make clean`: Clean build artifacts
- `make help`: Show all available commands

## Project Structure

```
.
├── cmd/
│   └── kube-a2s/    # Application entry point
│       └── main.go
├── pkg/             # Package directory for shared code
│   └── apis/        # API type definitions
├── bin/             # Compiled binaries
├── Makefile         # Build automation
├── go.mod          # Go module definition
└── README.md       # Project documentation
```

## License

Copyright 2024 The Kube-a2s Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
