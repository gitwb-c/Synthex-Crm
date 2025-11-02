import type Pipeline from "$lib/components/kanban/stage.svelte";
import type Deal from "./deal";
import type Field from "./field";
import type Queue from "./queue";

export default interface Stage {
    id: string,
    title: string,
    color: string,
    pipeline: Pipeline,
    queue: Queue,
    requiredFields: Field[],
    deals: Omit<Deal, "stage" | "pipeline">[]
}