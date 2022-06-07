
## IAM 

## Cloud Watch

## EC2: Elastic Compute Cloud
- cloud computer / someone else's computer
- computing power, ram, storage

### Instances:
pricing models
1. on demand 
2. reserved
3. spot 
4. dedicated 

### Instance store EC2 vs EBS
instance store ec2 - temporary block store, 
- delete instance, delete data

object storage 
- includes meta data (data of data)


EBS - Elastic block storage
- ssd; non ephemeral block storage
- excellent back up functionality via snapshots
- multiple ebs to ec2; not reverse
- same availability zone for multi ebs
- pay what you use
- delta (difference)of snapshots 

AMI - Amazon Machine Umage
- read only filesystem image; 
- includes os and dependencies

## Ports
- a way / channel to communicate to the system/network
- one app per port


Security Groups - Ports
- restricting what can connect; security rules

CIDR Range 
- classless inter-domain routing; 
- method for allocating IP address and for IP routing

## SSH - Secure shell
- connect to a remote linux machine
- ssh key pairs
- must use proper OS user aka:
```
Amazon Linux => ec2.user
Redhat => root user
Ubuntu => Ubuntu user
```

## Elastic IPs
- dedicated Ip address to keep; unique address;
- attach, detach instances
- a dollar a month


Demo
login sytnax in shell
ssh -i filename.pem ec2-user@<--Public IP address-->

```
Warning: Identity file kinton_secret.pem not accessible: No such file or directory.
The authenticity of host '34.221.59.188 (34.221.59.188)' can't be established.
ED25519 key fingerprint is SHA256:xTe/fBuNdTsiKvekw7aF12z8Gi1fHOn25/qjqUkA1Gg.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? 
```

will receive an error
```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0644 for 'kinton_secret_220520.pem' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
Load key "kinton_secret_220520.pem": bad permissions
ec2-user@34.221.59.188: Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
```
on command line
```
chmod 0400 filename.pem 
```

once logged in
```
sudo yum update
sudo amazon-linux-extras install nginx1
sudo service nginx start
```

## Create EC2 (check camp cloud 4 of 7)
EC2 > Instance > Launch Instance
1. Add Name
2. Select OS
3. Create Key Pair
4. Download keypair .pem file and secure file
5. Launch
6. To prevent error using key:
```
chmod 0400 filename.pem 
```

Use / Login
on terminal
```
ssh -i name-of-file.pem ec2-user@<--public upv4 address-->
# ec2-user depends on OS
# -i tells to use identity file
# public ipv4 address data is on instance
```
## EC Recovery
### Option A
EC2 Serial Console (must be set up)

### OPTION B
Systems Manager (must be set up)

### Option C
Enter user-data
1. Create a new keypair
2. EC2 > Network and Security 
3. Press Create Key-Pair
4. Download key-par
5. On terminal
```
ssh-keygen -y -f ./new-key.pem
```
6. Stop instance > Go to instance state (new ip address)
7. Actions> Instance Settings > Edit User data
```
Content-Type: multipart/mixed; boundary="//"
MIME-Version: 1.0

--//
Content-Type: text/cloud-config; charset="us-ascii"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit
Content-Disposition: attachment; filename="cloud-config.txt"

#cloud-config
cloud_final_modules:
- [users-groups, once]
users:
  - name: ec2-user
    ssh-authorized-keys: 
    - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDRR/+dKPa56x3a3iJBkhowOE3FyI5xwcUPaVDwu3Gt3FsWveili2XsSIXicgT9Hx4s/gT1X07bqfEPa99mpZQvdDi/7kSjum2uw0VS/C3rzZ/xXO0yUWTQACQxpCpgz3b8X5DH6gpDBLrAkbDuo8bNFhV+g55iAlByOyDU+BdWoO4+uB+RSY80hd9LJ22wm/pDjy/ww2Vll5Tels4OM+c16kyTnEqiG/VGLyMiI1E2a3BQRwUmRlgM0pf3GNmSUkHxKk34JSPXr8NAMQZHtKCF7vqk+8kQgEj1NQyPk0zQ+r80AxT8gm1V8ayIUmY74ECcqVCwBee9tGDpmcuzCQav <--keypairname-->
```
8. Save
9. Restart EC2
10 On terminal
```
ssh -i newKey.pem user@newIpAddress
```

## Install Web server
1. Launch instance
2. Select OS
3. Select Key Pair
4. On Network Settings select Allow http/https
Firewall configuration
- port 443 - https
5. Get Ip address
6. Log in via SSH
7. On terminal:
resource: https://www.linuxshelltips.com/install-nginx-in-linux/
```
sudo yum update
sudo apt update

// tell package manager to look another place to look for packages
sudo amazon-linux-extras install epel

// do update again
sudo yum update

// install nginx / apache etc
sudo yum install nginx

// start server
sudo systemctl start nginx
sudo systemctl status nginx
```

Side Note
```
vim /etc/nginx/nginx.conf
vim /usr/share/nginx/html

sudo !! // last command with sudo
sudo vim /usr/share/nginx/html

download file : https://classroom.google.com/c/NTEwNDkzMzU3MDcx/m/NTMxNzczODk4NDg0/details

sudo mkdir -p /var/www/html
sudo wget https://www.quackit.com/html/templates/download/bootstrap/business-2.zip 
sudo unzip business-2.zip
ls -l business-2/

nano /etc/nginx/nginx.conf // edit root path of business-2
```

restart service
```
sudo systemctl restart nginx
```


## Run as root user all the time
```
sudo su -
```

## Add EBS to EC2 instance
1. Must be in the same availability zone
2. Go to volumes