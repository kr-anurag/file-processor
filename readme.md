# File Processor

## Overview

This project processes text files in a specified directory, counting the number of lines and words in each file. It uses a concurrent approach to handle multiple files efficiently.

## Features

- Processes all `.txt` files in a given directory.
- Counts the number of lines and words in each file.
- Handles errors gracefully and reports them.
- Outputs the results to the console.

## Usage

### Running the Application

1. **Build the application:**

   ```sh
   go build -o fileprocessor cmd/main.go
   
2. **Run the application:**

   ```sh
   go run cmd/main.go ./files/
   ```

3. **Run the tests:**

   ```sh
   go test ./...
   ```