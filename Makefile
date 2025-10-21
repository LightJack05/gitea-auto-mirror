CONTAINER_RUNTIME=podman
USERNAME=$(USER)

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

data-update-clean: down data-fix-perms
	cp -r data/* data-clean/

data-fix-perms:
	sudo chown -R $(USERNAME) data/ data-clean/
	chmod 777 -R data/ data-clean/

data-reset:
	@echo "Resetting dev environment..."
	$(CONTAINER_RUNTIME) compose down
	rm -r data/*
	cp -r data-clean/* data/
	$(CONTAINER_RUNTIME) compose build
	$(CONTAINER_RUNTIME) compose up -d
	@echo "Environment has been reset!"
