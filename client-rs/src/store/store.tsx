
import { createStore } from "solid-js/store"
import { IFile } from "../interfaces/IFile"
import { createEffect } from "solid-js"
import { ILog } from "../interfaces/ILog"

export const [state, setState] = createStore({
    files: JSON.parse(localStorage.getItem("files") || "[]") as IFile[],
    page: "files",
    logs: JSON.parse(localStorage.getItem("logs") || "[]") as ILog[],
})

createEffect(() => {
    localStorage.setItem("files", JSON.stringify(state.files))
    localStorage.setItem("logs", JSON.stringify(state.logs))
})