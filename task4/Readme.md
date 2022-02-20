# Linux machine

Для windows ``` GOOS=windows go build .```<br>
Для plan9 ``` GOOS=plan9 go build. ```<br>
Для linux ``` go build .```

# Windows machine
Для linux ``` $ENV:GOOS="linux" ``` потом ``` go build .```<br>
Для plan9 ``` $ENV:GOOS="plan9" ``` потом ``` go build .```<br>
Для windows ``` go build .```