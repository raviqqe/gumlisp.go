use serde_cbor::ser;
use serde_cbor::de;

use address::{GlobalAddress, LocalAddress};
use processor::ProcessorId;
use object::SerializedObject;
use reference::Ref;
use thunk::SerializedApp;
use weight::Weight;



#[derive(Debug, Serialize, Deserialize)]
pub enum Message {
  Fetch {
    from: GlobalAddress,
    address: LocalAddress,
  },
  Resume {
    to: LocalAddress,
    address: GlobalAddress,
    object: SerializedObject,
  },

  Fish { from: ProcessorId },
  Schedule {
    task: SerializedApp,
    neighbors: Vec<(GlobalAddress, SerializedObject)>,
  },

  AddWeight { address: LocalAddress, delta: Weight },
  SubWeight { address: LocalAddress, delta: Weight },

  DepReady { to: LocalAddress },
  Finish,
}

impl From<Vec<u8>> for Message {
  fn from(data: Vec<u8>) -> Self {
    de::from_slice(data.as_ref()).unwrap()
  }
}

impl From<Message> for Vec<u8> {
  fn from(m: Message) -> Self {
    ser::to_vec(&m).unwrap()
  }
}
