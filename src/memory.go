use std::collections::BTreeMap;
use std::mem::size_of;
use std::ops::{Deref, DerefMut};

use address::{GlobalAddress, LocalAddress};
use processor::ProcessorId;
use reference::Ref;
use weight::Weight;
use weighted::Weighted;
use object::Object;
use thunk::{Thunk, Waits};
use libc::{malloc, free, c_void};



pub trait ThunkMemory {
  fn store(&self, Thunk) -> Ref;
  fn load(&self, Ref) -> Option<&Thunk>;
}


#[derive(Debug)]
pub struct Memory {
  proc_id: ProcessorId,
  globals: BTreeMap<GlobalAddress, Thunk>,
}

impl ThunkMemory for Memory {
  fn store(&self, t: Thunk) -> Ref {
    let w: &mut Weighted<Thunk>
        = unsafe { malloc(size_of::<Weighted<Thunk>>()) }.into();
    *w = Weighted::new(t);
    w.get_ref(self.proc_id)
  }

  fn load(&self, r: Ref) -> Option<&Thunk> {
    if r.proc_id() == self.proc_id {
      let w: &Weighted<Thunk> = r.local_address().into();
      Some(w.deref())
    } else {
      self.globals.get(&r.global_address())
    }
  }
}

impl Memory {
  pub fn new(id: ProcessorId) -> Self {
    Memory {
      proc_id: id,
      globals: BTreeMap::new(),
    }
  }

  pub fn get_ref(&self, mut a: LocalAddress) -> Ref {
    a.get_ref(self.proc_id)
  }

  pub fn store_global(&mut self, a: GlobalAddress, o: Box<Object>) {
    self.globals.insert(a, o.into());
  }

  pub fn load_mut(&self, r: Ref) -> Option<&mut Thunk> {
    let a = r.local_address();

    if r.proc_id() == self.proc_id {
      let w: &mut Weighted<Thunk> = a.into();
      Some(w.deref_mut())
    } else {
      None
    }
  }

  pub fn add_weight(&self, mut a: LocalAddress, dw: Weight) {
    a.add_weight(dw);
  }

  pub fn sub_weight(&self, mut a: LocalAddress, dw: Weight) {
    a.sub_weight(dw);

    if a.is_orphan() {
      unsafe { free(a.into()) }
    }
  }
}
