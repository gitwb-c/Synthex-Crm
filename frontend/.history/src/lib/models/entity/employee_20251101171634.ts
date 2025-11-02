import type Deal from "./deal";
import type Person from "./person";

export default interface Employee extends Person {
    deals: Deal[],
    email: string,
}