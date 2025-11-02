<script lang="ts">
  import type Deal from "$lib/models/entity/deal";
  import type Stage from "$lib/models/entity/stage";

  export let width: string = "20%";

  export let stageParams: Stage;

  const { title, color, deals } = stageParams;
</script>

<div id="stage-container" style:width>
  <div id="stage-title"></div>
</div>

<style>
  #stage-container {
    height: 100%;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
  }
</style>
