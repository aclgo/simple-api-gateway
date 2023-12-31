push:
    sudo docker compose up --build -d
migrate_up:
    migrate -database postgres://grpc-admin:grpc-admin@localhost:5432/grpc-admin?sslmode=disable -path migrations up 1
migrate_down:
    migrate -database postgres://grpc-admin:grpc-admin@localhost:5432/grpc-admin?sslmode=disable -path migrations down 1