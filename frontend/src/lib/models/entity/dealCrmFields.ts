import type Deal from "./deal";
import type Field from "./crmField";

export default interface DealCrmFields {
  value: string | number;
  field: Field;
  deal: Deal;
}
