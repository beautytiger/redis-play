#!/usr/bin/env bash

sysctl vm.overcommit_memory=1
# sysctl net.core.somaxconn=4096
echo never > /sys/kernel/mm/transparent_hugepage/enabled
