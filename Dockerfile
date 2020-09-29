FROM scratch
COPY grproxy /
ENTRYPOINT ["/grproxy"]
