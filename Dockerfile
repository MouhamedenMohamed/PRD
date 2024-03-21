# Use a base image with your preferred Linux distribution
FROM ubuntu:latest

# Install Go Lang
RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean

# Install MySQL client
RUN apt-get update && \
    apt-get install -y mysql-client && \
    apt-get clean

# Set the working directory
WORKDIR /app

# Copy your Pydio Cells application files into the Docker image
COPY . .

# Expose the required ports
EXPOSE 8080

# Command to start your Pydio Cells application
CMD ["./cells"]

