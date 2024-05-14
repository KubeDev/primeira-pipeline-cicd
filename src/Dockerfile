FROM golang:1.22.0-alpine3.19 as builder
WORKDIR /build 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app 
 
FROM alpine:3.19 as final
RUN adduser --uid 1000 --disabled-password appuser
USER appuser
WORKDIR /app
COPY --from=builder --chown=appuser:appuser /build/app .
EXPOSE 5000
CMD [ "./app" ]