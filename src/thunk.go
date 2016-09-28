use address::GlobalAddress;
use object::Object;
use reference::Ref;



pub type Waits = u64;

pub type BlackHole = Vec<GlobalAddress>;

#[derive(Debug)]
pub enum Thunk {
  Object(Box<Object>),
  App {
    app: App,
    black_hole: BlackHole,
    waits: Waits,
  },
}

impl Thunk {
  pub fn update(&mut self, v: ThunkValue) -> Vec<GlobalAddress> {
    assert!(!self.is_object());

    match v {
      ThunkValue::Object(o) => self.update_with_object(o),
      ThunkValue::App(a) => {
        self.update_with_app(a);
        Vec::new()
      }
    }
  }

  fn update_with_object(&mut self, o: Box<Object>) -> Vec<GlobalAddress> {
    let gas = self.take_from_black_hole();

    *self = Thunk::Object(o);

    gas
  }

  fn update_with_app(&mut self, a: App) {
    match *self {
      Thunk::App { ref mut app, waits, .. } => {
        assert_eq!(waits, 0);
        *app = a;
      },
      _ => unreachable!(),
    }
  }

  pub fn is_object(&self) -> bool {
    match *self {
      Thunk::Object(_) => true,
      _ => false,
    }
  }

  pub fn object(&self) -> Option<&Object> {
    match *self {
      Thunk::Object(ref o) => Some(o.as_ref()),
      _ => None,
    }
  }

  pub fn set_waits(&mut self, w: Waits) {
    match *self {
      Thunk::App { ref mut waits, .. } => *waits = w,
      _ => unreachable!(),
    }
  }

  pub fn decre_waits(&mut self) {
    match *self {
      Thunk::App { ref mut waits, .. } => *waits -= 1,
      _ => unreachable!(),
    }
  }

  pub fn is_ready(&self) -> bool {
    match *self {
      Thunk::App { waits, .. } => waits == 0,
      _ => unreachable!(),
    }
  }

  pub fn put_into_black_hole(&mut self, a: GlobalAddress) {
    match *self {
      Thunk::App { ref mut waits, .. } => *waits -= 1,
      _ => unreachable!(),
    }
  }

  fn take_from_black_hole(&mut self) -> Vec<GlobalAddress> {
    match *self {
      Thunk::App { ref mut black_hole, .. } => {
        let l = black_hole.len();
        black_hole.split_off(l)
      }
      _ => unreachable!(),
    }
  }
}

impl From<Box<Object>> for Thunk {
  fn from(o: Box<Object>) -> Self {
    Thunk::Object(o)
  }
}

#[derive(Debug, Serialize, Deserialize)]
pub struct App {
  func: Ref,
  arg: Ref,
}

#[derive(Debug)]
pub enum ThunkValue {
  Object(Box<Object>),
  App(App),
}

impl From<ThunkValue> for Thunk {
  fn from(v: ThunkValue) -> Self {
    match v {
      ThunkValue::Object(o) => Thunk::Object(o),
      ThunkValue::App(a) => Thunk::App {
        app: a,
        black_hole: BlackHole::new(),
        waits: 0,
      },
    }
  }
}


#[derive(Debug, Serialize, Deserialize)]
pub struct SerializedApp {
  app: App,
  black_hole: BlackHole,
  waits: Waits,
}

impl From<SerializedApp> for Thunk {
  fn from(SerializedApp { app, black_hole, waits }: SerializedApp) -> Thunk {
    Thunk::App { app: app, black_hole: black_hole, waits: waits }
  }
}

impl From<Thunk> for SerializedApp {
  fn from(t: Thunk) -> SerializedApp {
    match t {
      Thunk::App { app, black_hole, waits } => SerializedApp {
        app: app,
        black_hole: black_hole,
        waits: waits
      },
      _ => panic!("The Thunk is not a Thunk::App. (value: {:?})", t),
    }
  }
}
