use processor::ProcessorId;
use reference::Ref;
use memory::ThunkMemory;
use message::Message;
use thunk::ThunkValue;



pub type Func = Fn (&mut ThunkMemory, Ref) -> FuncResult;
pub type FuncResult = Result<ThunkValue, Vec<(ProcessorId, Message)>>;
