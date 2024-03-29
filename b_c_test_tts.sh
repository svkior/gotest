#!/bin/sh
#IP=192.168.3.14
IP=192.168.220.134
LOGIN=root
PASSWD=1

PROG=test_exec_fpga



GOARCH=arm GOARM=5 GOOS=linux go build src/$PROG.go

if [ $? -eq 0 ]; then
    echo OK
    echo $IP
    ncftpput -u $LOGIN -p $PASSWD $IP /root/ ./arm_distro/1.bit
    ncftpput -u $LOGIN -p $PASSWD $IP /root/ ./$PROG
    echo chmod +x /root/$PROG >./arm_distro/run_test_mmap.sh
    echo /root/$PROG >>./arm_distro/run_test_mmap.sh
    ncftpput -u $LOGIN -p $PASSWD $IP /root/ ./arm_distro/run_test_mmap.sh

    ./run_on_arm.sh $IP
else
    echo FAIL
fi
