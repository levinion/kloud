use super::block::Block;

#[derive(Clone, serde::Deserialize, serde::Serialize, Debug)]
pub struct BlockList {
    pub list: Vec<Block>,
}

impl BlockList {
    pub fn new() -> Self {
        Self { list: vec![] }
    }

    pub fn append(&mut self, block: super::block::Block) {
        self.list.push(block)
    }

    pub fn diff(&self, remote_blocks_hashs: &Vec<String>) -> BlockList {
        let mut diff_blocks = BlockList::new();
        let blocks = self
            .list
            .clone()
            .into_iter()
            .filter(|item| !remote_blocks_hashs.contains(&item.hash))
            .collect::<Vec<Block>>();
        diff_blocks.list = blocks;
        diff_blocks
    }

    pub fn get_blocks_hash_list(&self) -> Vec<String> {
        self.list.iter().map(|block| block.hash.clone()).collect()
    }
}
