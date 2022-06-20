# Linux

Resources 
1. Amazon re/start

## Linode
SSH
Login to Linode

On terminal, 
```
ssh root@<--ipaddress-->
```
Enter password

## Packages
- collection of files that perform a task

## Package Managers
- help install, upgrade, configure, remove

apt vs dpkg
dpkg installs .deb packages

Prep ubuntu
```
apt update // get information
apt upgrade // upgrades system
```


Option 1
```
curl -o <--filename--> <--url-->
dpkg -i <--filename-->

dpkg -r tree
dpkg -l
```


Option 2
```
apt install <-- packagename -->
apt remove tree
apt list --installed
apt list --installed | grep tree

searching /
```

Option 3
```
apt-get install <--- package name -->
```

### Search for packages
```
dpkg-query -S package-name

apt search package-name

https://www.debian.org/distrib/packages
```

Add a repository
- packages not part of official
```
sudo add-apt-repository ppa:ppaName apt update
//personal package archive

apt install golang-go
cd /etc/apt/sources.list.d

```


Check distributio
```
// UBUNTU
cat /etc/issue

// RED HAT
cat /etc/apt/redhat-release
yum // redehat
```

## Build Source 
- install from source
- manually compiling an apps source code 

going to repos, get source code, compile

OPTION 1
```
// UBUNTU 
apt install build-essential

// REDHAT
yum groupinstall "Development Tools"
```

```
// install dependencies
sudo apt-get install ninja-build gettext libtool libtool-bin autoconf automake cmake g++ pkg-config unzip curl doxygen

// clone repo
git clone https://github.com/neovim/neovim

cd to repo
cd neovim/

// build
make
```

## June 03, 2022
## 5 of 7: Managing Linux

Manage processes
```
//snapshot of processes
ps 

  PID TTY          TIME CMD
78066 pts/0    00:00:00 bash
78123 pts/0    00:00:00 ps

ps -e
ps -ef
ps a
ps auwx

ps -ef | grep -i python
```

top : view processin real time 
what status mean: 
https://www.journaldev.com/43930/process-management-in-linux
```
top

uptime
```

SIDENOTE:
- load average;
- scales for systems loads;
- amount of process and number of cores
- 24 cores 24 load average; max utilization
- trend data (mins)
```
1-5     5-10    10
0.00    0.00    0.00
```


```
PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                                                    
                                                                    
 2 root      20   0       0      0      0 S   0.0   0.0   0:00.04 kthreadd                                                                   
 3 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 rcu_gp                                                                     
13 root     -51   0       0      0      0 S   0.0   0.0   0:00.00 idle_inject/0                                                              
14 root      20   0       0      0      0 S   0.0   0.0   0:00.00 cpuhp/0  
```
- PR priority; lower number highrt priority
- NI change priority/ Nice values
to kill a PID
```
k
provide pid number
```

htop - more helpful than top; easier to read


Sigterm: https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
sigterm - shutdown gracefully

kill - stop a process
```
sudo apt install nginx
systemctl status nginx
ps auwx | grep nginx

kill <processid>
systemctl start nginx


kill -9  <pid> // kills all process in the group
```

free : checks memory in bytes
```
free

free -m
free -g
free -h

              total        used        free      shared  buff/cache   available
Mem:          1.9Gi       117Mi       376Mi       0.0Ki       1.5Gi       1.6Gi
Swap:         511Mi       1.0Mi       510Mi
```

swap - offloading from ram to disk space

```
df // file system information
du // disk usage; how much space is this folder taking up
```

Linux System Start up
run levels : modes it runs on
| Level | Mode |
| --- | --- |
| 0 | halt - shutdown |
| 1 | single user mode (God mode) physical access |
| 2 | multi user mode (safe mode) physical access |
| 3 | multi network w/ networking (normal boot; cmd only) |
| 4 | Undefined |
| 5 | X11 (desktop linux) |
| 6 | Reboot  |

```
runlevel

who -r

systemctl get-default
w
```

change level
```
init <runlevel>

init 0
init 1

reboot now
```

Permanent Runlevel Change
```
systemctl set-default <graphical.target>
systemctl set-default reboot.target //dont do this 0 and 6
```

See what's loaded
```
systemctl list-units
```

start at boot 
```
systemctl enable <service-name>
```

Manage processes
```
systemctl start nginx
systemctl stop nginx
systemctl restart nginx
```

system nginx status
```
ip addr
```

### CRON
- run more than once; schedule when to run
cron.daily
cron.monthly
cron.weekly

Cron expression
Resource: crontab.guru
/etc/crontab
```
vim /etc/crontab

* * * * *   root touch ~/`date +%s`.txt


watch date
cat /var/log/cron
```

## 6 of 7 : Basic Networking 

Internet 
- a netword of interconnected computers 
- wires

TCP/IP
- is a suite of protocols (set of rules) that allow computers to communicate over a network

IP - Internet Protocol 
- the part that obtains the address that data will be sent to.
- 192.168.1.1

TCP - Transmission Control Protocol 
- delivers the data to the IP address. 
- the process that happens when you receice data

Network Interfaces
- ifconfig utility; allow to get/change information about interfaces available
- multiple interfaces (wifi, cable)

```
ip addr
ifconfig

// interface where system can talk to itself
// 127.0.0.1 - my ip address 

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever


// wired connection -  139.xxx.xxx.xxx (1st one)

2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether f2:3c:93:ac:5e:d0 brd ff:ff:ff:ff:ff:ff
    inet 139.xxx.xxx.xxx/24 brd 139.xxx.xxx.xxx scope global eth0
```

setting interfaces up or down 
```
ip link ser eth0 up
ip link ser eth0 down
```

## Traceroute
See the path that a request takes throughtout network
resource: redhat 7 great network commands
```
mtr <url>
mtr -4 <url>
traceroute
```
Packet loss - where is it going
packet - how data is transferred

## DNS Domain Name System
- mapped to ip address
- translate address
Going to websites
1. Local DBS - Cache
2. Recursive name server - cache websites; gives ip add
3. Authoritative Name Server - tell internet where to find information; DNS record => goes to RNS

## Record
host record - google.com => 8.8.8.8; IPV4
AAAA record - google.com => fe80::1ff:fe23:4567:890a; IPV6
CNAME record; canonical name - google.com => public.02.fun.com; domain name aliases
MX records Mail exchange - google.com => aspmx.l.google.com 
NS Records; Name Server - google.com => nsl.google.com; sets server respnsible for providing dns information 

```
dig <url>
dig mx <url>

ip address 
- not all sites will allow this way to access 
- maybe usable for single server
```

## /etc/hosts
- manual set a dns override; local mapping

## Connectins and Sockets
gather information about about ports and sockets a computer is using 
```
lsat
netstat -tulpn
lsof -i -P// see listening ports 

cat /etc/services
```

## 7 of 7:  Bash Scripting 

## Scripting 
- writing code to automate a task
- often interpreted (line by line); best for scripting

``` 
#!/bin/bash // tell which interpreter which to use; first line of code

if - else if - else

&&
||
!
```

Exit Codes
- linux process return an exit code when they are finished; 
- status of the run
```
0 - Success
1 - General unkown error
2 - Misuse of shell command

126 - Shell cant execute
127 - command not found
128 - invalid exit argument
128+x - fatal error with Linux signal x
130 - command terminated with ctrl + c
255 exit status out of range