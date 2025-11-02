import Helper from "$lib/helper";
import type { LayoutServerLoad } from "./$types";
import { BACKEND_API_URL } from "$env/static/private";

export const load = (async ({ params }) => {
  const { companyId } = params;

  return {
    exists: true,
  };
}) satisfies LayoutServerLoad;
