#!/bin/sh
IP=192.168.2.5
LOGIN=root
PASSWD=1

PROG=test_exec_fpga



GOARCH=arm GOARM=5 GOOS=linux go build src/$PROG.go

if [ $? -eq 0 ]; then
    echo OK
    echo $IP
    ncftpput -u $LOGIN -p $PASSWD $IP /root/ ./$PROG
    echo chmod +x /root/$PROG >./arm_distro/run_test_mmap.sh
    echo /root/$PROG >>./arm_distro/run_test_mmap.sh
    ncftpput -u $LOGIN -p $PASSWD $IP /root/ ./arm_distro/run_test_mmap.sh
else
    echo FAIL
fi
