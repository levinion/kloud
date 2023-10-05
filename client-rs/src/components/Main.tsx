import { Dynamic } from "solid-js/web"
import Files from "./Files"
import History from "./History"
import { state, setState } from "../store/store"
import Dialog from "./basic/Dialog"

const page = {
    "files": Files,
    "history": History,
    "dialog": Dialog,
}

export default function Main() {
    return (<>
        <Dynamic component={page[state.page as keyof typeof page]}></Dynamic>
    </>)
}