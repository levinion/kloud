import { Show, createSignal, For } from "solid-js"

export interface Callback {
    display: string, func: () => void
}

export type Callbacks = Callback[]

interface ContextMenuProps {
    visiable?: boolean,
    position?: { x: number, y: number }
    callbacks?: Callbacks
}

// 接收一个菜单列表，返回菜单对象的回调函数
export function createContextMenu(callbacks: Callbacks) {
    const [visiable, setVisiable] = createSignal(false)
    const [position, setPosition] = createSignal({ x: 0, y: 0 })
    const open = (e: MouseEvent) => {
        e.preventDefault()
        setPosition({ x: e.pageX, y: e.pageY })
        setVisiable(true)
        addEventListener("click", () => { setVisiable(false); })
    }
    return {
        visiable, setVisiable,
        position, setPosition,
        open,callbacks
    }
}



export function ContextMenu(props: ContextMenuProps) {
    return (
        <Show when={props.visiable || false}>
            {/* mask */}
            <div class="absolute w-screen h-screen left-0 top-0 z-10">
                {/* menu */}
                <div class="absolute bg-slate-700 rounded-lg 
            p-2 flex flex-col space-y-2 border-2 border-purple-800"
                    style={`left:${props.position?.x}px;top:${props.position?.y}px `}>
                    <For each={props.callbacks}>
                        {(callback) => (<button onClick={() => callback.func()}>{callback.display}</button>)}
                    </For>
                </div>
            </div>
        </Show>
    )
}
