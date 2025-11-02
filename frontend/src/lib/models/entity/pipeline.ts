import type Stage from "./stage";

export default interface Pipeline {
    id: string,
    title: string,
    stages: Stage[]
}