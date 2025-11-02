import type Chat from "./chat";
import type Costumer from "./costumer";
import type Field from "./field";

export default interface Deal {
    id: string,
    title: string,
    source: string,
    costumer: Omit<Costumer, "deal">,
    fields: Field[],
    chat: Chat,
    createdAt: Date,
    updatedAt: Date,
}
