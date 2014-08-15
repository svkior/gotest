#!/usr/bin/expect

set timeout 20
set channel [open "/dev/tty.usbserial" RDWR]
fconfigure $channel -mode 115200,n,8,1 -blocking 0 -translation binary

spawn -open $channel

send "\r"

expect "gin: "

send "root\r"

expect "# "

send "infconfig eth0 192.168.3.14 netmask 255.255.255.0 up\r"

expect "# "

send "exit\r"

expect "gin:"
