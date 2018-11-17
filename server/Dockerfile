FROM instrumentisto/dep:0.5 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/rdettai/test-owkin/server
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
ENV GIN_MODE release
ENV MODEL_FOLDER /model
ENV PORT 8081
VOLUME ["/model"]
EXPOSE 8081
COPY --from=builder /app ./
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./app"]