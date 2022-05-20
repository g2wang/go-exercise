test:
	docker-compose up --build --abort-on-container-exit
	docker-compose down --volumes
