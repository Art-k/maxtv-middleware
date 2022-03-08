ssh root@159.203.47.150 'sudo service maxtv-middleware stop'
ssh root@159.203.47.150 'rm /root/maxtv-middleware/maxtv_middleware'

ssh root@maxtv.tech 'cd /root/maxtv-middleware/maxtv-middleware; git pull; /snap/bin/go build -v;'
#ssh root@maxtv.tech 'scp /root/maxtv-middleware/maxtv-middleware/maxtv-middleware root@159.203.47.150:/root/maxtv-middleware/maxtv-middleware'
scp root@maxtv.tech:/root/maxtv-middleware/maxtv-middleware/maxtv_middleware root@159.203.47.150:/root/maxtv-middleware/maxtv_middleware

ssh root@159.203.47.150 'sudo service maxtv-middleware start'
ssh root@159.203.47.150 'sudo service maxtv-middleware status'


# rm maxtv_middleware
#go build
#ssh root@159.203.47.150 'sudo service maxtv-middleware stop'
#scp maxtv_middleware root@159.203.47.150:/root/maxtv-middleware/maxtv_middleware
#ssh root@159.203.47.150 'sudo service maxtv-middleware start'
#ssh root@159.203.47.150 'sudo service maxtv-middleware status'