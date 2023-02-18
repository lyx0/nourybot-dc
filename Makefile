BINARY_NAME=Nourybot.out

xd:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME}

jq:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} | jq
