import type DealCrmFields from "../entity/dealCrmFields";
import type Field from "../entity/crmField";
import type Stage from "../entity/stage";
import type Deal from "../entity/deal";

export default interface PipelineContract {
  pipelines: {
    edges: {
      node: {
        id: string;
        name: string;
        stages: (Pick<Stage, "id" | "color" > & deals: Pick<"deals", "title">);
      }[];
    };
  };
}
