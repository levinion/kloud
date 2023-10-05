use std::path::Path;

use futures::{future::BoxFuture, Future, FutureExt};

pub fn walk<F, Fut>(root: String, callback: F) -> BoxFuture<'static, ()>
where
    F: Fn(String) -> Fut + Copy + Send + std::marker::Sync+ 'static,
    Fut: Future<Output = ()> + Send,
{
    async move {
        for path in Path::new(&root).read_dir().unwrap() {
            let path = path.unwrap();
            // dbg!(path.path());
            if path.path().is_dir() {
                walk(path.path().to_string_lossy().to_string(), callback.clone()).await;
            } else {
                callback(path.path().to_string_lossy().to_string()).await;
            };
        }
    }
    .boxed()
}

pub fn get_remote_path(current: &str, root: &str, remote_root: &str) -> String {
    let related = get_related_path(current, root);
    Path::new(remote_root)
        .join(related)
        .to_string_lossy()
        .to_string()
}

fn get_related_path(current: &str, root: &str) -> String {
    let current_path = Path::new(current);
    current_path
        .strip_prefix(root)
        .unwrap()
        .to_string_lossy()
        .to_string()
}

#[cfg(test)]
mod tests {
    use super::{get_related_path, get_remote_path};

    #[test]
    fn test_get_related_path() {
        let related_path = get_related_path("/home/user/dir/file.txt", "/home/user/dir");
        assert_eq!(related_path, String::from("file.txt"))
    }

    #[test]
    fn test_get_remote_path() {
        let remote_path = get_remote_path("/home/user/dir/file.txt", "/home/user/dir", "/dir");
        assert_eq!(remote_path, String::from("/dir/file.txt"))
    }
}
