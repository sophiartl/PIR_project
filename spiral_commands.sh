#!/bin/bash

cd pir_project/PIR_project/spiral-rs/spiral-rs

echo "running commands"

cargo run --release --bin e2e 15 256 > bench-15_256B &&
cargo run --release --bin e2e 15 1024 > bench-15_1KB &&
cargo run --release --bin e2e 15 4096 > bench-15_4KB &&
cargo run --release --bin e2e 15 8192 > bench-15_8KB &&
cargo run --release --bin e2e 15 16384 > bench-15_16KB &&
cargo run --release --bin e2e 15 32768 > bench-15_32KB &&
cargo run --release --bin e2e 15 65536 > bench-15_64KB &&
cargo run --release --bin e2e 15 131072 > bench-15_128KB &&
cargo run --release --bin e2e 15 262144 > bench-15_256KB 