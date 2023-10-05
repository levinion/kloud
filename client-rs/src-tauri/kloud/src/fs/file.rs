use anyhow::Ok;

use super::block::Block;
use super::block_list::BlockList;
use std::io::{self, BufRead, BufReader, Read};

#[derive(Clone, Debug)]
pub struct File {
    pub local_path: String,
    pub remote_path: String, // base64
    pub version: i64,
    pub blocks: BlockList,
}

const SERVERADDR: &str = "http://127.0.0.1:8080";

impl File {
    /// 创建一个文件的抽象；分别传入本地文件路径以及云端文件路径，从而将本地文件映射到云端
    pub fn new(local: String, remote: String) -> Self {
        let blocks = get_blocks(&local);
        #[allow(deprecated)]
        let remote_base62 = base_62::encode(remote.as_bytes());
        Self {
            local_path: local,
            remote_path: remote_base62,
            version: chrono::Utc::now().timestamp(),
            blocks,
        }
    }

    fn update_blocks(&self) -> BlockList {
        get_blocks(&self.local_path)
    }

    fn update_version(&mut self) {
        self.version = chrono::Utc::now().timestamp();
    }

    /// 同步。从云端拉取旧信息，对比并上传新的存储块，从而更新云端数据
    pub async fn sync(&mut self) -> anyhow::Result<()> {
        self.update_version();
        let online_blocks = self.fetch_online_blocks_hashs().await?;
        let diff_blocks = self.blocks.diff(&online_blocks);
        self.post_diff_blocks(&diff_blocks).await?;
        Ok(())
    }

    /// 获取云端文件存储块
    pub async fn fetch_online_blocks_hashs(&self) -> anyhow::Result<Vec<String>> {
        let path = SERVERADDR
            .parse::<url::Url>()
            .unwrap()
            .join("file/")?
            .join("maruka/")?
            .join(self.remote_path.clone().as_str())?;
        let res = reqwest::get(path.to_string()).await?;
        let hashs = res.json::<Vec<String>>().await?;
        Ok(hashs)
    }

    /// 上传当前文件存储块信息及差异块
    pub async fn post_diff_blocks(&self, diff_blocks: &BlockList) -> anyhow::Result<()> {
        let client = reqwest::Client::new();
        let payload = PostPayload {
            hashs: self.blocks.get_blocks_hash_list(),
            diffs: diff_blocks.list.clone(),
        };
        // dbg!(serde_json::to_string(&payload));
        let path = SERVERADDR
            .parse::<url::Url>()
            .unwrap()
            .join("file/")?
            .join("maruka/")?
            .join(self.remote_path.clone().as_str())?;
        let _ = client.post(path).json(&payload).send().await?;
        Ok(())
    }
}

/// important!: 分块函数
fn get_blocks(filename: &String) -> BlockList {
    let mut f = std::fs::File::open(filename).unwrap();
    let blocks =
        get_utf8_blocks(f.by_ref()).unwrap_or_else(|_| get_non_utf8_blocks(f.by_ref()).unwrap());
    blocks
}

fn get_utf8_blocks(file: &mut std::fs::File) -> anyhow::Result<BlockList> {
    let reader = io::BufReader::new(file);
    let mut blocks = BlockList::new();
    for line in reader.lines() {
        #[allow(deprecated)]
        let content = base64::encode(line? + "\n");
        blocks.append(Block::new(content));
    }
    Ok(blocks)
}

// TODO: 二进制文件分块
fn get_non_utf8_blocks(file: &mut std::fs::File) -> anyhow::Result<BlockList> {
    let mut blocks = BlockList::new();
    let mut buffer = Vec::new();
    file.read_to_end(&mut buffer)?;
    #[allow(deprecated)]
    let b64 = base64::encode(buffer.to_vec());
    // 按4MB每块读取
    let chunk_size = 4 * 1024 * 1024; // 4MB
    for chunk in b64.as_bytes().chunks(chunk_size) {
        blocks.append(Block::new(String::from_utf8(chunk.to_vec()).unwrap()));
    }
    Ok(blocks)
}

/// 上传的块信息
#[derive(serde::Serialize, Debug)]
struct PostPayload {
    hashs: Vec<String>,
    diffs: Vec<Block>,
}
