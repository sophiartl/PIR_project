
cd  simplepir-main/pir

LOG_N=14 D=32768 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_4KB  &&


LOG_N=20 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_256B  &&
LOG_N=20 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_1KB &&
LOG_N=20 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_4KB  &&
LOG_N=20 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_8KB &&
LOG_N=20 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_16KB &&
LOG_N=20 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_32KB  &&
LOG_N=20 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_64KB &&
LOG_N=20 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-20_128KB &&

LOG_N=21 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_256B  &&
LOG_N=21 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_1KB &&
LOG_N=21 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_4KB  &&
LOG_N=21 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_8KB &&
LOG_N=21 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_16KB &&
LOG_N=21 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_32KB  &&
LOG_N=21 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_64KB &&
LOG_N=21 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-21_128KB &&

LOG_N=14 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_256B  &&
LOG_N=14 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_1KB &&
LOG_N=14 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_4KB  &&
LOG_N=14 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-14_8KB &&


LOG_N=15 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_256B  &&
LOG_N=15 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_1KB &&
LOG_N=15 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_4KB  &&
LOG_N=15 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_8KB &&


