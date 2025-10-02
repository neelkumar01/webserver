<div align="center">

  ### Go Web Server
  <img width="250" height="250" alt="image" src="https://cdn.jsdelivr.net/gh/neelkumar01/webserver@main/myserver/assets/image.jpeg" />
</div>
<br>
<p align="center">
<b>This project is a minimal yet fully functional web server made from scratch in go programming language and this web server is built on top of golang's built-in "net/http" package</b></p>

### ⭐️ Features

- Dynamic route handling (`/page1`, `/page2`)
- Static file serving (e.g., HTML, CSS, JS) from the `public/` directory
- Dynamic API proxying using a public API ([catfact.ninja](https://catfact.ninja))
- HTTPS support with self-signed certificates (`cert.pem` and `key.pem`) that enables TLS encryption
- Web server works with http1.1 protocol

### 🧩 Quick Start

### Prerequisites

- Go 1.25.1+ installed
- (Optional) run following command in terminal for generating local TLS certificates
```bash
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes

```

### 📦 Clone and Run

```bash
git clone https://github.com/neelkumar01/webserver.git
cd webserver/myserver
go run main.go
```

### 📁 Project Structure

```bash
webserver/
└── myserver/
    ├── public/                 # Static files directory
    │   ├── about/              # Served at /about/
    │   │   ├── index.html
    │   │   └── style.css
    │   └── root/               # Served at /
    │       ├── index.html
    │       └── style.css
    ├── cert.pem                # TLS certificate (for HTTPS)
    ├── key.pem                 # TLS private key
    ├── go.mod                  # Go module file
    └── main.go                 # Main application entry point
```

The server listens on **port 9999** for https
```bash
https://localhost:9999/

```

### 🗳️ Dependency Management

No external dependencies are used. All logic is built using Go's standard library:
- net/http
- fmt
- io
- log
