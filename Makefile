BINARY_NAME=Nourybot.out

xd:
	go build -o ${BINARY_NAME} cmd/bot/main.go 
	./${BINARY_NAME} -env="dev"

xdprod:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} -env="prod"

jq:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} -env="dev" | jq 

jqprod:
	go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} -env="prod" | jq 
