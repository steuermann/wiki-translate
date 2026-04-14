# Wikipedia Summary Fetcher

A high-performance, robust Go application for fetching article summaries from Wikipedia using the official API.

## Features

- **Robust Error Handling**: Comprehensive error checking for network requests, JSON parsing, and input validation.
- **Performance Optimized**:
  - Uses HTTP connection pooling with custom timeouts.
  - Efficient JSON parsing instead of fragile string manipulation.
  - Proper resource management (closing response bodies).
- **Safe & Secure**:
  - Input validation and URL encoding to prevent injection attacks.
  - Sets required `User-Agent` header to comply with Wikimedia usage policies.
- **Clean Code**:
  - No global variables.
  - Structured data models.
  - Clear, idiomatic Go patterns.

## Requirements

- Go 1.16 or higher
- Internet connection

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install dependencies (if using Go modules):
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o wiki-fetcher
   ```

## Usage

Run the application with a search term as an argument:

```bash
./wiki-fetcher "Golang"
```

Or run directly without building:

```bash
go run main.go "Artificial Intelligence"
```

### Example Output

```text
Fetching summary for: Golang
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
```

## How It Works

1. **Input Validation**: Checks if the search term is provided and not empty.
2. **URL Construction**: Safely encodes the search term and constructs the Wikipedia API URL.
3. **HTTP Request**: Sends a GET request with a proper `User-Agent` and timeout.
4. **JSON Parsing**: Decodes the API response into structured Go types.
5. **Output**: Prints the extracted summary or a friendly error message.

## API Endpoint

This tool uses the [Wikipedia Text Extracts API](https://www.mediawiki.org/wiki/API:Textextracts):
`https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&exintro&explaintext&titles=<title>`

## Error Handling

The application handles various error scenarios gracefully:
- Missing command-line arguments.
- Network connectivity issues.
- Invalid API responses.
- Non-existent articles.

## License

MIT License

## Contributing

Feel free to submit issues or pull requests to improve performance, add features, or fix bugs.
