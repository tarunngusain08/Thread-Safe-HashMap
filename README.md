# Thread-Safe HashMap

This repository contains implementations of concurrent hashmaps in Go, utilizing `sync.RWMutex` and Go concurrency primitives such as goroutines, waitgroups to simulate concurrent usage of a distributed hashmap. Additionally, one of the implementations utilizes the concept of generics to create a generic hashmap using comparable interface as the key.

## Repository Structure

- **`concurrent_map.go`**: Contains the implementation of a concurrent hashmap using `sync.RWMutex`.
- **`generic_concurrent_map.go`**: Implements a generic hashmap using comparable interface as the key and `sync.RWMutex`.
- **`main.go`**: Entry point of the application. It includes test functions for both concurrent map implementations.
- **`README.md`**: Overview and documentation of the repository.

## Usage

To test the concurrent hashmap implementations, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/tarunngusain08/Thread-Safe-HashMap.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Thread-Safe-HashMap
    ```

3. Run the `main.go` file:

    ```bash
    go run main.go
    ```

## Implementation Details

- **`concurrent_map.go`**: This file contains the implementation of a concurrent hashmap using `sync.RWMutex` for synchronization. It provides thread-safe operations for setting, getting, and deleting key-value pairs in the hashmap.

- **`generic_concurrent_map.go`**: Implements a generic hashmap using comparable interface as the key. It allows for the storage of key-value pairs where the key implements the `comparable` interface. This implementation ensures type safety and thread safety.

## Contributing

Contributions to enhance the functionality, performance, or documentation of this repository are welcome. Feel free to open issues for bug reports or feature requests, and submit pull requests with improvements.

---
Feel free to reach out with any questions, feedback, or suggestions!

## Screenshots

<img width="812" alt="Screenshot 2024-02-01 at 9 39 08 AM" src="https://github.com/tarunngusain08/Thread-Safe-HashMap/assets/36428256/915f989c-1fe5-46df-9e77-afe2598ce92d">

<img width="891" alt="Screenshot 2024-02-01 at 9 38 51 AM" src="https://github.com/tarunngusain08/Thread-Safe-HashMap/assets/36428256/057aa93d-c6fc-49d9-a16c-86c8d8a92ad0">

<img width="1061" alt="Screenshot 2024-02-01 at 9 38 34 AM" src="https://github.com/tarunngusain08/Thread-Safe-HashMap/assets/36428256/af665499-d229-4fa4-8f9f-8e53ee4fb2a7">

