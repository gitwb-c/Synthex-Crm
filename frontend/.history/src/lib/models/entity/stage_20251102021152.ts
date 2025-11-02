import type Deal from "./deal";
import type Field from "./crmField";
import type Pipeline from "./pipeline";
import type Queue from "./queue";

export default interface Stage {
  id: string;
  title: string;
  color: string;
  pipeline?: Omit<Pipeline, "stages">;
  queue?: Queue;
  requiredFields?: Field[];
  deals?: Omit<Deal, "stage" | "pipeline">[];
}
