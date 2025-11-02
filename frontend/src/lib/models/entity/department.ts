import type Employee from "./employee";
import type Queue from "./queue";

export default interface Department {
    id: string,
    name: string,
    employees: Employee[],
    queues: Omit<Queue, "department">[]
}