#!/bin/bash

sed -s "s/^127.0.0.1.*/127.0.0.1 localhost $(ec2metadata --instance-id)/g" -i /etc/hosts
ec2metadata --instance-id > /etc/hostname
hostname -F /etc/hostname
