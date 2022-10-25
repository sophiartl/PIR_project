use std::env;
use std::fs::File;
use std::io::Write;

use spiral_rs::server::*;
use spiral_rs::util::*;
use std::time::Instant;

fn main() {
    // do so that we can get our own json file from command line
    let mut file = OpenOptions::new()
        .write(true)
        .append(true)
        .open("results.csv")
        .unwrap();
    let mut wtr = csv::Writer::from_writer(file);

    wtr.write_record(&["helloe".to_string(), 4.to_string()]);

    wtr.flush();

    let base_params = params_from_json(&CFG_10_256.replace("'", "\""));

    let params = &base_params;
    println!("db_item_size {}", base_params.db_item_size);

    let args: Vec<String> = env::args().collect();
    let inp_db_path: &String = &args[1];
    let out_db_path: &String = &args[2];

    let mut inp_file = File::open(inp_db_path).unwrap();

    let db = load_db_from_seek(params, &mut inp_file);
    let db_slice = db.as_slice();

    let mut out_file = File::create(out_db_path).unwrap();
    let now = Instant::now();
    for i in 0..db.len() {
        let coeff = db_slice[i];
        out_file.write_all(&coeff.to_ne_bytes()).unwrap();
    }
    let end = now.elapsed().as_micros();
    println!("Database proprocessing time {} us", end);
}
