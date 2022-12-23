use csv::Writer;
use rand::thread_rng;
use rand::Rng;
use spiral_rs::arith::*;
use spiral_rs::client::*;
use spiral_rs::params::*;
use spiral_rs::server::*;
use spiral_rs::util::*;
use std::env;
use std::fs;

use std::path::Path;
use std::time::Instant;

fn print_params_summary(params: &Params) {
    let db_elem_size = params.item_size();
    let total_size = params.num_items() * db_elem_size;
    println!(
        "Using a {} x {} byte database ({} bytes total)",
        params.num_items(),
        db_elem_size,
        total_size
    );
}

fn main() {
    let params;
    let args: Vec<String> = env::args().collect();

    // new configuration
    let _cfg_10_256 = r#"
    {
        "n": 2,
        "nu_1": 6,
        "nu_2": 4,
        "p": 4,
        's_e': 87.62938774292914,
        "q2_bits": 13,
        "t_gsw": 4,
        "t_conv": 4,
        "t_exp_left": 4,
        "t_exp_right": 56,
        "instances": 1,
        "db_item_size": 2048
    }
"#;
    let mut target_num_log2: usize = 0;
    let mut item_size_bytes: usize = 0;
    if args.len() == 2 {
        //input parameter file
        let inp_params_fname = &args[1];
        let params_store_str = fs::read_to_string(inp_params_fname).unwrap();
        params = params_from_json(&params_store_str.replace("'", "\""));
    } else {
        // get predefined parameters
        target_num_log2 = args[1].parse().unwrap(); // power of 2 DB size
        item_size_bytes = args[2].parse().unwrap();
        // println!("hola {}", item_size_bytes);

        params = get_params_from_store(target_num_log2, item_size_bytes);
    }

    print_params_summary(&params);

    let mut rng = thread_rng();
    let idx_target: usize = rng.gen::<usize>() % params.num_items();

    println!(
        "fetching index {} out of {} items",
        idx_target,
        params.num_items()
    );
    println!("initializing client");
    let mut client = Client::init(&params);
    println!("generating public parameters");
    let pub_params = client.generate_keys();
    let pub_params_buf = pub_params.serialize();

    let mut procc: [u128; 5] = [0, 0, 0, 0, 0];
    let mut gen: [u128; 5] = [0, 0, 0, 0, 0];
    let mut dec: [u128; 5] = [0, 0, 0, 0, 0];
    let mut quer_size = 0;
    let mut asnw_size = 0;

    println!("public parameters size: {} bytes", pub_params_buf.len());
    for test in 0..5 {
        let now_gen = Instant::now();
        let query = client.generate_query(idx_target);
        let time_gen = now_gen.elapsed().as_micros();
        let query_buf = query.serialize();
        println!("initial query size: {} bytes", query_buf.len());
        quer_size = query_buf.len();
        gen[test] = time_gen;
        println!("generating db");
        let (corr_item, db) = generate_random_db_and_get_item(&params, idx_target);

        println!("processing query");
        let now = Instant::now();
        let response = process_query(&params, &pub_params, &query, db.as_slice());
        let end = now.elapsed().as_micros();
        println!("processing query (took {} us).", end);
        procc[test] = end;
        println!("response size: {} bytes", response.len());
        asnw_size = response.len();
        println!("decoding response");
        let now_dec = Instant::now();
        let result = client.decode_response(response.as_slice());
        let end_dec = now_dec.elapsed().as_micros();
        println!("decoding reponse (took {} us).", end_dec);
        dec[test] = end;
        let p_bits = log2_ceil(params.pt_modulus) as usize;
        let corr_result = corr_item.to_vec(p_bits, params.modp_words_per_chunk());

        assert_eq!(result.len(), corr_result.len());
        for z in 0..corr_result.len() {
            assert_eq!(result[z], corr_result[z], "error in response at {:?}", z);
        }

        println!("completed correctly!");
    }

    let sum: u128 = gen.iter().sum();
    let gen_avg = sum as f64 / gen.len() as f64;
    let sum3: u128 = procc.iter().sum();
    let proc_avg = sum3 as f64 / procc.len() as f64;

    let file_path = Path::new("results.csv");
    let mut writer = Writer::from_path(file_path).expect("Unable to write file");
    writer.write_record(&[
        target_num_log2.to_string(),
        item_size_bytes.to_string(),
        gen_avg.to_string(),
        proc_avg.to_string(),
        pub_params_buf.len().to_string(),
        quer_size.to_string(),
        asnw_size.to_string(),
    ]);

    // Flush the writer to ensure all data is written to the file
    writer.flush();

    // let data = format!(
    //     "
    // n={}; db size={}; entry size ={}; record size ={};
    // public parameter size={} bytes;
    // query_size={} bytes; query generation time={} us;
    // answer size={} bytes; query processing time={} us;
    // decoding time={}.
    // ",
    //     target_num_log2,
    //     params.num_items().to_string(),
    //     item_size_bytes,
    //     params.item_size().to_string(),
    //     pub_params_buf.len().to_string(),
    //     query_buf.len().to_string(),
    //     time_gen,
    //     response.len().to_string(),
    //     end.to_string(),
    //     end_dec.to_string(),
    // );

    // let path_to: String = format!("spiral-{}_{}.txt", target_num_log2, item_size_bytes / 1024);

    // fs::write(path_to, data).expect("Unable to write file");

    // let file = OpenOptions::new()
    //     .write(true)
    //     .append(true)
    //     .open("results.csv")
    //     .unwrap();
    // let mut wtr = csv::Writer::from_writer(file);

    // wtr.write_record(&[
    // params.num_items().to_string(),
    // params.item_size().to_string(),
    // pub_params_buf.len().to_string(),
    // query_buf.len().to_string(),
    // end.to_string(),
    // response.len().to_string(),
    // end_dec.to_string(),
    // ]);

    // wtr.flush();
}
