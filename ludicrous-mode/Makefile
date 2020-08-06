.PHONY: up ps down clean load data-info schema-info ui

up:
	docker-compose up --detach

ps:
	docker-compose ps

down:
	docker-compose down

clean:
	sudo rm -rf ./data/*

load:
	./live.sh

data-info:
	@echo "1million.rdf.gz: $(shell zcat datasets/1million.rdf.gz | wc -l) edges"

schema-info:
	cat datasets/1million.schema

ui:
	python -m webbrowser "https://localhost:8000"
