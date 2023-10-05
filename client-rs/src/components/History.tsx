import { For } from "solid-js"
import { state } from "../store/store"
import { ILog } from "../interfaces/ILog"

export default function History() {
    return (
        <>
            <div class="flex flex-col w-full h-full">
                <For each={state.logs}>{
                    (log) => (
                        <div class=" bg-slate-500 my-2 mx-10 rounded-lg p-5">
                            <div class=" text-lime-80 text-sky-600 bg-red-300 w-fit rounded-3xl px-5 py-1 text">
                                {log.action}
                            </div>
                            <div>
                                {log.file}
                            </div>
                            <div>
                                {log.time}
                            </div>
                        </div>
                    )
                }
                </For>
            </div>
        </>
    )
}