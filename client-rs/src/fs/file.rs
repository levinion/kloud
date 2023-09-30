use super::block::Block;
use super::block_list::BlockList;
use std::io::{self, BufRead};

#[derive(Clone, Debug)]
pub struct File {
    pub local_path: String,
    pub remote_path: String,
    pub hash: String,
    pub version: i64,
    pub blocks: BlockList,
}

const SERVERADDR: &str = "http://127.0.0.1:8080";

impl File {
    /// 创建一个文件的抽象；分别传入本地文件路径以及云端文件路径，从而将本地文件映射到云端
    pub fn new<T>(local: T, remote: T) -> Self
    where
        T: Into<String>,
    {
        let path: String = local.into();
        let hash = file_hash(&path);
        let blocks = get_blocks(&path);
        Self {
            local_path: path.clone(),
            remote_path: remote.into(),
            hash,
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
            blocks_hash_list: self.blocks.get_blocks_hash_list(),
            diff_blocks: diff_blocks.list.clone(),
        };
        dbg!(serde_json::to_string(&payload));
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

/// 根据文件内容生成当前文件哈希
fn file_hash(filename: &String) -> String {
    let content = std::fs::read_to_string(filename).unwrap();
    let blake3_hash = blake3::hash(content.as_bytes());
    #[allow(deprecated)]
    let base64_hash = base64::encode(blake3_hash.as_bytes());
    base64_hash
}

/// important!: 分块函数
fn get_blocks(filename: &String) -> BlockList {
    let f = std::fs::File::open(filename).unwrap();
    let reader = io::BufReader::new(f);
    let mut blocks = BlockList::new();
    reader
        .lines()
        .for_each(|line| blocks.append(Block::new(line.unwrap() + "\n")));
    blocks
}

/// 上传的块信息
#[derive(serde::Serialize, Debug)]
struct PostPayload {
    #[serde[rename="hashs"]]
    blocks_hash_list: Vec<String>,
    #[serde[rename="diffs"]]
    diff_blocks: Vec<Block>,
}
