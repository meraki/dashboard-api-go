API_VERSION := v1.34.0
GENERATOR_JAR := .openapi/generator/openapi-generator-cli.jar
SPEC_FILE := $(API_VERSION).zip
SPEC_PATH := spec3.json

.PHONY: install-generator download-spec generate-code cleanup all

install-generator:
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.1/openapi-generator-cli-6.2.1.jar -O $(GENERATOR_JAR)

download-spec:
	wget https://github.com/meraki/openapi/archive/refs/tags/$(API_VERSION).zip && unzip -j $(SPEC_FILE) '*/$(SPEC_PATH)'

generate-code:
	java -jar $(GENERATOR_JAR) generate \
		-i $(SPEC_PATH) \
		-g go \
		-o client \
		-p enumClassPrefix=true \
		-p structPrefix=true \
		-p packageVersion=$(API_VERSION) \
		-t .openapi/templates \
		--package-name client \
		--git-user-id meraki \
		--git-repo-id dashboard-api-go \
		--git-host github.com

cleanup:
	rm $(SPEC_FILE); rm $(SPEC_PATH); rm $(GENERATOR_JAR)

all: install-generator download-spec generate-code cleanup
