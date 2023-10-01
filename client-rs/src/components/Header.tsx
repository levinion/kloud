import { Button } from "./basic/Button"

export default function Header() {
    return (
        <div class="flex flex-row h-full w-full items-center space-x-5">
            <input type="text" class="h-3/4 bg-zinc-800 ml-5 px-2" />
            <Button icon="mdi:new-box"></Button>
        </div>
    )
}