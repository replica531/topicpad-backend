up:
		docker-compose up -d
down:
		docker-compose down --remove-orphans
build:
		docker-compose build
ps:
		docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Networks}}\t{{.Mounts}}"
