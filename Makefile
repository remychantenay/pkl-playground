PKL_BASE_PATH 	:= "github.com/remychantenay/pkl-playground"

default: generate

.PHONY: generate
generate:
	pkl-gen-go pkl/config.pkl --base-path $(PKL_BASE_PATH)
	# Generating output files.
	pkl eval -f json -o file-json.json pkl/config.pkl
	pkl eval -f yaml -o file-yaml.yaml pkl/config.pkl

.PHONY: run
run: generate
	go run main.go