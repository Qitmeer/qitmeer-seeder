# qitmeer-seeder

The seeder of the Qitmeer network

## Usage

```bash
git clone https://github.com/apefuu/qitmeer-seeder.git
```
### Build qitmeer-seeder

```bash
cd qitmeer-seeder
go get -u -x
go mod tidy
go build
```

### Build Linux

```bash
./build-linux.sh
```

### Start qitmeer-seeder

```bash
./qitmeer-seeder --testnet -H seed.example.com -n example.com -s nodeip
```
or modify the start-seeder.sh and
```bash
./start-seeder.sh
```

### Build Qitmeer

```bash
git clone https://github.com/HalalChain/qitmeer.git
```

```bash
cd qitmeer
go get -u -x
go mod tidy
go build
```

### Create start script

```bash
vim start-qitmeer.sh
```

```bash
#!/usr/bin/env bash

net="--testnet"
mining="--miningaddr TmRqga4jcJsKDYTZDSgQfWvQb9oK6HzVgxY"
debug="-d trace --printorigin"
rpc="--rpclisten 0.0.0.0:1234 --rpcuser test --rpcpass test"
path="-b "$(pwd)
index="--txindex"
listen="0.0.0.0:18130"
rpcmaxclients="2000"
debuglevel="debug"
rpcmaxclients="10000000"

./qitmeer ${net} ${mining} ${debug} ${rpc} ${path} ${index} ${listen} ${rpcmaxclients} ${debuglevel} ${rpcmaxclients} "$@"
```
Start the first node on server

```bash
chmod 755 start-qitmeer.sh
./start-qitmeer.sh
```

*note : you can delete testnet folder and restart qitmeer-seeder*

Then you can see sync blocks,and then you can use

```bash
dig -t NS seed.example.com
```

```zsh
; <<>> DiG 9.14.3 <<>> -t NS seed.example.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 13277
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.example.com.		IN	NS

;; ANSWER SECTION:
seed.example.com.	1	IN	NS	xps.example.com.

;; Query time: 10 msec
;; SERVER: 192.168.31.1#53(192.168.31.1)
;; WHEN: 二  7 02 20:14:01 CST 2019
;; MSG SIZE  rcvd: 54
```

```bash
dig seed.example.com
```

*Note : You should wait the blocks sync finished*

```zsh
; <<>> DiG 9.14.3 <<>> seed.example.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 28250
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.example.com.		IN	A

;; ANSWER SECTION:
seed.example.com.	1	IN	A	xxx.xxx.xxx.xxx
seed.example.com.	1	IN	A	xxx.xxx.xxx.xxx

;; Query time: 1655 msec
;; SERVER: xxx.xxx.xxx.xxx#53(xxx.xxx.xxx.xxx)
;; WHEN: 二  7 02 20:15:04 CST 2019
;; MSG SIZE  rcvd: 52
```