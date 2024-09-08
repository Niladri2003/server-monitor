sudo apt-get update
sudo apt-get install libpcap-dev



go build -o myapp
./myapp --config /path/to/config.yaml



sysmos-package/
├── usr/
│   ├── local/
│   │   └── bin/
│   │       └── sysmos-app            # Binary executable
│   └── share/
│       └── doc/
│           └── sysmos/
│               ├── README.md         # Documentation
│               └── LICENSE           # License file
└── etc/
└── sysmos/
└── config.yaml               # Configuration file



fpm -s dir -t deb -n sysmos-app -v 1.0.1 \
--prefix / \
--description "SYSMOS System Monitoring Tool" \
--maintainer "Niladri Adak" \
--license "MIT" \
--url "https://sysmos.niladriadak.tech" \
usr/ etc/
