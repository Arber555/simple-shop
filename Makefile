# Variables
DOCKER_USERNAME ?= stefanprifti
VERSION ?= v1

WEB_IMAGE = $(DOCKER_USERNAME)/simpleshop-web:$(VERSION)
PRODUCTS_API_IMAGE = $(DOCKER_USERNAME)/simpleshop-products-api:$(VERSION)
STOCK_API_IMAGE = $(DOCKER_USERNAME)/simpleshop-stock-api:$(VERSION)
PRODUCTS_DB_IMAGE = $(DOCKER_USERNAME)/simpleshop-products-db:$(VERSION)

.PHONY: all build-web build-products-api build-stock-api build-products-db push-web push-products-api push-stock-api push-products-db push-all

# Build all images
all: build-web build-products-api build-stock-api build-products-db

# Build web image
build-web:
	@echo "Building Web Image..."
	docker build -t $(WEB_IMAGE) ./web

# Build products-api image
build-products-api:
	@echo "Building Products API Image..."
	docker build -t $(PRODUCTS_API_IMAGE) ./products-api

# Build stock-api image
build-stock-api:
	@echo "Building Stock API Image..."
	docker build -t $(STOCK_API_IMAGE) ./stock-api

# Build products-db image
build-products-db:
	@echo "Building Products Database Image..."
	docker build -t $(PRODUCTS_DB_IMAGE) ./products-db

# Push all images to Docker Hub
push-all: push-web push-products-api push-stock-api push-products-db

# Push web image
push-web:
	@echo "Pushing Web Image to Docker Hub..."
	docker push $(WEB_IMAGE)

# Push products-api image
push-products-api:
	@echo "Pushing Products API Image to Docker Hub..."
	docker push $(PRODUCTS_API_IMAGE)

# Push stock-api image
push-stock-api:
	@echo "Pushing Stock API Image to Docker Hub..."
	docker push $(STOCK_API_IMAGE)

# Push products-db image
push-products-db:
	@echo "Pushing Products Database Image to Docker Hub..."
	docker push $(PRODUCTS_DB_IMAGE)
