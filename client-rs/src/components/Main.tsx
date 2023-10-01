import { For } from "solid-js"
import { IFile } from "../interfaces/IFile"
import File from "./File"

const files: IFile[] = new Array(10).fill({ local: "/home/maruka/Downloads/test.txt", remote: "test.txt" })

export default function Main() {
    return (
        <div class="flex flex-wrap overflow-auto">
            <For each={files}>
                {(file) => (
                    <div>
                        <File local={file.local} remote={file.remote}></File>
                    </div>
                )}
            </For >
        </div >
    )
}