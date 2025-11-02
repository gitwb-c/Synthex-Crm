<script lang="ts">
  import { goto } from "$app/navigation";
  import Button from "$lib/components/button.svelte";
  import Input from "$lib/components/input.svelte";
  import Stage from "$lib/components/kanban/stage.svelte";
  import Select from "$lib/components/select.svelte";
  import type { PageProps } from "./$types";
  import { Search } from "lucide-svelte";

  let { data, params }: PageProps = $props();

  let choosenPipeline = params.pipelineId;

  const goToAnotherPipeline = async (pipelineId: string) => {
    try {
      await goto(`/company/${params.companyId}/pipeline/${pipelineId}`);
    } catch (e) {
      alert(e);
    }
  };

  const stageToIterate = data.currentPipeline.stages as unknown as Pick<
    Stage,
    "name" | "color"
  >;

  stageToIterate.
</script>

{#if !data.exists}
  <h1>Etapa {params.pipelineId} não existe!</h1>
{:else}
  <div id="first-wrapper">
    <Input width={"60%"} type="text" leftIcon={Search} />

    <Button
      textColor="white"
      backgroundColor={"var(--azul-claro)"}
      title="Gerenciar automações"
    ></Button>
  </div>

  <div id="second-wrapper">
    <Select
      value={choosenPipeline}
      on:change={(e) =>
        goToAnotherPipeline((e.target as HTMLSelectElement).value)}
      width={"20%"}
      fontSize={"1rem"}
      options={data.availablePipelines.map((pipe) => ({
        title: pipe.node.name,
        value: pipe.node.id,
      }))}
    ></Select>

    <Button
      textColor="white"
      backgroundColor="var(--azul-marinho)"
      title="Configurar etapa"
    ></Button>
  </div>

  <section id="kanban">
    {#each data.currentPipeline.stages as stage}
      <Stage width="19rem" stageParams={stage}></Stage>
    {/each}

    <Stage
      width="19rem"
      stageParams={{
        color: "rgb(0,0,0)",
        id: "newStage",
        name: "Criar novo estágio",
      }}
    ></Stage>
  </section>
{/if}

<style>
  #first-wrapper {
    width: 100%;
    height: fit-content;
    margin-bottom: 1rem;
    display: flex;
    justify-content: space-between;
  }

  #second-wrapper {
    width: 100%;
    display: flex;
    gap: 2rem;
  }

  #kanban {
    margin-top: 1.6rem;
    flex: 1;
    overflow-x: auto;
    overflow-y: hidden;
    display: flex;
    gap: 2rem;
  }
</style>
