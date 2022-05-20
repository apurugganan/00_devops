
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

