module github.com/Qitmeer/qitmeer-seeder

go 1.12

require (
	github.com/Qitmeer/qitmeer v0.0.0-20190910090212-1b7cecf2a95a
	github.com/Qitmeer/qitmeer-lib v0.0.0-20190910085745-2d3d9b8d3e06
	github.com/jessevdk/go-flags v1.4.0
	github.com/miekg/dns v1.1.15
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190312203227-4b39c73a6495
	golang.org/x/image => github.com/golang/image v0.0.0-20190227222117-0694c2d4d067
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190312151609-d3739f865fa6
	golang.org/x/net => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190511041617-99f201b6807e
	gonum.org/v1/gonum => github.com/gonum/gonum v0.0.0-20190608115022-c5f01565d866
	gonum.org/v1/netlib => github.com/gonum/netlib v0.0.0-20190313105609-8cb42192e0e0
	gopkg.in/check.v1 => github.com/go-check/check v0.0.0-20161208181325-20d25e280405
	gopkg.in/fsnotify.v1 => github.com/fsnotify/fsnotify v1.4.7+incompatible
	gopkg.in/tomb.v1 => github.com/go-tomb/tomb v1.0.0-20141024135613-dd632973f1e7
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20180328195020-5420a8b6744d
)
