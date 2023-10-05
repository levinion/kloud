import { JSX } from "solid-js/jsx-runtime"

interface CardProps extends JSX.BaseHTMLAttributes<HTMLDivElement> {}

export default function Card(props: CardProps) {
    return (
        <div class="w-4/5 bg-white shadow-xl h-96 rounded-lg mx-auto">
            {props.children}
        </div>
    )
}