# what the docker?

Repro:
```
$ git clone https://github.com/voutasaurus/whatthedocker
$ docker build -t whatthedocker .
$ docker run whatthedocker
standard_init_linux.go:211: exec user process caused "no such file or directory"
```

What's happening?

Go is trying to compile with CGO enabled, and scratch does not contain gcc.

Fix is to add CGO_ENABLED=0 to the build step:
```
diff --git a/Dockerfile b/Dockerfile
index 57837bd..50d84b0 100644
--- a/Dockerfile
+++ b/Dockerfile
@@ -1,6 +1,6 @@
 FROM golang:1.12.9 AS build
 ADD . /app
-RUN cd /app && go build -o app
+RUN cd /app && CGO_ENABLED=0 go build -o app
 
 FROM scratch
 COPY --from=build /app/app /
```
