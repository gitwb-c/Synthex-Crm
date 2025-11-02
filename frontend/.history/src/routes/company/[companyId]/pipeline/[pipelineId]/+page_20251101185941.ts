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
            id: "stage:1",
            color: "red",
            title: 'Entrada',
            deals: [{
                id: '',
                source: '',
                costumer: {
                    id: '',
                    name: '',
                    photo: ''
                },
                chat: {
                    id: "chat:1",
                    messages: [
                        {
                            id: '',
                            author: {id: "costumer:1", name: "Bebe"},
                            type: 'image',
                            message: '',
                            createdAt: new Date("2025-11-27").toLocaleDateString("pt-BR")
                        }
                    ]
                }
                createdAt: new Date("2025-11-27").toLocaleDateString("pt-BR"),
            }]
        }
    ]

    return {
        exists: true,
        pipelines,
    };
}) satisfies PageLoad;