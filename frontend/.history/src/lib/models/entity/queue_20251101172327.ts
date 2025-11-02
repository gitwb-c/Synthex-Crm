import type Department from "./department";
import type Employee from "./employee";

export default interface Queue {
    id: string,
    name: string,
    timeToSkip: number,
    department: Department,
    employees: Employee[]
}