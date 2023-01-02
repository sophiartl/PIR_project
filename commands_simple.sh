

cd  simplepir-main/pir

LOG_N=14 D=2048 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_256B  &&
LOG_N=14 D=8192 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_1KB &&
LOG_N=14 D=329768 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_4KB  &&
LOG_N=14 D=65536 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_8KB &&
LOG_N=14 D=131072 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_16KB &&
LOG_N=14 D=262144 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_32KB  &&
LOG_N=14 D=524288 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_64KB &&
LOG_N=14 D=1048576 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_128KB &&
LOG_N=14 D=2097152 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-14_256KB &&


LOG_N=15 D=2048 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_256B  &&
LOG_N=15 D=8192 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_1KB &&
LOG_N=15 D=329768 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_4KB  &&
LOG_N=15 D=65536 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_8KB &&
LOG_N=15 D=131072 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_16KB &&
LOG_N=15 D=262144 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_32KB  &&
LOG_N=15 D=524288 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_64KB &&
LOG_N=15 D=1048576 go test -bench DoublePirSingle -timeout 0 -run=^$ > bench-15_128KB &&
LOG_N=15 D=2097152 go test -bench PirSingle -timeout 0 -run=^$ > bench-15_256KB &&


LOG_N=16 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_256B  &&
LOG_N=16 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_1KB &&
LOG_N=16 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_4KB  &&
LOG_N=16 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_8KB &&
LOG_N=16 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_16KB &&
LOG_N=16 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_32KB  &&
LOG_N=16 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_64KB &&
LOG_N=16 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_128KB &&
LOG_N=16 D=2097152 go test -bench PirSingle -timeout 0 -run=^$ > bench-16_256KB &&

LOG_N=17 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_256B  &&
LOG_N=17 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_1KB &&
LOG_N=17 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_4KB  &&
LOG_N=17 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_8KB &&
LOG_N=17 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_16KB &&
LOG_N=17 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_32KB  &&
LOG_N=17 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_64KB &&
LOG_N=17 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_128KB &&
LOG_N=17 D=2097152 go test -bench PirSingle -timeout 0 -run=^$ > bench-17_256KB &&

LOG_N=18 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_256B  &&
LOG_N=18 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_1KB &&
LOG_N=18 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_4KB  &&
LOG_N=18 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_8KB &&
LOG_N=18 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_16KB &&
LOG_N=18 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_32KB  &&
LOG_N=18 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_64KB &&
LOG_N=18 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_128KB &&
LOG_N=18 D=2097152 go test -bench PirSingle -timeout 0 -run=^$ > bench-18_256KB &&

LOG_N=19 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_256B  &&
LOG_N=19 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_1KB &&
LOG_N=19 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_4KB  &&
LOG_N=19 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_8KB &&
LOG_N=19 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_16KB &&
LOG_N=19 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_32KB  &&
LOG_N=19 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_64KB &&
LOG_N=19 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-19_128KB &&

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

LOG_N=22 D=2048 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_256B  &&
LOG_N=22 D=8192 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_1KB &&
LOG_N=22 D=329768 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_4KB  &&
LOG_N=22 D=65536 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_8KB &&
LOG_N=22 D=131072 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_16KB &&
LOG_N=22 D=262144 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_32KB  &&
LOG_N=22 D=524288 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_64KB &&
LOG_N=22 D=1048576 go test -bench PirSingle -timeout 0 -run=^$ > bench-22_128KB

