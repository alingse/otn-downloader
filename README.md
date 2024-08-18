# otn-downloader

single direction optical transport network (aka data over video)



```bash
otn-downloader encode --fps 10 --loop 100 --chunk-size 60 -f index.html -s 100 -s 104 -s 113 -s 124
```

## Data Flow

1. data -> otn-downloader -> qrcode video -> terminal

2. html -> video -> mobile -> download -> data (use js) need https

3. otn-downloader start a server

4. <s>video -> otn-downloader extract -> data (use go), this is not todo</s>
