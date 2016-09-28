use std::ops::{Deref, DerefMut};

use libc::c_void;

use weight::Weight;
use reference::Ref;
use processor::ProcessorId;
use address::GlobalAddress;



#[derive(Debug)]
pub struct Weighted<T> {
  value: T,
  weight: Weight,
}

impl<T> Weighted<T> {
  pub fn new(x: T) -> Self {
    Weighted { value: x, weight: Weight::default() }
  }

  pub fn add_weight(&mut self, w: Weight) {
    self.weight += w;
  }

  pub fn sub_weight(&mut self, w: Weight) {
    self.weight -= w;
  }

  pub fn is_orphan(&self) -> bool {
    self.weight == Weight::new(0)
  }

  pub fn get_ref(&mut self, id: ProcessorId) -> Ref {
    let w = Weight::default();
    self.add_weight(w);
    Ref::new(GlobalAddress::new(id, self.into()), w)
  }
}

impl<T> Deref for Weighted<T> {
  type Target = T;

  fn deref(&self) -> &T {
    &self.value
  }
}

impl<T> DerefMut for Weighted<T> {
  fn deref_mut(&mut self) -> &mut T {
    &mut self.value
  }
}

impl<'a, T> From<*mut c_void> for &'a mut Weighted<T> {
  fn from(p: *mut c_void) -> &'a mut Weighted<T> {
    unsafe { &mut *(p as *mut Weighted<T>) }
  }
}
