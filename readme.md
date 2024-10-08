# Rigo Client CLI

This is a simple CLI tool that connects to a Rigo queue server over TCP, sends commands, and receives responses. It allows you to interact with the Rigo server via the terminal.

## Features

- Connects to a Rigo server via TCP.
- Sends commands interactively.
- Receives and displays responses from the Rigo server.
- Gracefully closes the connection when the user exits.

## Prerequisites

- Go 1.16+ installed on your machine.
- A running instance of the Rigo server accessible over TCP.

## Installation

To install and use the Rigo Client CLI, first, clone the repository and then build the Go binary.

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/rigo-client-cli.git
   cd rigo-client-cli
   ```

2. Build the binary:

   ```bash
   go build -o rigo-client
   ```

3. The executable `rigo-client` will be created in the current directory.

## Usage

### Running the CLI

You can start the client by running the `rigo-client` command, providing the **host** and **port** of the Rigo server as arguments:

```bash
./rigo-client <host> <port>
```

### Example

```bash
./rigo-client 127.0.0.1 8080
```

This command will connect to a Rigo server running on `localhost` at port `8080`.

### Interactive Commands

Once the client is connected, you can enter commands to be sent to the Rigo server:

```bash
127.0.0.1:8080> your-command-here
```

The client will wait for a response from the server, which will be printed in the terminal:

```bash
Received response: <response-from-server>
```

The command prompt will appear again, allowing you to send more commands to the server.

### Exiting the Client

To exit the CLI and close the connection to the Rigo server, simply press **Ctrl + C** or terminate the process.

When the client disconnects, it will display the following message:

```bash
Closing the connection to Rigo server ❌
```

## Error Handling

- If the connection to the server fails, you will see an error message indicating the issue, and the client will terminate:

  ```bash
  Error: <error-details>
  ```

- If an error occurs while sending a command or reading a response, the client will display an appropriate error message and exit.

## Example Workflow

```bash
./rigo-client 127.0.0.1 8080
Connecting to Rigo server at 127.0.0.1:8080... 🛫
Connected to Rigo server at 127.0.0.1:8080 ✅
127.0.0.1:8080> CHANNEL job123
Received response: OK
127.0.0.1:8080> PRODUCE job123 something
Received response: OK
127.0.0.1:8080> CONSUME job123
Received response: something
127.0.0.1:8080> CONSUME job123
Received response: EMPTY
127.0.0.1:8080> exit
Closing the connection to Rigo server ❌
```

## Future Improvements

- Add command validation before sending to the server.
- Implement command history in the interactive shell.
- Support for secure connections (TLS/SSL).

Server Implementation at: [rigo](https://github.com/midedickson/rigo)
