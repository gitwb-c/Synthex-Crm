import type Field from "./field";

export default interface Deal {
    id: string,
    title: string,
    source: string,
    fields: Field[],
    createdAt: Date,
    updatedAt: Date,
}