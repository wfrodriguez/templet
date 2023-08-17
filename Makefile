Name := Templet
Version := 0.0.1
Binary := templet
MainFile := cmd/templet/main.go
Args := --help

Green  := $(shell tput -Txterm setaf 2)
Yellow := $(shell tput -Txterm setaf 3)
White  := $(shell tput -Txterm setaf 7)
Cyan   := $(shell tput -Txterm setaf 6)
Reset  := $(shell tput -Txterm sgr0)

all: help

.PHONY: build # Compila el proyecto y genera el binario en ./bin
build:
	@echo '${Green}Compilando...${Reset}'
	@mkdir -p bin
	@go build -o bin/$(Binary) $(MainFile)

.PHONY: clean # Borra los archivos creados por el proyecto
clean:
	@echo '${Green}Limpiando archivos creados por el proyecto...${Reset}'
	@rm -fr ./bin

.PHONY: test # Corre los tests del proyecto
test:
	@echo '${Green}Cobertura del proyecto...${Reset}'
	go test -v ./...

.PHONY: coverage # Corre los tests del proyecto y genera el reporte de cobertura
coverage:
	@echo '${Green}Cobertura del proyecto...${Reset}'
	go test -cover -covermode=count -coverprofile=profile.cov ./...
	go tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	@go get -u github.com/AlekSi/gocov-xml
	@go get -u github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif

.PHONY: run/go # Inicia la aplicación
run/go:
	@echo '${Green}Iniciando aplicación...${Reset}'
	@go run $(MainFile) $(Args)

.PHONY: run/bin # Inicia el binario
run/bin:
	@echo '${Green}Iniciando aplicación...${Reset}'
	@./bin/$(Binary)

.PHONY: help # Muestra esta ayuda.
help:
	@echo ''
	@echo 'Modo de uso:'
	@echo '  ${Yellow}make${Reset} ${Green}<target>${Reset}'
	@echo ''
	@echo 'Targets:'
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/  $(Yellow)\1\t$(Cyan)\2$(Reset)/' | expand -t20
