PKL_VERSION	 	:= v0.8.0
PKL_BASE_PATH 	:= "github.com/remychantenay/pkl-playground"

default: generate

.PHONY: prepare
prepare:
	brew install pkl
	go install github.com/apple/pkl-go/cmd/pkl-gen-go@$(PKL_VERSION)

.PHONY: generate
generate:
	pkl-gen-go pkl/config.pkl --base-path $(PKL_BASE_PATH)

.PHONY: yaml
yaml: generate
	pkl eval -f yaml -o file-yaml.yaml pkl/config.pkl

.PHONY: json
json: generate
	pkl eval -f json -o file-json.json pkl/config.pkl
