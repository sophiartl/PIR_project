use std::env;
use std::fs::File;
use std::io::Write;

use spiral_rs::server::*;
use spiral_rs::util::*;

fn main() {
    let base_params = params_from_json(&CFG_10_256.replace("'", "\""));

    let params = &base_params;
    println!("db_item_size {}", base_params.db_item_size);

    let args: Vec<String> = env::args().collect();
    let inp_db_path: &String = &args[1];
    let out_db_path: &String = &args[2];

    let mut inp_file = File::open(inp_db_path).unwrap();

    let db = load_db_from_seek(params, &mut inp_file);
    let db_slice = db.as_slice();
    println!("before output");

    let mut out_file = File::create(out_db_path).unwrap();
    println!("db length: {}", db.len());
    for i in 0..db.len() {
        let coeff = db_slice[i];
        out_file.write_all(&coeff.to_ne_bytes()).unwrap();
    }
}
