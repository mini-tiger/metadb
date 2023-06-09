monstache [link](https://github.com/rwynn/monstache/tree/v6.7.11)
monstache config [link](https://rwynn.github.io/monstache-site/config/#mongo-config-url)
bk 3.9.x [link](https://github.com/TencentBlueKing/bk-cmdb/blob/release-v3.9.39/docs/overview/installation.md)
## bk-monstache (忽略)
```shell
https://github.com/TencentBlueKing/bk-cmdb/tree/release-v3.10.26/src/tools/monstache
cd `$pwd`
cp * .

GO111MODULE=off 
go get github.com/rwynn/monstache/monstachemap
go get github.com/BurntSushi/toml
make # 3.10 版本增加文件
```

## start dev monstache
```shell
version 6.7.11
cd src/configcenter/src/tools/monstache
./monstache -f build/monstache/etc/config.toml
```

## docker image monstache
```shell
docker pull rwynn/monstache-builder-cache-rel6:1.0.7
docker pull rwynn/monstache-alpine:3.15.0

cd src/configcenter/src/tools/monstache
mkdir src
cd src
git clone https://github.com/rwynn/monstache.git
cd monstache/
git checkout v6.7.11
```

```shell
cd src
modify Dockerfile
RUN export GOPROXY=https://proxy.golang.com.cn,direct && go mod download

# cmd
docker build -f Dockerfile -t  harbor.dev.21vianet.com/taojun/monstache:latest .

```

## helm monstache
```shell
cd src/configcenter/src/tools/monstache/helm
 helm -n cmdbv4 install monstache .
```