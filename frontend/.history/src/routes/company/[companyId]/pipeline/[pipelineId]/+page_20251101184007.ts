import type { PageLoad } from './$types';

export const load = (async ({params}) => {
    const { pipelineId} = params;

    const delay = (seconds: number)=>new Promise(resolve => setTimeout(resolve, seconds*1000))

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