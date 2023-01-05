#!/bin/bash

cd spiral-rs/spiral-rs

echo "running  commands"

cargo run --release --bin e2e 14 262144 > bench-20_256B 
