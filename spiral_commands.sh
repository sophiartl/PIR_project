#!/bin/bash

cd spiral-rs/spiral-rs

echo "running  commands"

cargo run --release --bin e2e 20 256 > bench-20_256B &&
cargo run --release --bin e2e 20 1024 > bench-20_1KB &&
cargo run --release --bin e2e 20 4096 > bench-20_4KB &&
cargo run --release --bin e2e 20 8192 > bench-20_8KB &&
cargo run --release --bin e2e 20 16384 > bench-20_16KB &&
cargo run --release --bin e2e 20 32768 > bench-20_32KB &&
cargo run --release --bin e2e 20 65536 > bench-20_64KB &&
cargo run --release --bin e2e 20 102400 > bench-20_100KB &&

cargo run --release --bin e2e 21 256 > bench-21_256B &&
cargo run --release --bin e2e 21 1024 > bench-21_1KB &&
cargo run --release --bin e2e 21 4096 > bench-21_4KB &&
cargo run --release --bin e2e 21 8192 > bench-21_8KB &&
cargo run --release --bin e2e 21 16384 > bench-21_16KB &&
cargo run --release --bin e2e 21 32768 > bench-21_32KB &&
cargo run --release --bin e2e 21 65536 > bench-21_64KB &&
cargo run --release --bin e2e 21 102400 > bench-21_100KB 


