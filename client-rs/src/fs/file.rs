use super::block::Block;
use super::block_list::BlockList;
use std::io::{self, BufRead};

#[derive(Clone)]
pub struct File {
    pub local_path: String,
    pub remote_path: String,
    pub hash: String,
    pub version: i64,
    blocks: BlockList,
}

impl File {
    pub fn into_blocks(&self) -> BlockList {
        let f = std::fs::File::open(self.local_path.clone()).unwrap();
        let reader = io::BufReader::new(f);
        let mut blocks = BlockList::new();
        reader
            .lines()
            .for_each(|line| blocks.append(Block::new(line.unwrap())));
        blocks
    }
    pub fn sync(&self){

    }
}
