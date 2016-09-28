use address::{GlobalAddress, LocalAddress};
use processor::ProcessorId;
use weight::Weight;



#[derive(Debug, Serialize, Deserialize)]
pub struct Ref {
  global_address: GlobalAddress,
  weight: Weight,
}

impl Ref {
  pub fn new(a: GlobalAddress, w: Weight) -> Ref {
    Ref { global_address: a, weight: w }
  }

  pub fn proc_id(&self) -> ProcessorId {
    self.global_address.proc_id
  }

  pub fn local_address(&self) -> LocalAddress {
    self.global_address.local_address
  }

  pub fn global_address(&self) -> GlobalAddress {
    self.global_address
  }

  pub fn delete(self) -> (GlobalAddress, Weight) {
    (self.global_address, self.weight)
  }

  pub fn split(&mut self) -> Option<Self> {
    match self.weight.split() {
      Some(w) => Some(Ref {
        global_address: self.global_address,
        weight: w,
      }),
      None => None,
    }
  }
}
