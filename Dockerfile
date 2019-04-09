FROM gcr.io/tetratelabs/tetrate-base:v0.1
ADD build/log /usr/local/bin/log
ENTRYPOINT ["/usr/local/bin/log"]
