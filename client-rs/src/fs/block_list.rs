use super::block::Block;

#[derive(Clone)]
pub struct BlockList {
    list: Vec<Block>,
}

impl BlockList {
    pub fn new() -> Self {
        Self { list: vec![] }
    }

    pub fn append(&mut self, block: super::block::Block) {
        self.list.push(block)
    }

    pub fn diff(&self, remote_blocks: &BlockList) -> Vec<Block> {
        self.list
            .clone()
            .into_iter()
            .filter(|item| remote_blocks.list.contains(item))
            .collect::<Vec<Block>>()
    }
}
