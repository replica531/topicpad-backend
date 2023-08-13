up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
build:
	docker-compose build
restart:
	@make down
	@make up
ps:
	docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Networks}}\t{{.Mounts}}"
logs:
	docker-compose logs -f
mysql:
	docker-compose exec mysql bash
	# mysql -h 127.0.0.1 -P 3306 -u root -p
api:
	docker-compose exec api bash
mysql-login:
	docker-compose exec mysql bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE' && use topicpad_database
	# show tables;
