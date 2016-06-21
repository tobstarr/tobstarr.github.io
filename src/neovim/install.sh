#!/bin/bash
set -xe

## Ubuntu
apt-get install software-properties-common
add-apt-repository ppa:neovim-ppa/unstable -y
apt-get update -y
apt-get install neovim python-dev python-pip python3-dev python3-pip -y
pip install -y neovim

## Fedora
dnf -y install dnf-plugins-core
dnf -y copr enable dperson/neovim
dnf -y install neovim
