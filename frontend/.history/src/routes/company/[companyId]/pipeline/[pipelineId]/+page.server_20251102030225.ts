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
      availablePipelines: pipelines {
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

  if (!res.data.currentPipeline.edges.length) {
    return {
      exists: false,
    };
  }

  return {
    currentPipeline: res.data.currentPipeline.edges[0].node,
    availablePipelines: res.data.availablePipelines.edges,
  };
}) satisfies PageServerLoad;
