FROM scratch AS runner
LABEL org.opencontainers.image.source https://github.com/kodflow/fizzbuzz

WORKDIR /var/run
COPY --chmod=0777 .build/fizzbuzz /var/run/fizzbuzz
HEALTHCHECK --interval=1m --timeout=30s --retries=3 CMD curl --fail http://localhost/v1/status/healthcheck || exit 1
EXPOSE 80 443

ENTRYPOINT [ "/var/run/fizzbuzz" ]