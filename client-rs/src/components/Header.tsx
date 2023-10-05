import { tauri } from "@tauri-apps/api"
import { Button } from "./basic/Button"
import { FilePicker } from "./basic/FilePicker"
import { state, setState } from "../store/store"
import { produce } from "solid-js/store"
import { callIsDir, callOpenDialog } from "../api/api"

export default function Header() {
    return (
        <div class="flex flex-row h-full w-full items-center space-x-5">
            <input type="text" class="h-3/4 bg-zinc-800 ml-2 px-2" />
            <Button onClick={clickCallbackAddFile} icon="mdi:new-box"></Button>
        </div>
    )
}

function clickCallbackAddFile() {
    callOpenDialog((names) => {
        if (Array.isArray(names)) {
            names.forEach((name) => addFileOrDir(name))
        } else {
            addFileOrDir(names)
        }
    })
}

function addFileOrDir(local: string) {
    callIsDir(local).then(is_dir => {
        if (is_dir) {
            addDir(local)
        } else {
            addFile(local)
        }
    })
}

function addFile(local: string) {
    setState("files", produce(files => files.push({ type: "file", local: local, remote: local })))
    let t = new Date
    setState("logs", produce((log) => log.unshift({ action: "ADD", time: t.toUTCString(), file: local })))
}

function addDir(local: string) {
    setState("files", produce(files => files.push({ type: "dir", local: local, remote: local })))
    let t = new Date
    setState("logs", produce((log) => log.unshift({ action: "ADD", time: t.toUTCString(), file: local })))
}