# Paths to the main files 
SERVICE1 = ./bikes/main.go
SERVICE2 = ./rentees/main.go

# Run both services in the background
run:
	@echo "Starting Service 1..."
	go run $(SERVICE1) > bikes.log 2>&1 &
	@echo "Starting Service 2..."
	go run $(SERVICE2) > rentees.log 2>&1 &

# Run both in foreground (useful for debugging)
run-foreground:
	@echo "Running Service 1 in foreground..."
	go run $(SERVICE1)

# Stop all running Go services
stop:
	@echo "Killing Go services..."
	@pkill -f $(SERVICE1) || true
	@pkill -f $(SERVICE2) || true

# Build binaries
build:
	go build -o bin/service1 $(SERVICE1)
	go build -o bin/service2 $(SERVICE2)

# Clean binaries
clean:
	rm -rf bin/

# Watch for file changes and rerun
watch:
	@echo "Watching for changes..."
	reflex -r '\.go$$' -R '\.log$$' -- sh -c "make stop && make run"





# # Paths to the main files
# SERVICE1 = ./bikes/main.go
# SERVICE2 = ./rentees/main.go

# # Run both services in the background
# run:
# 	@echo "Starting Service 1..."
# 	go run $(SERVICE1) > bikes.log 2>&1 &
# 	@echo "Starting Service 2..."
# 	go run $(SERVICE2) > rentees.log 2>&1 &


# # Run both in foreground (useful for debugging)
# run-foreground:
# 	@echo "Running Service 1 in foreground..."
# 	go run $(SERVICE1)

# # Stop all running Go services
# stop:
# 	@echo "Killing Go services..."
# 	@pkill -f $(SERVICE1) || true
# 	@pkill -f $(SERVICE2) || true

# # Build binaries
# build:
# 	go build -o bin/service1 $(SERVICE1)
# 	go build -o bin/service2 $(SERVICE2)

# # Clean binaries
# clean:
# 	rm -rf bin/
