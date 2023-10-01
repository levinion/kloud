import { IFile } from "../interfaces/IFile";
import { Button } from "./basic/Button";

export default function File(props: IFile) {
    return (
        <div class="p-1 m-auto flex flex-col items-center" onContextMenu={onRightClick()}>
            <Button icon="solar:file-bold" height="64"></Button>
            <p class="break-all text-black">{props.local}</p>
        </div>
    )
}

function onRightClick() {
    return (e: MouseEvent) => {
        e.preventDefault()
    }
}