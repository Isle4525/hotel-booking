.PHONY: up down logs restart clean

# Запустить PostgreSQL в фоновом режиме
up:
	docker compose up -d

# Остановить контейнеры
down:
	docker compose down

# Посмотреть логи базы данных
logs:
	docker compose logs -f postgres

# Перезапустить базу данных
restart: down up

# Полная очистка (удаление контейнеров и ВСЕХ данных базы)
clean:
	docker compose down -v

