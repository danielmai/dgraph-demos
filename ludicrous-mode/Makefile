.PHONY: up down ps logs clean load data-info schema-info ui

PROJECT = dgraph

up:
	docker-compose -p $(PROJECT) up --detach

down:
	docker-compose -p $(PROJECT) down

ps:
	docker-compose -p $(PROJECT) ps

logs:
	docker-compose -p $(PROJECT) logs -f

clean:
	rm -rf ./data/*

load:
	./live.sh

data-info:
	@echo "1million.rdf.gz: $(shell zcat datasets/1million.rdf.gz | wc -l) edges"

schema-info:
	cat datasets/1million.schema

ui:
	python -m webbrowser "https://localhost:8000"
