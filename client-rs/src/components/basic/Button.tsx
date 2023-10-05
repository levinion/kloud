import { Icon } from "@iconify-icon/solid"
import { Switch, Match } from "solid-js"
import { JSX } from "solid-js/jsx-runtime"

export interface ButtonProps extends JSX.BaseHTMLAttributes<HTMLSpanElement> {
    icon?: string
    height?: string
    text?: string
}

export function Button(props: ButtonProps) {
    return (
        <div class="hover:bg-blue-300 rounded-lg">
            <Switch fallback={<button onClick={props.onClick}>{props.text || "button"}</button>}>
                <Match when={props.icon}>
                    <Icon
                        icon={props.icon!}
                        height={props.height || 24}
                        onClick={props.onClick}
                        onContextMenu={props.onContextMenu}
                    >
                    </Icon>
                </Match>
            </Switch>
        </div>
    )
}