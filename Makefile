up:
		docker-compose up -d
down:
		docker-compose down --remove-orphans
ps:
		docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Networks}}\t{{.Mounts}}"
