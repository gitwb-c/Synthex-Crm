<script lang="ts">
  import type Deal from "$lib/models/entity/deal";
  import type Stage from "$lib/models/entity/stage";

  export let width: string = "20%";

  export let stageParams: Stage;

  const { title, color, deals } = stageParams;
</script>

<div id="stage-container" style:width>
  <div id="stage-title" style:background-color={color}>{title}</div>
  <div id="stage-deals"></div>
</div>

<style>
  #stage-container {
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  #stage-title {
    padding: 1rem;
    display: grid;
    place-items: center;
  }

  #stage-deals {
    flex: 1;
    overflow-y: auto;
  }
</style>
