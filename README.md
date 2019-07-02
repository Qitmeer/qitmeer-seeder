# hlc-seeder

The seeder of the Qitmeer network

## Usage

modify config.go 

```zsh
~ go build
```

*Build Linux*

```zsh
./build-linux.sh
```

```zsh
./hlc-seeder --testnet -H seed.example.com -n xps.example.com -s ip
```

```zsh
dig -t NS seed.fulingjie.com
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
seed.fulingjie.com.	1	IN	NS	xps.fulingjie.com.

;; Query time: 10 msec
;; SERVER: 192.168.31.1#53(192.168.31.1)
;; WHEN: 二  7 02 20:14:01 CST 2019
;; MSG SIZE  rcvd: 54
```

```zsh
dig seed.fulingjie.com
```

```zsh
; <<>> DiG 9.14.3 <<>> seed.fulingjie.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 28250
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;seed.fulingjie.com.		IN	A

;; ANSWER SECTION:
seed.fulingjie.com.	1	IN	A	104.220.88.225

;; Query time: 1655 msec
;; SERVER: 192.168.31.1#53(192.168.31.1)
;; WHEN: 二  7 02 20:15:04 CST 2019
;; MSG SIZE  rcvd: 52
```