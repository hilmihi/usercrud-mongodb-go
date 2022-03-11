<!-- ABOUT THE PROJECT -->

## About The Project

CRUD Users using MongoDB in Go.

Building the project with layered architecture, and clean code approach for the structure, with the intention of simplicity when the app is scaling up and ease of maintenance

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

This project structure is built using

- [Golang]
- [MongoDB]
- [Labstack/Echo]

<p align="right">(<a href="#top">back to top</a>)</p>

### Features

- USERS CRUD

### Folder Structure

```
├── app                             # Main.go
├── config/                         # Configuration to connect to database
├── entities/                       # Create entities for user
├── delivery/ 
    ├── common                      # Create common response
    ├── controller                  # Create controller for user
    ├── middleware                  # Create middleware
    ├── router                      # Create router
├── repository/                     # Get all required data from database for category, comment, event, participant, and user
├── util/                           # Create utilization.

```

<!-- GETTING STARTED -->

## Getting Started

To start project, just clone this repo

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/hilmihi/e-commerce-project
   ```
2. Create .env file in main directory
   ```bash
   touch .env
   ```
3. Write the following example environment
   ```
   export DB_CONNECTION_STRING='mongodb+srv://<username>:<password>@cluster0.gu5gt.mongodb.net/<database>?retryWrites=true&w=majority'
   ```
4. Run the server
   ```bash
   source .env && go run ./app/main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>
