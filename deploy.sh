ssh root@159.203.47.150 'sudo service maxtv-middleware stop'

ssh root@159.203.47.150 'cd /root/maxtv-middleware/code/maxtv-middleware/; git pull; go build;'
ssh root@159.203.47.150 'mv /root/maxtv-middleware/code/maxtv-middleware/maxtv-middleware /root/maxtv-middleware/maxtv-middleware'

ssh root@159.203.47.150 'sudo service maxtv-middleware start'
ssh root@159.203.47.150 'sudo service maxtv-middleware status'


# rm maxtv_middleware
#go build
#ssh root@159.203.47.150 'sudo service maxtv-middleware stop'
#scp maxtv_middleware root@159.203.47.150:/root/maxtv-middleware/maxtv_middleware
#ssh root@159.203.47.150 'sudo service maxtv-middleware start'
#ssh root@159.203.47.150 'sudo service maxtv-middleware status'