#!/usr/bin/expect

set timeout 20
set name [lindex $argv  0]

spawn telnet $name

expect " # "

send "if \[ -f '/root/1.prog' \]; then \r echo Already programmed \r else \r fpga_loader /root/1.bit\rfi\r"

expect " # "

send "echo Hello World >/root/1.prog\r"

expect " # "

send "sh /root/run_test_mmap.sh\r"

expect " # "

send "exit\r"

expect "host"
