<script lang="ts">
  import Helper from "$lib/helper";
  import type Deal from "$lib/models/entity/deal";
  import type Stage from "$lib/models/entity/stage";
  import DealCard from "./dealCard.svelte";

  export let width: string = "20%";

  export let stageParams: Stage;

  const { title, color, deals } = stageParams;
</script>

<div id="stage-container" style:width>
  <div
    id="stage-title"
    style:background-color={color}
    style:color={"Helper.getContrastColor(color)"}
  >
    {title}
  </div>
  <div id="stage-deals">
    {#each deals as deal}
      <DealCard themeColor={`solid ${color} 8px`} dealParams={deal}></DealCard>
    {/each}
  </div>
</div>

<style>
  #stage-container {
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  #stage-title {
    padding-inline: 0.7rem;
    padding-block: 0.5rem;
    box-shadow: 0.2rem 0.2rem 0.2rem rgba(0, 0, 0, 0.3);
    border-top-right-radius: 1.2rem;
    border-bottom-right-radius: 1.2rem;
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
