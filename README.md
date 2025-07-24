# gobolt-view

[![GoGin](https://img.shields.io/badge/Go-Gin-blue)](https://go.dev/)
[![Svelte](https://img.shields.io/badge/Svelte-3-orange)](https://svelte.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/username/gobolt-view/blob/main/LICENSE)
[![BoltDB](https://img.shields.io/badge/Bolt-DB-green)](https://dbdb.io/db/boltdb)
[![BoltDB](https://img.shields.io/badge/Bolt-DB-blue)](https://tailwindcss.com/)

A modern web UI for managing BoltDB databases, built with Svelte.

## Features

- Intuitive web interface for BoltDB
- Built with Svelte 3 for a fast, reactive experience
- Efficient data browsing, editing, and management
- Clean, streamlined design

## Getting Started

### Prerequisites

- [Go Lang](https://go.dev/)
- [Node.js](https://nodejs.org/) (v18 or newer)
- [npm](https://www.npmjs.com/)

### Installation

1. **Clone the repository**
  ```bash
  git clone https://github.com/username/gobolt-view.git
  cd gobolt-view
  ```

2. **Install dependencies**
  ```bash
  go install
  cd web/
  npm install
  ```

### Running the Application

You can start the development server with:

```bash
cd server/
go run main.go [flags]
```

Then you can start the web server with:

```bash
cd web/
npm run start [flags]
```

Then open your browser at [http://localhost:4000](http://localhost:4000) to use gobolt-view.

#### Using the `start-gobolt.sh` Script

For convenience, you can use the provided `start-gobolt.sh` script to launch the application. This script automates the startup process. Run:

```bash
./start-gobolt.sh
```

Make sure the script has execute permissions:
```bash
chmod +x start-gobolt.sh
```

## Dependencies

- Svelte 3

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a new branch for your feature or fix
3. Make your changes
4. Submit a pull request

Please follow the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct/).

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/username/gobolt-view/blob/main/LICENSE) for details.

## Screenshots

### Main Dashboard
![Main Dashboard](./screenshots/Screenshot%202025-07-22%20at%2000.21.39.png)

### Bucket View
![Bucket View](./screenshots/Screenshot%202025-07-22%20at%2000.22.10.png)

### Key/Value View
![Key/Value View](./screenshots/Screenshot%202025-07-24%20at%2023.28.31.png)

### Key/Value Editor
![Key/Value Editor](./screenshots/Screenshot%202025-07-24%20at%2023.28.46.png)

### Query View
![Key/Value Editor](./screenshots/Screenshot%202025-07-22%20at%2000.26.40.png)