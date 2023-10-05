import { createSignal, Switch, Match } from "solid-js"
import { JSX } from "solid-js/jsx-runtime"
interface DialogProps extends JSX.BaseHTMLAttributes<HTMLDivElement> {}

let [visiable, setVisiable] = createSignal(false)

export default function Dialog(props: DialogProps) {
    return (
        <>
            <Switch>
                <Match when={visiable()}>
                    <div class="absolute w-screen h-screen z-10 left-0 top-0 flex flex-col justify-center">
                        <div class=" w-2/5 bg-white shadow-xl h-48 self-center rounded-lg mx-auto">
                            {props.children}
                        </div>
                    </div>
                </Match>
            </Switch>
        </>
    )
}

export function openDialog() {
    setVisiable(true)
}

export function closeDialog() {
    setVisiable(false)
}