build:
	docker compose build 

up: 
	docker compose up

tailwind:
	npm run dev

to: 
	docker compose run --service-ports web bash


.PHONY: build up to