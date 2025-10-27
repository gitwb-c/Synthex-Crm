<script lang="ts">
    import type { PageProps } from './$types';

    let { data, params }: PageProps = $props();
    const {pipelineId} = params
</script>

<h1>Etapa {pipelineId}</h1>
