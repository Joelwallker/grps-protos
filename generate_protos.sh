#!/bin/bash

# Добавляем директорию $GOPATH/bin в PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Проверяем наличие protoc-gen-go и protoc-gen-go-grpc
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "Ошибка: protoc-gen-go или protoc-gen-go-grpc не найдены. Устанавливаем..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    echo "Установка завершена."
fi

PROTO_DIR="./"
OUT_DIR="./gen"

mkdir -p "$OUT_DIR"

for dir in $(find "$PROTO_DIR" -type d); do
    if ls "$dir"/*.proto 1> /dev/null 2>&1; then
        echo "Генерация gRPC файлов для $dir"
        protoc -I="$PROTO_DIR" \
               --go_out="$OUT_DIR" \
               --go-grpc_out="$OUT_DIR" \
               "$dir"/*.proto
    fi
done

echo "Генерация завершена. Файлы находятся в $OUT_DIR."