# Makefile to run the Go program with Air and watch TailwindCSS

# Target for starting the Go application with Air for live reloading
run:
	@echo "Starting the Go application with Air..."
	air

# Target for watching TailwindCSS files and building the CSS
watch:
	@echo "Watching TailwindCSS files for changes..."
	npx tailwindcss -i ./web/static/css/input.css -o ./web/static/css/tailwind.css --watch

# Default target to do both: run the Go application and watch TailwindCSS
all: run watch-tailwind

