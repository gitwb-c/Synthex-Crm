import type Costumer from "./costumer";
import type Employee from "./employee";

export default interface Message {
    id: string,
    author: Employee | Costumer,
    type: "image" | "audio" | "text",
    message: string,
    timestamp: Date
}