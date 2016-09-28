use processor::ProcessorId;
use message::Message;
use rand::random;



pub struct Transceiver<R: Receiver, S: Sender> {
  receiver: R,
  senders: Vec<S>,
}

impl<R: Receiver, S: Sender> Transceiver<R, S> {
  pub fn new(r: R, ss: Vec<S>) -> Self {
    Transceiver { receiver: r, senders: ss }
  }

  pub fn send(&mut self, i: ProcessorId, m: Message) {
    self.senders[i as usize].send(m.into())
  }

  pub fn send_at_random(&mut self, m: Message) {
    let l = self.senders.len() as ProcessorId;
    self.send(random::<ProcessorId>() % l, m)
  }

  pub fn receive(&mut self) -> Option<Message> {
    self.receiver.receive().map(Message::from)
  }
}

pub trait Sender {
  fn send(&mut self, data: Vec<u8>);
}

pub trait Receiver {
  fn receive(&mut self) -> Option<Vec<u8>>;
}
