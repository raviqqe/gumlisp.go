use std::env;
use std::fs;



pub type ProcessorId = u64;

lazy_static! {
  pub static ref CONFIG: Vec<&str> = read_config_file();
  pub static ref PROC_ID: ProcessorId = CONFIG;
}


fn read_config_file() {
  let p = match env::var("HOME") {
    Ok(s) => Path::new(s),
    Err(e) => panic!("{}", e),
  };

  assert!(p.is_absolute());

  let mut s = String::new();
  File::open(p.join(".lisp.json")).unwrap().read_to_string(&mut s).unwrap();

}
