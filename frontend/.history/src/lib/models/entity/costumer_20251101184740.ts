import type Deal from "./deal";
import type Person from "./person";

export default interface Costumer extends Person {
    deal: Omit<Deal, "costumer">
}