# Gorthanc

## **Attention**

This project is currently in early development and is not yet considered a stable release. The API, features, and implementation details may change significantly before the first v1.0.0 release. We welcome contributions, feedback, and suggestions from the community to help shape the direction of this library.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue.svg)](https://golang.org/)

A comprehensive, type-safe Go client library for the [Orthanc](https://www.orthanc-server.com/) DICOM server REST API.

## Overview

Gorthanc provides a clean, idiomatic Go interface for interacting with Orthanc DICOM servers. It abstracts the complexity of the REST API and provides strongly-typed methods for all major operations including managing studies, series, instances, patients and system operations.

## Key Features

- **🔒 Type-Safe**: Fully typed API with comprehensive struct definitions for all DICOM resources
- **🎯 Idiomatic Go**: Clean, Go-style API design with proper error handling
- **📝 Well Documented**: Extensive documentation with examples for every method
- **🔐 Authentication**: Built-in support for HTTP Basic Authentication
- **⚡ Efficient**: Support for pagination, query parameters, and selective expansion
- **🎨 Flexible**: Customizable HTTP client with timeout and transport options
- **🧪 Production Ready**: Error handling with custom error types for better debugging

## Advantages

- **Simplified Development**: No need to manually construct HTTP requests or parse JSON responses
- **Reduced Errors**: Strong typing catches errors at compile time rather than runtime
- **Better Maintenance**: Clear separation of concerns with dedicated types package
- **Comprehensive Examples**: Ready-to-use examples for common workflows
- **DICOM Expertise Not Required**: Abstract away DICOM complexity while still providing access to all tags
- **Consistent API**: All resources follow the same patterns for CRUD operations

## Upcoming

We plan to add a full roadmp until have coverage of most of the [Orthanc REST API](https://orthanc.uclouvain.be/api/#tag/System), right now it is covering a lot of more basics endpoint!

## Installation

```bash
go get github.com/proencaj/gorthanc
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/proencaj/gorthanc/pkg/gorthanc"
)

func main() {
    // Create a client
    client, err := gorthanc.NewClient(
        "http://localhost:8042",
        gorthanc.WithBasicAuth("orthanc", "orthanc"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Get system information
    info, err := client.GetSystem()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Orthanc Version: %s\n", info.Version)

    // List all studies
    studies, err := client.GetStudies(nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d studies\n", len(studies))
}
```

## API Coverage

Will be added soon

## Examples

The `examples/` directory contains comprehensive examples:

```bash
# System information
go run examples/basic/main.go

# Studies operations
go run examples/studies/main.go

# Series operations
go run examples/series/main.go

# Instances operations
go run examples/instances/main.go
```

## Project Structure

```
gorthanc/
├── pkg/gorthanc/          # Main package
│   ├── client.go          # HTTP client implementation
│   ├── system.go          # System endpoints
│   ├── studies.go         # Studies endpoints
│   ├── series.go          # Series endpoints
│   ├── instances.go       # Instances endpoints
│   ├── patients.go        # Patients endpoints
│   ├── instances.go       # Modalities endpoints
│   ├── errors.go          # Custom error types
│   └── types/             # Type definitions
│       ├── system.go      # System types
│       ├── study.go       # Study types
│       ├── series.go      # Series types
│       └── instance.go    # Instance types
│       └── patient.go     # Patient types
│       └── modality.go    # Modality types
└── examples/              # Usage examples
    ├── basic/             # Basic system info
    ├── studies/           # Studies examples
    ├── series/            # Series examples
    └── instances/         # Instances examples
    └── patients/          # Patients examples
    └── modalities/        # Modalities examples
```

## Requirements

- Go 1.21 or higher
- Access to an Orthanc server (local or remote)

## Testing

```bash
# Run all examples
go run examples/basic/main.go
go run examples/studies/main.go
go run examples/series/main.go
go run examples/instances/main.go

# Build the project
go build ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Resources

- [Orthanc Book](https://book.orthanc-server.com/) - Official Orthanc documentation
- [Orthanc REST API](https://orthanc.uclouvain.be/api/) - Complete API reference
- [DICOM Standard](https://www.dicomstandard.org/) - DICOM specification
- [DICOM NEMA Documentation](https://dicom.nema.org/medical/dicom/current/output/html/part01.html) - DICOM Documentation from NEMA

## Acknowledgments

- The [Orthanc](https://www.orthanc-server.com/) team for creating an excellent DICOM server

## Support

If you encounter any issues or have questions:

- Open an issue on GitHub
- Check the [examples](examples/) directory for usage patterns
- Refer to the [Orthanc REST API documentation](https://orthanc.uclouvain.be/api/)

---

Made with ❤️ for the medical imaging community
