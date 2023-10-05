import { Icon } from '@iconify-icon/solid';
import { Button } from "./basic/Button"
import { setState,state } from '../store/store';
import { produce } from 'solid-js/store';

export default function SideBar() {
    return (<>
        <div class="flex flex-col w-full items-center space-y-4 py-4">
            <Button icon='ep:menu' onClick={()=>goto("files")}></Button>
            <Button icon="fluent:history-24-filled" onClick={()=>goto("history")}></Button>
            <Button icon="solar:bookmark-linear" onClick={()=>goto("dialog")}></Button>
        </div >
    </>)
}

function goto(page:string) {
    setState("page", page)
}