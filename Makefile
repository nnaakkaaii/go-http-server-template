.PHONE: fix-lint
fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/nnaakkaaii/go-http-server-template"

.PHONY: gen-api
gen-api:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
	mkdir -p ./gen/api/
	oapi-codegen --config config/oapi-codegen/server.yaml ./spec/openapi.yaml

.PHONY: __init-db-args
__init-db-args:
ifndef DB_HOST
	$(warning DB_HOST was not set; 127.0.0.1 is used)
	$(eval DB_HOST := 127.0.0.1)
endif
ifndef DB_PORT
	$(warning DB_PORT was not set; 33066 is used)
	$(eval DB_PORT := 33066)
endif
ifndef DB_USER
	$(warning DB_USER was not set; http_server is used)
	$(eval DB_USER := http_server)
endif
ifndef DB_PASS
	$(warning DB_PASS was not set; passw0rd is used)
	$(eval DB_PASS := passw0rd)
endif
ifndef DB_NAME
	$(warning DB_NAME was not set; http_server is used)
	$(eval DB_NAME := http_server)
endif
ifndef TEST_DB_NAME
	$(warning TEST_DB_NAME was not set; test_http_server is used)
	$(eval TEST_DB_NAME := test_http_server)
endif

.PHONY: db-migrate
db-migrate: __init-db-args
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	migrate -source "file://ddl" -database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up
	$(MAKE) gen-xo

.PHONY: gen-xo
gen-xo: __init-db-args
	go install github.com/xo/xo@42b11c7999bc6ac5be620949723f44bd0ec63e02
	xo schema --out "gen/dbschema" -t json "mysql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)"

.PHONY: gen-db
gen-db:
	go run ./script/dbgen ./gen/dbschema/xo.xo.json

.PHONY: fix-db
fix-db: __init-db-args
	echo "DELETE FROM schema_migrations;" | mysql -u$(DB_USER) -p$(DB_PASS) -P$(DB_PORT) -D$(DB_NAME) -h$(DB_HOST)
	for file in $$(find ./ddl -type f -name '*.down.sql'); do mysql -u$(DB_USER) -p$(DB_PASS) -h$(DB_HOST) -P$(DB_PORT) -D$(DB_NAME) < $$file; done
