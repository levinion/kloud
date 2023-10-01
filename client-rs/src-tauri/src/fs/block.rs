use super::file::File;

#[derive(Clone, serde::Deserialize, serde::Serialize, Debug, Eq, Hash)]
pub struct Block {
    pub hash: String,    // blake3 + base64
    pub content: String, // zstd default + base64
}

impl Block {
    #[allow(deprecated)]
    pub fn new(data: String) -> Self {
        let hash = blake3::hash(data.as_bytes()).to_string();
        let base = base64::encode(hash);
        let content = base64::encode(zstd::encode_all(data.as_bytes(), 0).unwrap());
        Self {
            hash: base,
            content: content,
        }
    }
}

impl PartialEq for Block {
    fn eq(&self, other: &Self) -> bool {
        self.hash == other.hash
    }
}
