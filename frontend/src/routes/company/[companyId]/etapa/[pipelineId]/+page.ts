import type { PageLoad } from './$types';

export const load = (async ({params}) => {
    const { pipelineId} = params;

    const pipelines =["Qualificação", "Oportunidade", "Pós-Vendas"];

    if (!pipelines.includes(pipelineId)) {
        return {
            exists: false
        }
    }
    return {
        exists: true,
        pipelines,
    };
}) satisfies PageLoad;