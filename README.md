# Url Shortener

- Built with Golang&HTMX

## Getting started

- You need to have Go installed on your system

  ```sh
      brew install go
  ```

- Clone repository

  ```sh
    git clone https://github.com/mscandan/url-shortener.git
  ```

- Download dependencies

  ```sh
    go mod download
  ```

- Create database on [MongoDB Atlas](https://www.mongodb.com)

- Create Redis instance on [Render](https://render.com/) or the provider of your choice

- Create and set environment variables in .env file

  ```sh
    touch .env
  ```

- Start development server

  ```sh
    go run .
  ```
