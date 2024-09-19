
# sigma to Golang gosigma

## Overview

Gosigma is a tool designed to convert code written in sigma into Golang.



## Installation

- **Clone the Repository:**

   ```bash
   git clone https://github.com/scott-mescudi/go_to_sigma.git
   ```

- **Navigate to the Project Directory:**

   ```bash
   cd go_to_sigma
   ```

- **Install Dependencies:**

   Ensure you have Go installed. If not, install it from [golang.org](https://golang.org/doc/install).

   ```bash
   go mod tidy
   ```

- **Build the Program:**

   ```bash
   go build -o gosigma main.go
   ```

## Usage

- **Run gosigma:**

   ```bash
   ./gosigma input.sigma -o output.go
   ```

   - `input.sigma`: Path to the sigma source file.
   - `-o output.go`: Path to the output Golang file.

- **Example:**

   ```bash
   ./gosigma example.sigma -o example.go
   ```
