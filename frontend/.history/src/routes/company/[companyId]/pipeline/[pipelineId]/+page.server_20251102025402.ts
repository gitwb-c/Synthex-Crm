import { BACKEND_API_URL } from "$env/static/private";
import Helper from "$lib/helper";
import type PipelineContract from "$lib/models/contracts/pipeline.contract";
import type { PageServerLoad } from "./$types";

export const load = (async ({ params }) => {
  const { companyId, pipelineId } = params;

  const query = `
    query {
      currentPipelines: pipelines (where: {id: ${pipelineId}}) {
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
                employee {
                  id
                  name
                  photo
                }
                chat {
                  messages (last: 1, orderBy: ["createdAt"]){
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
                }
              }
            }
          }
        }
      }
      avaiblePipelines: pipelines {
        edges {
          node {
            id
            name
          }
        }
      }
    }
  `;

  const res = await Helper.fetchUrl<PipelineContract>(
    `${BACKEND_API_URL}/graphql/query/${companyId}`,
    {
      method: "POST",
      body: JSON.stringify({ query }),
    }
  );

  const pipeline = res.data.pipelines.edges[0].node;
  return {};
}) satisfies PageServerLoad;
