<script lang="ts">
  import type Deal from "$lib/models/entity/deal";
  import type Stage from "$lib/models/entity/stage";
  import DealCard from "./dealCard.svelte";

  export let width: string = "20%";

  export let stageParams: Stage;

  const { title, color, deals } = stageParams;
</script>

<div id="stage-container" style:width>
  <div id="stage-title" style:background-color={color}>{title}</div>
  <div id="stage-deals">
    {#each deals as deal}
      <DealCard dealParams={deal}></DealCard>
    {/each}
  </div>
</div>

<style>
  #stage-container {
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  #stage-title {
    padding-inline: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  #stage-deals {
    flex: 1;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
</style>
