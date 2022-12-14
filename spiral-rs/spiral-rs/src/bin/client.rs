use serde_json::Value;
use spiral_rs::client::*;
use spiral_rs::util::*;
use std::env;
use std::fs::OpenOptions;
use std::time::Instant;

fn send_api_req_text(api_url: &str, path: &str, data: Vec<u8>) -> Result<String, reqwest::Error> {
    let client = reqwest::blocking::Client::builder()
        .timeout(None)
        .build()
        .unwrap();

    client
        .post(format!("{}{}", api_url, path))
        .body(data)
        .send()?
        .text()
}

fn send_api_req_vec(api_url: &str, path: &str, data: Vec<u8>) -> Result<Vec<u8>, reqwest::Error> {
    let client = reqwest::blocking::Client::builder()
        .timeout(None)
        .build()
        .unwrap();
    Ok(client
        .post(format!("{}{}", api_url, path))
        .body(data)
        .send()?
        .bytes()?
        .to_vec())
}

fn main() {
    let file = OpenOptions::new()
        .write(true)
        .append(true)
        .open("console.csv")
        .unwrap();
    let mut wtr = csv::Writer::from_writer(file);

    wtr.write_record(&["hello", "hi"]);

    wtr.flush();

    let _cfg_expand = r#"
        {'n': 2,
        'nu_1': 10,
        'nu_2': 6,
        'p': 512,
        'q2_bits': 21,
        's_e': 85.83255142749422,
        't_gsw': 10,
        't_conv': 4,
        't_exp_left': 16,
        't_exp_right': 56,
        'instances': 11,
        'db_item_size': 100000 }
    "#;

    //change parameters
    //define parameters
    //do json read put parrms in seperate file /Users/sophiaartioli/PIR_project/PIR_project/spiral-rs/spiral-rs/src/bin/client.rs
    //whar stats to collect
    let _cfg_direct = r#"
        {'direct_upload': 1,
        'n': 5,
        'nu_1': 11,
        'nu_2': 3,
        'p': 65536,
        'q2_bits': 27,
        's_e': 57.793748020122216,
        't_gsw': 3,
        't_conv': 56,
        't_exp_left': 56,
        't_exp_right': 56}
    "#;
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

    let cfg = _cfg_10_256;
    let cfg = cfg.replace("'", "\"");
    let params = params_from_json(&cfg);

    let args: Vec<String> = env::args().collect();
    let api_url = &args[1];

    let idx_target: usize = (&args[2]).parse().unwrap();

    println!("initializing client");
    let mut c = Client::init(&params);
    println!("generating public parameters");
    let pub_params = c.generate_keys();
    let pub_params_buf = pub_params.serialize();
    println!("pub_params size {}", pub_params_buf.len());
    let query = c.generate_query(idx_target);
    let mut query_buf = query.serialize();
    println!("initial query size {}", query_buf.len());

    let setup_resp_str = send_api_req_text(api_url, "/setup", pub_params_buf).unwrap();
    println!("{}", setup_resp_str);
    let resp_json: Value = serde_json::from_str(&setup_resp_str).unwrap();
    let id = resp_json["id"].as_str().unwrap();
    let mut full_query_buf = id.as_bytes().to_vec();
    full_query_buf.append(&mut query_buf);

    let now = Instant::now();
    let query_resp = send_api_req_vec(api_url, "/query", full_query_buf).unwrap();
    let duration = now.elapsed().as_millis();
    println!("duration of query processing is {} ms", duration);
    println!("query_resp len {}", query_resp.len());
    // println!("query_resp {:x?}", query_resp);

    let result = c.decode_response(query_resp.as_slice());
    println!("{:x?}", result);
}
