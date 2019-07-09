# hlc-seeder

The seeder of the Qitmeer network

## Usage

```zsh
git clone https://github.com/HalalChain/hlc-seeder.git
```

modify config.go 

```golang
if cfg.TestNet {
		activeNetParams = &params.TestNetParams
		activeNetParams.Name = "testnet"
		activeNetParams.Net = protocol.TestNet
		activeNetParams.DefaultPort = "18130"
		seed := "seed1.hlcseeder.xyz"
		activeNetParams.DNSSeeds = []params.DNSSeed{
			{seed, true},
			{seed, true},
			{seed, true},
		}
	}
```

```zsh
go build
```

*Build Linux*

```zsh
./build-linux.sh
```

```zsh
./hlc-seeder --testnet -H seed.example.com -n example.com -s ip
```

Build Qitmeer

```zsh
git clone https://github.com/HalalChain/qitmeer.git
```

modify config.go

```golang
if cfg.TestNet {
		numNets++
		activeNetParams = &testNetParams
		activeNetParams.Name = "testnet"
		activeNetParams.Net = protocol.TestNet
		activeNetParams.DefaultPort = "18130"
		seed := "seed1.hlcseeder.xyz"
		activeNetParams.DNSSeeds = []params.DNSSeed{
			{seed, true},
			{seed, true},
			{seed, true},
		}
	}
```

```zsh
go build
```

create start script

```zsh
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
#notls="--notls"
debuglevel="debug"
rpcmaxclients="10000000"


./qitmeer ${net} ${mining} ${debug} ${rpc} ${path} ${index} ${listen} ${rpcmaxclients} ${notls} ${debuglevel} ${rpcmaxclients} "$@"
```
Start the first node on server

```zsh
chmod 755 start-qitmeer.sh
./start-qitmeer.sh
```

*note : you can delete testnet folder and restart hlc-seeder*

Then you can see sync blocks

```zsh
dig -t NS seed.example.com
```

```zsh
; <<>> DiG 9.14.3 <<>> -t NS seed.fulingjie.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 13277
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.fulingjie.com.		IN	NS

;; ANSWER SECTION:
seed.example.com.	1	IN	NS	xps.example.com.

;; Query time: 10 msec
;; SERVER: 192.168.31.1#53(192.168.31.1)
;; WHEN: 二  7 02 20:14:01 CST 2019
;; MSG SIZE  rcvd: 54
```

```zsh
dig seed.fulingjie.com
```

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

;; Query time: 1655 msec
;; SERVER: xxx.xxx.xxx.xxx#53(xxx.xxx.xxx.xxx)
;; WHEN: 二  7 02 20:15:04 CST 2019
;; MSG SIZE  rcvd: 52
```