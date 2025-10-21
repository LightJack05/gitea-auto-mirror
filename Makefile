CONTAINER_RUNTIME=podman

default: 
	$(CONTAINER_RUNTIME) compose build
down:
	$(CONTAINER_RUNTIME) compose down

up:
	$(CONTAINER_RUNTIME) compose up -d

restart:
	$(CONTAINER_RUNTIME) compose down
	$(CONTAINER_RUNTIME) compose build
	$(CONTAINER_RUNTIME) compose up -d

reset:
	@echo "Resetting dev environment..."
	$(CONTAINER_RUNTIME) compose down
	sudo cp -r data-clean/* data/
	$(CONTAINER_RUNTIME) compose build
	$(CONTAINER_RUNTIME) compose up -d
	@echo "Environment has been reset!"
