# traffic_sniffer

Sniff TCP/IP packets in 443 port

Example run:
  - go build
  - sudo ./traffic_sniffer

Docker run:
  - docker build -t traffic_sniffer .
  - docker run --net=host -e DEVICE='lo' traffic_sniffer
  
  DEVICE - device to sniff, default 'lo'

Testing:
  - go test
