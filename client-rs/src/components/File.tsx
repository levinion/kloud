import { JSX } from "solid-js/jsx-runtime";
import { IFile } from "../interfaces/IFile";
import { Button } from "./basic/Button";

interface FileProps extends JSX.BaseHTMLAttributes<HTMLSpanElement> {
    file: IFile
}

export default function File(props: FileProps) {
    return (
        <>
            <div class="p-1 flex flex-col items-center text-center">
                <Button onClick={props.onClick} onContextMenu={props.onContextMenu} icon={props.file.type==="file"?"solar:file-bold":"game-icons:files"} height="56"></Button>
                <p class="break-word text-black w-40">{props.file.local}</p>
            </div>
        </>
    )
}