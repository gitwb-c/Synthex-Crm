import Helper from "$lib/helper";
import type { LayoutServerLoad } from "./$types";

export const load = (async ({ params }) => {
  const { companyId } = params;

  const res = Helper.fetchUrl("");

  return {
    exists: true,
  };
}) satisfies LayoutServerLoad;
