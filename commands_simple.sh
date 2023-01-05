
cd  simplepir-main/pir

LOG_N=14 D=32768 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_4KB  &&


LOG_N=15 D=32768 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_4KB  
