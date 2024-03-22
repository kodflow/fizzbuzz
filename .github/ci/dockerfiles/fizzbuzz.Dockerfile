FROM --platform=$BUILDPLATFORM scratch AS runner
LABEL org.opencontainers.image.source https://github.com/kodflow/fizzbuzz

ARG BUILDPLATFORM
ENV BUILDPLATFORM=$BUILDPLATFORM

ARG BINARY_VERSION
ENV BINARY_VERSION=$BINARY_VERSION

ARG TARGETARCH
ENV TARGETARCH=$TARGETARCH

WORKDIR /var/run
ADD https://github.com/kodflow/fizzbuzz/releases/download/$BINARY_VERSION/fizzbuzz-$TARGETARCH /var/run/fizzbuzz
HEALTHCHECK --interval=1m --timeout=30s --retries=3 CMD curl --fail http://localhost/v1/status/healthcheck || exit 1
EXPOSE 80 443

ENTRYPOINT [ "/var/run/fizzbuzz" ]