go install github.com/discord-gophers/goapi-gen@latest
goapi-gen -p server -o server/api.gen.go wargops-api.json
echo "Generated API"