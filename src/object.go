use std::any::Any;
use std::fmt;
use std::fmt::Debug;



pub trait Object: Any {
  fn serialize(&self) -> Vec<u8>;
}

pub trait ObjectLike: Object {
  fn deserialize(Vec<u8>) -> Box<Object>;
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SerializedObject {
  id: u64,
  data: Vec<u8>,
}

impl<'a> From<&'a Object> for SerializedObject {
  fn from(o: &'a Object) -> Self {
    let mut h = 0; // dummy value

    unimplemented!(); // get_type_id()

    SerializedObject { id: h, data: o.serialize() }
  }
}

impl Into<Box<Object>> for SerializedObject {
  fn into(self) -> Box<Object> {
    unimplemented!() // TYPES[self.id].deserialize(s.data)
  }
}

impl Debug for Box<Object> {
  fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
    write!(f, "Object")
  }
}
