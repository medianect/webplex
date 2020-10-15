# Copyright © 2018 Xvdfe <xvdfe@medianect.com>
# SPDX-License-Identifier: BSD-2-Clause

STATIC_FILES := $(wildcard src/static_files/*.css) \
		$(wildcard src/static_files/*.html) \
		$(wildcard src/static_files/*.js)

SOURCES :=	src/audio_file.go \
		src/directory.go \
		src/directory_file.go \
		src/file_types.go \
		src/image_file.go \
		src/main.go \
		src/static_files.go \
		src/static_handlers.go \
		src/text_file.go \
		src/unknown_file.go \
		src/video_file.go \
		src/views.go

all: src/static_files.go src/static_handlers.go webplex

src/static_files.go: ${STATIC_FILES} Makefile
	printf "package main\n" > "$@"
	@(for sf in ${STATIC_FILES}; do \
		varname=$$(printf "$$(basename "$$sf")"|sed "s/\./_/g"); \
		printf "\nvar %s = " "$$varname";\
		xxdd -f go -i "$$sf"; \
		printf "\n"; \
	done) >> "$@"

src/static_handlers.go: ${STATIC_FILES} Makefile
	printf "package main\n\n" > "$@"
	@printf "import (\n"       >> "$@"
	@printf "\t\"net/http\"\n" >> "$@"
	@printf ")\n"              >> "$@"
	@printf "\n"               >> "$@"
	@printf "func init() {\n"  >> "$@"
	@(for sf in ${STATIC_FILES}; do \
		filename="$$(basename -- "$$sf")"; \
		varname=$$(printf "%s" "$$filename"|sed "s/\./_/g"); \
		ext="$${filename##*.}"; \
		[ "$$ext" = "html" ] && continue; \
		case "$$ext" in \
			css) content_type=text/css;; \
			js)  content_type=text/javascript;; \
		esac; \
		printf "\n"; \
		printf "\thttp.HandleFunc(\n"; \
		printf "\t\t\"/%s/%s\",\n" "$$ext" "$$filename"; \
		printf "\t\tfunc(w http.ResponseWriter, r *http.Request) {\n"; \
		printf "\t\t\tw.Header().Set(\"Content-Type\", \"%s\")\n" "$$content_type"; \
		printf "\t\t\tw.Write([]byte(%s))\n" "$$varname"; \
		printf "\n\t\t},\n"; \
		printf "\n\t)\n"; \
	done) >> "$@"
	@printf "}\n" >> "$@"

node_modules/:
	npm install

webplex: ${SOURCES} node_modules/
	go build -o webplex ${SOURCES}
