<script lang="ts">
  import { goto } from "$app/navigation";
  import Button from "$lib/components/button.svelte";
  import Input from "$lib/components/input.svelte";
  import Select from "$lib/components/select.svelte";
  import type { PageProps } from "./$types";
  import { Search } from "lucide-svelte";

  let { data, params }: PageProps = $props();
  const { pipelineId } = params;

  let choosenPipeline = params.pipelineId;

  const goToAnotherPipeline = async (pipelineId: string) => {
    try {
      await goto(`/company/${params.companyId}/pipeline/${pipelineId}`);
    } catch (e) {
      alert(e);
    }
  };
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
      options={data.pipelines!.map((pipe) => ({ title: pipe, value: pipe }))}
    ></Select>

    <Button
      textColor="white"
      backgroundColor="var(--azul-marinho)"
      title="Configurar etapa"
    ></Button>
  </div>
{/if}

<style>
  #first-wrapper {
    width: 100%;
    height: 5%;
    margin-bottom: 1rem;
    display: flex;
    justify-content: space-between;
  }

  #second-wrapper {
    width: 100%;
    display: flex;
    gap: 2rem;
  }
</style>
