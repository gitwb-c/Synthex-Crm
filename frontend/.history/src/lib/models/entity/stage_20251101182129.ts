import type Pipeline from "$lib/components/kanban/pipeline.svelte";

export default interface Stage {
    id: string,
    title: string,
    color: string,
    pipeline: Pipeline
}