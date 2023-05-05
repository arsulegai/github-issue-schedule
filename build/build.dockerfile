FROM golang:1.20.4-bullseye

# Add make command
RUN apt update && apt install make

WORKDIR /app
