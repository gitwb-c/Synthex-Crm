import type { PageServerLoad } from "./$types";

export const load = (async ({ params }) => {
  const { companyId } = params;
  return {};
}) satisfies PageServerLoad;
