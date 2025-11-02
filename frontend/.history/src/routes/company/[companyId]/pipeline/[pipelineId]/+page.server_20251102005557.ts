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
            name
            stages {
              id
              name
              color
              deals {
                id
                title
                source
                createdAt
                chat {
                  messages (limit: 1, orderBy: 'DESC'){
                    type
                    text {
                      text
                    }
                    file {
                      url
                    }
                  }
                }
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
