import type DealCrmFields from "../entity/dealCrmFields";
import type Field from "../entity/crmField";
import type Stage from "../entity/stage";
import type Deal from "../entity/deal";
import type Costumer from "../entity/costumer";
import type Employee from "../entity/employee";

export default interface PipelineContract {
  query1: {
    pipelines: {
      edges: {
        node: {
          id: string;
          name: string;
          stages: Pick<Stage, "id" | "color"> &
            {
              deals: Pick<Deal, "title" | "createdAt"> & {
                costumer: Pick<Costumer, "id" | "name">;
              } & { employee: Pick<Employee, "id" | "name" | "photo"> }[];
            }[];
        };
      }[];
    };
  };
}
