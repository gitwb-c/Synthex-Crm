import type Deal from "./deal";
import type Message from "./message";
import type Queue from "./queue";

export default interface Chat {
    id: string,
    source?: string,
    phone?: string, 
    messages: Message[],
    queue?: Queue,
    deal?: Deal
}