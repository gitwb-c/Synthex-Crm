import type Deal from "./deal";
import type Queue from "./queue";

export default interface Chat {
    id: string,
    source: string,
    phone: string, 
    queue: Queue,
    deal: Deal
}