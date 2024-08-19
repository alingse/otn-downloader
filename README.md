# otn-downloader

一个基于单向光学传输网络的下载工具

a downloader based on the single direction optical transport network (aka data over video)

把数据转换成二维码流, 拿手机扫，但是避免使用任何传统互联网,USB 等设备，只需要一个屏幕，就能下载电脑的数据。


## Quick Start

1. 安装 cli 工具

```bash
go install github.com/alingse/otn-downloader@master
```


2. 访问 https://alingse.github.io/otn-downloader/index.html

点击开始，并给一下摄像头权限，如果有手机支架，建议使用手机支架。


3. 将数据转换为二维码流

```bash
otn-downloader encode --fps 10 --loop 10 --chunk-size 60 -f index.html
```

4. 如果某些切片缺失了，可以使用如下命令, 只输出特定切片的数据，以便补充

```bash
otn-downloader encode --fps 10 --loop 100 --chunk-size 60 -f index.html -s 100 -s 104 -s 113 -s 124
```


## 数据流向

1. data -> otn-downloader -> qrcode video -> terminal

2. html -> video -> mobile -> download -> data (use js) need https

3. otn-downloader start a server

4. <s>video -> otn-downloader extract -> data (use go), this is not todo</s>

## TODO

未来想做的事

1. 优化下载速度
2. 优化 index.html 的 js
3. 支持 Video 模式，先录制视频，再根据视频解码，可以使用 cli 也可以使用网页解码
4. 两台有摄像头和屏幕的电脑，可以通过互相扫描二维码，进行 TCP 通信。真正的 光学传输网络
