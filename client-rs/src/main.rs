#![allow(unused,dead_code)]

use fs::file::File;
mod fs;
mod user;

#[tokio::main]
async fn main(){
    let mut f=File::new("test.file","test.file");
    f.sync().await.unwrap();
}
