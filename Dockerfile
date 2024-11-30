# Step 1: Use the official Go image for Go 1.23 based on Alpine
FROM golang:1.23-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go source code into the container
COPY . .

# Step 4: Build the Go application
RUN go build -o main .

# Step 5: Expose the port that the app will run on
EXPOSE 8080

# Step 6: Define the default command to run the app
CMD ["./main"]

