# what the docker?

Repro:
```
$ git clone https://github.com/voutasaurus/whatthedocker
$ docker build -t whatthedocker .
$ docker run whatthedocker
standard_init_linux.go:211: exec user process caused "no such file or directory"
```
