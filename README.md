# qitmeer-seeder

The seeder of the Qitmeer 

## Usage

### build qitmeer-seeder

```bash
git clone https://github.com/apefuu/qitmeer-seeder.git
cd qitmeer-seeder
go build
```

### build qitmeer

see [qitmeer](https://github.com/HalalChain/qitmeer)

## config seed domain

You must have 2 domain names
```
# example
seed.example.xxx  # DNS type namesever(ns) to ns.examplex.xxx
ns.examplex.xxx   # DNS type A to your seed server ip
```

## start qitmeer

if the network peers count less than 5,you should add start parameter “--getaddrpercent=100” to you qitmeer 

The qitmeer start parameter,see [qitmeer readme](https://github.com/HalalChain/qitmeer)

The qitmeer p2p port must use default port (mainnet 830,testnet 1830,seed qitmmer help)


## start qitmeer-seeder

deafult dns server port 53,so your should config your seeder firewall and open udp port 53

```bash
# example
./qitmeer-seeder --testnet -H seed.example.xxx -n example.xxx -l 0.0.0.0:53 -s your-qitmeer-ip
```

## example: qitmeer start bash


```bash
#!/usr/bin/env bash

net="--testnet"
mining="--miningaddr TmRqga4jcJsKDYTZDSgQfWvQb9oK6HzVgxY"
debug="-d trace"
rpc="--rpclisten 0.0.0.0:1234 --rpcuser test --rpcpass test --rpcmaxclients=2000"
path="-b "$(pwd)
index="--txindex"
debuglevel="--debuglevel debug"
getaddrpercent="--getaddrpercent=100"

./qitmeer  ${net} ${mining} ${debug} ${rpc} ${path} ${index} ${debuglevel} ${getaddrpercent}"$@"

```
 
## test

### test seed domain 

```bash

# if your seed domian is seed.example.xxx
# if your ns domain is ns.example.xxx
# "dig -t ns" will list ns.exaple.xxx

dig -t ns seed.example.xxx
```

```zsh
; <<>> DiG 9.14.3 <<>> -t NS seed.example.xxx
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 13277
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.example.xxx.		IN	NS

;; ANSWER SECTION:
seed.example.xxx.	1	IN	NS	ns.example.xxx.
```

### test your seed 

```bash
# "dig" wil list all ips 
dig seed.example.xxx
```

*Note : You should wait the blocks sync finished*

```zsh
; <<>> DiG 9.14.3 <<>> seed.example.xxx
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 28250
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.example.xxx.		IN	A

;; ANSWER SECTION:
seed.example.xxx.	1	IN	A	xxx.xxx.xxx.xxx
seed.example.xxx.	1	IN	A	xxx.xxx.xxx.xxx

```