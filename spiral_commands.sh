#!/bin/bash

cd pir_project/PIR_project/spiral-rs/spiral-rs

echo "running commands"


cargo run --release --bin e2e 14 256 > bench-14_256B &&
cargo run --release --bin e2e 14 1024 > bench-14_1KB &&
cargo run --release --bin e2e 14 4096 > bench-14_4KB &&
cargo run --release --bin e2e 14 8192 > bench-14_8KB &&
cargo run --release --bin e2e 14 16384 > bench-14_16KB &&
cargo run --release --bin e2e 14 32768 > bench-14_32KB &&
cargo run --release --bin e2e 14 65536 > bench-14_64KB &&
cargo run --release --bin e2e 14 131072 > bench-14_128KB &&
cargo run --release --bin e2e 14 262144 > bench-14_256KB &&

cargo run --release --bin e2e 15 256 > bench-15_256B &&
cargo run --release --bin e2e 15 1024 > bench-15_1KB &&
cargo run --release --bin e2e 15 4096 > bench-15_4KB &&
cargo run --release --bin e2e 15 8192 > bench-15_8KB &&
cargo run --release --bin e2e 15 16384 > bench-15_16KB &&
cargo run --release --bin e2e 15 32768 > bench-15_32KB &&
cargo run --release --bin e2e 15 65536 > bench-15_64KB &&
cargo run --release --bin e2e 15 131072 > bench-15_128KB &&
cargo run --release --bin e2e 15 262144 > bench-15_256KB &&

cargo run --release --bin e2e 16 256 > bench-16_256B &&
cargo run --release --bin e2e 16 1024 > bench-16_1KB &&
cargo run --release --bin e2e 16 4096 > bench-16_4KB &&
cargo run --release --bin e2e 16 8192 > bench-16_8KB &&
cargo run --release --bin e2e 16 16384 > bench-16_16KB &&
cargo run --release --bin e2e 16 32768 > bench-16_32KB &&
cargo run --release --bin e2e 16 65536 > bench-16_64KB &&
cargo run --release --bin e2e 16 131072 > bench-16_128KB &&
cargo run --release --bin e2e 16 262144 > bench-16_256KB &&

cargo run --release --bin e2e 17 256 > bench-17_256B &&
cargo run --release --bin e2e 17 1024 > bench-17_1KB &&
cargo run --release --bin e2e 17 4096 > bench-17_4KB &&
cargo run --release --bin e2e 17 8192 > bench-17_8KB &&
cargo run --release --bin e2e 17 16384 > bench-17_16KB &&
cargo run --release --bin e2e 17 32768 > bench-17_32KB &&
cargo run --release --bin e2e 17 65536 > bench-17_64KB &&
cargo run --release --bin e2e 17 131072 > bench-17_128KB &&
cargo run --release --bin e2e 17 262144 > bench-17_256KB &&

cargo run --release --bin e2e 18 256 > bench-18_256B &&
cargo run --release --bin e2e 18 1024 > bench-18_1KB &&
cargo run --release --bin e2e 18 4096 > bench-18_4KB &&
cargo run --release --bin e2e 18 8192 > bench-18_8KB &&
cargo run --release --bin e2e 18 16384 > bench-18_16KB &&
cargo run --release --bin e2e 18 32768 > bench-18_32KB &&
cargo run --release --bin e2e 18 65536 > bench-18_64KB &&
cargo run --release --bin e2e 18 131072 > bench-18_128KB &&
cargo run --release --bin e2e 18 262144 > bench-18_256KB  &&

cargo run --release --bin e2e 19 256 > bench-19_256B &&
cargo run --release --bin e2e 19 1024 > bench-19_1KB &&
cargo run --release --bin e2e 19 4096 > bench-19_4KB &&
cargo run --release --bin e2e 19 8192 > bench-19_8KB &&
cargo run --release --bin e2e 19 16384 > bench-19_16KB &&
cargo run --release --bin e2e 19 32768 > bench-19_32KB &&
cargo run --release --bin e2e 19 65536 > bench-19_64KB &&
cargo run --release --bin e2e 19 131072 > bench-19_128KB &&

cargo run --release --bin e2e 20 256 > bench-20_256B &&
cargo run --release --bin e2e 20 1024 > bench-20_1KB &&
cargo run --release --bin e2e 20 4096 > bench-20_4KB &&
cargo run --release --bin e2e 20 8192 > bench-20_8KB &&
cargo run --release --bin e2e 20 16384 > bench-20_16KB &&
cargo run --release --bin e2e 20 32768 > bench-20_32KB &&
cargo run --release --bin e2e 20 65536 > bench-20_64KB &&
cargo run --release --bin e2e 20 131072 > bench-20_128KB &&

cargo run --release --bin e2e 21 256 > bench-21_256B &&
cargo run --release --bin e2e 21 1024 > bench-21_1KB &&
cargo run --release --bin e2e 21 4096 > bench-21_4KB &&
cargo run --release --bin e2e 21 8192 > bench-21_8KB &&
cargo run --release --bin e2e 21 16384 > bench-21_16KB &&
cargo run --release --bin e2e 21 32768 > bench-21_32KB &&
cargo run --release --bin e2e 21 65536 > bench-21_64KB &&
cargo run --release --bin e2e 21 131072 > bench-21_128KB &&


