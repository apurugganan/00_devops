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
