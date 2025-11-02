import { BACKEND_API_URL } from "$env/static/private";
import type { PageServerLoad } from "./$types";

export const load = (async ({ params }) => {
  const { companyId, pipelineId } = params;

  const query = `
    query {
      pipelines (where: {id: ${pipelineId}}) {
        edges {
          node {
            id
            title
            stages {
              id
              title
              color
              deals {
                id
                costumer {
              id
              name
              photo
            }
              }
            }
          }
        }
      }
    }
  `;

  const res = await fetch(`${BACKEND_API_URL}/graphql/query/${companyId}`, {
    method: "POST",
    body,
  });
  return {};
}) satisfies PageServerLoad;
