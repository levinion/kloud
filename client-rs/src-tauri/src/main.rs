// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]
#![feature(async_closure)]
use kloud::fs;
use std::{path::Path, process::Command};
mod sync;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command

/// 同步文件
#[tauri::command]
async fn sync_file(local: String, remote: String) {
    let local: &str = local.leak();
    let remote: &str = remote.leak();
    if is_dir(local.to_string()) {
        sync::walk(local.to_string(), async move |path| {
            tokio::spawn(async move {
                let remote_path = sync::get_remote_path(&path, local, remote);
                let mut file = fs::file::File::new(path.to_string(), remote_path);
                file.sync().await.unwrap();
            })
            .await
            .unwrap();
        }).await;
    } else {
        let mut file = fs::file::File::new(local.to_string(), remote.to_string());
        file.sync().await.unwrap();
    }
}

/// 打开文件
#[tauri::command]
fn open(local: String) {
    Command::new("xdg-open").arg(local).spawn().unwrap();
}

#[tauri::command]
fn is_dir(path: String) -> bool {
    Path::new(&path).is_dir()
}

#[tokio::main]
async fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![sync_file, open, is_dir])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
