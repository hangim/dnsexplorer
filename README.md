# dnsexplorer
Find out available DNS server.


First, you should get a DNS server list from below website.

> https://github.com/jedisct1/dnscrypt-proxy/blob/master/dnscrypt-resolvers.csv

Another optional website.

> https://dns.d0wn.biz/


You can execute this command to extract ipv4 address to get sorted and unique result.

> grep -o '[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}' dnscrypt-resolvers.csv | sort -nu > server-list.txt


Then execute this command to get available DNS server.

> go run dnsexplorer.go < server-list.txt 

