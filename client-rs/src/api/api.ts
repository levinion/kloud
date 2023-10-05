import { dialog, invoke } from "@tauri-apps/api";

export function callSyncFile(local: string, remote: string) {
    invoke("sync_file", { local: local, remote: remote}).then().catch(err => console.error(err))
}

export function callOpenDialog(callback: (selected: string | string[]) => void) {
    dialog.open({ multiple: true }).then((value) => {
        if (value) {
            callback(value)
        }
    })
}

export function callOpen(local: string) {
    invoke("open", { local: local }).then()
}

// 判断路径是否文件夹
export function callIsDir(path: string) {
    return invoke("is_dir", { path: path }) as Promise<boolean>
}

export function walkDir(root: string) {
    // for all files/dirs;if dir, run addDirFiles, else run add
}