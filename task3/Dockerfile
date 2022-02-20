FROM golang:1.16-alpine AS builder
WORKDIR /source
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -o linuxApp .
RUN CGO_ENABLED=0 GOOS=windows go build -o windowsApp.exe .
RUN CGO_ENABLED=0 GOOS=plan9 go build -o plan9App .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /source/linuxApp /usr/local/bin
ENTRYPOINT ["linuxApp"]

FROM plan9d/go-p9p:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /source/plan9App /usr/local/bin
ENTRYPOINT ["plan9App"]

FROM mcr.microsoft.com/windows/servercore:ltsc2022-KB5010354
RUN POWERSHELL mkdir /app
WORKDIR /app
COPY --from=builder /source/windowsApp.exe /app
ENTRYPOINT ["windowsApp.exe"]

# not sure about plan9 and windows