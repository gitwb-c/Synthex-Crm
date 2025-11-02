import type { LayoutServerLoad } from "./$types";

export const load = (async ({ params }) => {
  const { companyId } = params;

  if (!["teste1", "teste2"].includes(companyId)) {
    return {
      exists: false,
    };
  }

  return {
    exists: true,
  };
}) satisfies LayoutServerLoad;
