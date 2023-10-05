import { Button, ButtonProps } from "./Button"

interface FilePicker extends ButtonProps {
    callback: (file: File) => void
    icon?: string
}

export function FilePicker(props: FilePicker) {
    let input: any
    return (
        <div>
            <input ref={input} type="file" class="hidden" onChange={e => { props.callback(e.target.files![0]) }} />
            <Button icon={props.icon} onClick={(e) => { input.click() }}></Button>
        </div>
    )
}