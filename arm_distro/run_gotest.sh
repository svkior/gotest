chmod +x /root/gotest
fpga_loader 1.bit
echo "Hello from ARM" >test
./gotest test

