import { Icon } from '@iconify-icon/solid';
import { Button } from "./basic/Button"

export default function SideBar() {
    return (<>
        <div class="flex flex-col w-full items-center space-y-4 py-4">
            <Button icon='ep:menu' onClick={()=>alert("hello icon")}></Button>
            <Button icon="fluent:history-24-filled"></Button>
            <Button icon="solar:bookmark-linear"></Button>
        </div >
    </>)
}