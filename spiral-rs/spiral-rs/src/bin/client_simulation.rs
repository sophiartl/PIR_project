use rand::thread_rng;
use rand::Rng;
use spiral_rs::arith::*;
use spiral_rs::client::*;
use spiral_rs::params::*;
use spiral_rs::server::*;
use spiral_rs::util::*;
use std::env;

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

    if args.len() == 2 {
        //input parameter file
        // let inp_params_fname = &args[1];

        params = params_from_json(&_cfg_10_256.replace("'", "\""));
    } else {
        // get predefined parameters
        let target_num_log2: usize = args[1].parse().unwrap(); // power of 2 DB size
        let item_size_bytes: usize = args[2].parse().unwrap();

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
    println!("public parameters size: {} bytes", pub_params_buf.len());
    let query = client.generate_query(idx_target);
    let query_buf = query.serialize();
    println!("initial query size: {} bytes", query_buf.len());

    println!("generating db");
    let (corr_item, db) = generate_random_db_and_get_item(&params, idx_target);

    println!("processing query");
    let now = Instant::now();
    let response = process_query(&params, &pub_params, &query, db.as_slice());
    println!("done processing (took {} us).", now.elapsed().as_micros());
    println!("response size: {} bytes", response.len());

    println!("decoding response");
    let result = client.decode_response(response.as_slice());

    let p_bits = log2_ceil(params.pt_modulus) as usize;
    let corr_result = corr_item.to_vec(p_bits, params.modp_words_per_chunk());

    assert_eq!(result.len(), corr_result.len());
    for z in 0..corr_result.len() {
        assert_eq!(result[z], corr_result[z], "error in response at {:?}", z);
    }

    println!("completed correctly!");
}
