use std::str::FromStr;
use metalog::{set_logger, Log, LogLevelFilter, LogMetadata, LogRecord,
              MaxLogLevelFilter};



struct Logger { max_log_level: MaxLogLevelFilter }

impl Log for Logger {
  fn enabled(&self, metadata: &LogMetadata) -> bool {
    metadata.level() <= self.max_log_level.get()
  }

  fn log(&self, record: &LogRecord) {
    if self.enabled(record.metadata()) {
      println!("{}: {}", record.level(), record.args());
    }
  }
}

pub fn init(log_level: &str) {
  set_logger(|max_ll| {
    max_ll.set(LogLevelFilter::from_str(log_level)
               .unwrap_or(LogLevelFilter::Error));
    Box::new(Logger { max_log_level: max_ll })
  }).unwrap();
}
