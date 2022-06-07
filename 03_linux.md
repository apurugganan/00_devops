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
https://www.journaldev.com/43930/process-management-in-linux
```
top

uptime
```

SIDENOTE:
- load average;
- scales for systems loads;
- amount of process and number of cores
- 24 cores 24 load average
- trend data
1-5     5-10    10
0.00    0.00    0.00

PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                                                    
                                                                    
 2 root      20   0       0      0      0 S   0.0   0.0   0:00.04 kthreadd                                                                   
 3 root       0 -20       0      0      0 I   0.0   0.0   0:00.00 rcu_gp                                                                     
13 root     -51   0       0      0      0 S   0.0   0.0   0:00.00 idle_inject/0                                                              
14 root      20   0       0      0      0 S   0.0   0.0   0:00.00 cpuhp/0  

- PR priority; lower number highrt priority
- NI change priority

