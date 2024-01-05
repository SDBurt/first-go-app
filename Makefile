run:
	@npx tailwindcss -i ./assets/src/tailwind.css -o ./static/css/tailwind.css --minify
	@templ generate
	@go build -o tmp/main cmd/main.go