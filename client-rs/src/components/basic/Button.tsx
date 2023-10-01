import { Icon } from "@iconify-icon/solid"
import { JSX } from "solid-js/jsx-runtime"

interface ButtonProps extends JSX.BaseHTMLAttributes<HTMLSpanElement> {
    icon: string
    height?: string
}

export function Button(props: ButtonProps) {
    return (
        <div class="hover:bg-blue-300 rounded-lg p-1">
            <Icon
                icon={props.icon}
                height={props.height || 24}
                onClick={props.onClick}
            >
            </Icon>
        </div>
    )
}