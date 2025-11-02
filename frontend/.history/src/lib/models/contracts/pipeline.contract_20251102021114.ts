import type DealCemFields from "../entity/dealCrmFields";
import type Field from "../entity/field";
import type Stage from "../entity/stage";

export default interface PipelineContract {
  pipelines: {
    edges: {
      node: {
        id: string;
        name: string;
        stages: Pick<Stage, "id" | "color" | "deals"> & {
          dealCrmFields: Pick<DealCemFields, "id">;
        };
      }[];
    };
  };
}
