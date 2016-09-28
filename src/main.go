package main

import (
  "fmt"
  "github.com/docopt/docopt-go"
  "encoding/gob"
)


func get_args() {
  let usage = "
Lisp?

Usage:
  lisp [-l <log_level>] [-p <proc_id>] <filename>
  lisp (-h | --help)

Options:
  -l, --log-level <log_level>  Set log level. Valid values are \"error\"
                               (default), \"warn\", \"info\", \"debug\",
                               \"trace\", and \"off\".
  -p, --proc-id <proc_id>
  -h, --help  Show help.
";

  Docopt::new(usage).and_then(|d| d.parse()).unwrap_or_else(|e| e.exit())
}

fn main() {
  let args = get_args();

  log::init(args.get_str("--log-level"));

  let i = parse_proc_id(args.get_str("--proc-id"));
  let mut p = Processor::new(i, read_config_file());

  if i == MASTER_ID {
    p.run_as_master(read_file(args.get_str("<filename>")));
  } else {
    p.run_as_slave();
  }
}

fn read_file(f: &str) -> String {
  let mut s = String::new();

  let n = File::open(f).unwrap().read_to_string(&mut s).unwrap();
  assert_eq!(n, s.len());

  s
}

fn parse_proc_id(s: &str) -> ProcessorId {
  if s == "" {
    MASTER_ID
  } else {
    ProcessorId::from_str(s).unwrap()
  }
}

fn read_config_file() -> Vec<Address> {
  Vec::from_iter(read_file("procs.conf").lines().map(|s| s.trim().into())
                                                .filter(|s| s != ""))
}
