import type Stage from '$lib/models/entity/stage';
import type { PageLoad } from './$types';

export const load = (async ({params}) => {
    const delay = (seconds: number)=> new Promise(resolve => setTimeout(resolve, seconds * 1000));
    
    await delay(2);

    const { pipelineId} = params;


    const pipelines =["Qualificação", "Oportunidade", "Pós-Vendas"];

    if (!pipelines.includes(pipelineId)) {
        return {
            exists: false
        }
    }

    const stages: Stage[] = [
        {
            id: "asas",
            color: "red",
            pipeline: {
                id: '',
                title: '',
            },
            title: '',
            queue: {
                id: '',
                name: '',
                timeToSkip: 0,
                department: {
                    id: '',
                    name: '',
                    employees: []
                },
                employees: []
            },
            requiredFields: [],
            deals: [{
                id: '',
                source: '',
                costumer: {
                    id: '',
                    name: '',
                    photo: ''
                },
                fields: [],
                chat: undefined,
                createdAt: undefined,
                updatedAt: undefined
            }]
        }
    ]

    return {
        exists: true,
        pipelines,
    };
}) satisfies PageLoad;