.PHONY: up down ps logs clean load data-info schema-info ui

up:
	docker-compose up --detach

down:
	docker-compose down

ps:
	docker-compose ps

logs:
	docker-compose logs -f

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
