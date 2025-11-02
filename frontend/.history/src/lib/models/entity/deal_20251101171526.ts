import type Costumer from "./costumer";
import type Field from "./field";

export default interface Deal {
    id: string,
    title: string,
    source: string,
    costumer: Costumer,
    fields: Field[],
    createdAt: Date,
    updatedAt: Date,
}
