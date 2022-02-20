FROM golang:1.16-alpine AS builder
WORKDIR /source
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -o linuxApp ./task3/.
RUN CGO_ENABLED=0 GOOS=windows go build -o windowsApp.exe ./task3/.
RUN CGO_ENABLED=0 GOOS=plan9 go build -o plan9App ./task3/.

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /source/linuxApp /usr/local/bin
COPY --from=builder /source/windowsApp.exe /usr/local/bin
COPY --from=builder /source/plan9App /usr/local/bin

ENTRYPOINT [ "linuxApp", "windowsApp.exe", "plan9App"]