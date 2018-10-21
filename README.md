```
Grab an uclibc toolchain to avoid needing a new kernel

libpcap:
./configure CC=/home/bluecmd/Downloads/armv5-eabi--uclibc--stable/bin/arm-linux-cc LD=/home/bluecmd/Downloads/armv5-eabi--uclibc--stable/bin/arm-linux-ld --build=arm-linux-gnueabi --host=x86_64-linux-gnu
make -j8

go:
CGO_LDFLAGS="-static -L/home/bluecmd/Downloads/libpcap-1.9.0/" CGO_CFLAGS="-I/home/bluecmd/Downloads/libpcap-1.9.0/" CC=/home/bluecmd/Downloads/armv5-eabi--uclibc--stable/bin/arm-linux-cc CGO_ENABLED=1 GOARCH=arm go build -v -ldflags="-s"
```

