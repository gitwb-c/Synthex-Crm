import type Chat from "./chat";
import type Costumer from "./costumer";
import type Field from "./crmField";
import type Employee from "./employee";

export default interface Deal {
  id: string;
  title?: string;
  source?: string;
  costumer: Omit<Costumer, "deal">;
  employee: Employee;
  fields?: Field[];
  chat?: Chat;
  createdAt?: string;
  updatedAt?: string;
}
