# qitmeer-seeder

The seeder of the Qitmeer 

## Usage

### 1, build qitmeer-seeder

```bash
git clone https://github.com/Qitmeer/qitmeer-seeder.git

cd qitmeer-seeder

go build
```

### 2, build and run Qitmeer

See [Qitmeer](https://github.com/Qitmeer/qitmeer)

Notice:

> --getaddrpercent=100 ,start your Qitmeer with this parameter. 

> The Qitmeer p2p port must use default port (mainnet 8130,testnet 18130,see Qitmmer help)

```bash
# start qitmeer
./qitmeer  --testnet --getaddrpercent=100
```

### 3, seeder domain

You have atleast 2 domain names

example:

> seed.example.com  # DNS type namesever(ns), to ns.examplex.com

> ns.examplex.com   # DNS type A , to your seed server ip

### 4, start Qitmeer-seeder

Notice:

> deafult dns server port 53,so you should config your server firewall and open udp port 53

### example

```bash
# start qitmeer-seeder

./qitmeer-seeder --testnet -H example.com -n ns.example.com -l 0.0.0.0:53 -s your-qitmeer-p2plisten-ip
```

 
## How to test seeder status

1. check seed.example.com's NS is ns.example.com
```bash
dig -t ns seed.example.com

# show ...
;; ANSWER SECTION:
seed.example.com.	1	IN	NS	ns.example.com.
```

2. check qitmeer good ip list

> You should wait the blocks sync finished*

```bash
dig seed.example.com 

# show ...
;; ANSWER SECTION:
seed.example.com.	1	IN	A	xxx.xxx.xxx.xxx
seed.example.com.	1	IN	A	xxx.xxx.xxx.xxx
...
```