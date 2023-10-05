use std::collections::HashSet;

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

    pub fn diff(&self, other_blocks_hashs: &Vec<String>) -> BlockList {
        let mut diff_blocks = BlockList::new();
        let other_blocks_set: HashSet<_> = other_blocks_hashs.into_iter().collect();
        let blocks = self
            .list
            .clone()
            .into_iter()
            .filter(|item| !other_blocks_set.contains(&item.hash)) // 列表中无法找到当前块
            .collect::<HashSet<_>>() // 当前块不重复
            .into_iter()
            .collect::<Vec<_>>();
        diff_blocks.list = blocks;
        diff_blocks
    }

    pub fn get_blocks_hash_list(&self) -> Vec<String> {
        self.list.iter().map(|block| block.hash.clone()).collect()
    }
}
