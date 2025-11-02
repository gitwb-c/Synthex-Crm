import { BACKEND_API_URL } from "$env/static/private";
import Helper from "$lib/helper";
import type PipelineContract from "$lib/models/contracts/pipeline.contract";
import UnauthorizedError from "$lib/models/errors/Unauthorized.error";
import { redirect } from "@sveltejs/kit";
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

  let res;

  try {
    res = await Helper.fetchUrl<PipelineContract>(
      `${BACKEND_API_URL}/graphql/query/${companyId}`,
      {
        method: "POST",
        body: JSON.stringify({ query }),
      }
    );
  } catch (error) {
    if (error instanceof UnauthorizedError) {
      throw redirect(300, `/company/${params.companyId}/login`);
    }
  }

  if (res!.data.currentPipeline.edges.length == 0) {
    return {
      exists: false,
      availablePipelines: res!.data.availablePipelines.edges,
      currentPipeline: res!.data.currentPipeline.edges[0].node,
    };
  }

  return {
    currentPipeline: res!.data.currentPipeline.edges[0].node,
    availablePipelines: res!.data.availablePipelines.edges,
  };
}) satisfies PageServerLoad;
