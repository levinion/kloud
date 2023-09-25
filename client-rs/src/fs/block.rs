use super::file::File;

#[derive(Clone)]
pub struct Block {
    pub hash: String,
    pub belong: Vec<File>,
    pub content: String,
}

impl Block {
    #[allow(deprecated)]
    pub fn new(data: String) -> Self {
        let hash = blake3::hash(data.as_bytes()).to_string();
        let base = base64::encode(hash);
        Self {
            hash: base,
            belong: vec![],
            content: data,
        }
    }
}

impl PartialEq for Block{
    fn eq(&self, other: &Self) -> bool {
        self.hash==other.hash
    }
}
