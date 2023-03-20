# goWOL

Simple application, that hosts a server to wake-on-lan machines in its local network

build and push to dockerhub
```docker buildx build --platform linux/amd64,linux/arm64 --push -t serpantes/wol:latest .
