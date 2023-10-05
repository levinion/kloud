import { For, createSignal, Match, Switch } from "solid-js"
import { IFile } from "../interfaces/IFile"
import { ContextMenu, createContextMenu } from "./basic/ContextMenu"
import File from "./File"
import { state, setState } from "../store/store"
import { produce } from "solid-js/store"
import { callSyncFile, callOpen } from "../api/api"

const menu = createContextMenu([
    { display: "DELETE", func: () => { deleteSyncFile(selectedFile!) } },
    { display: "SYNC", func: () => { syncFileOrDir(selectedFile!) } }
])

let selectedFile: IFile | null = null

export default function Files() {
    return (
        <>
            <Switch fallback={<></>}>
                <Match when={state.files.length !== 0}>
                    <div class="flex flex-wrap w-full">
                        <For each={state.files}>
                            {(file) => (
                                <div>
                                    <File file={file} onClick={() => onFileClick(file)} onContextMenu={e => onFileRightClick(e, file)} ></File>
                                </div>
                            )}
                        </For >
                    </div>
                    <ContextMenu visiable={menu.visiable()} position={menu.position()} callbacks={menu.callbacks}></ContextMenu>
                </Match>
            </Switch>
        </>
    )
}

function onFileClick(file: IFile) {
    callOpen(file.local)
}

function onFileRightClick(e: MouseEvent, file: IFile) {
    selectedFile = file
    menu.open(e)
}

function deleteSyncFile(file: IFile) {
    setState("files", produce((files) => { files.splice(files.findIndex((f) => f.local === file.local), 1) }))
    let t = new Date
    setState("logs", produce((log) => log.unshift({ action: "DELETE", time: t.toUTCString(), file: file.local })))
}

function syncFileOrDir(file: IFile) {
    callSyncFile(file.local, file.remote)
    let t = new Date
    setState("logs", produce((log) => log.unshift({ action: "SYNC", time: t.toUTCString(), file: file.local })))
}