docker run -d --restart=unless-stopped -v /var/rancher/:/var/lib/rancher/ -v /var/log/auditlog:/var/log/auditlog -p 8084:80 -p 8443:443 rancher/rancher:v2.3.6
