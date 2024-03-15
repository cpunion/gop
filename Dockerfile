ARG BASE_IMAGE=golang:1.22-bookworm

FROM $BASE_IMAGE AS build
WORKDIR /usr/local/src/gop
COPY . .
ENV GOPROOT=/usr/local/gop
RUN set -eux; \
	mkdir -p $GOPROOT/bin; \
	cp LICENSE *.mod *.sum *.md $GOPROOT/; \
	for PATTERN in "*.go" "*.md"; do \
		find . -mindepth 2 -name "$PATTERN" -exec cp --parents {} $GOPROOT/ \;; \
	done; \
	if [ -d .dist ]; then \
		GOARCH=$(go env GOARCH); \
		BIN_DIR_SUFFIX=linux_$GOARCH; \
		[ $GOARCH = "amd64" ] && BIN_DIR_SUFFIX=${BIN_DIR_SUFFIX}_v1; \
		[ $GOARCH = "arm" ] && BIN_DIR_SUFFIX=${BIN_DIR_SUFFIX}_$(go env GOARM | cut -d , -f 1); \
		cp .dist/gop_$BIN_DIR_SUFFIX/bin/gop .dist/gopfmt_$BIN_DIR_SUFFIX/bin/gopfmt $GOPROOT/bin/; \
	else \
		./all.bash; \
		cp bin/gop bin/gopfmt $GOPROOT/bin/; \
	fi

FROM $BASE_IMAGE
ENV GOPROOT=/usr/local/gop
COPY --from=build $GOPROOT/ $GOPROOT/
ENV PATH=$GOPROOT/bin:$PATH
WORKDIR /gop
