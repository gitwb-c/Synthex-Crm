<script lang="ts">
  import type Deal from "$lib/models/entity/deal";

  export let width: string = "20%";

  export let dealParams: Deal;
</script>

<div id="stage-container" style:width></div>

<style>
  #stage-container {
    height: 100%;
    overflow-y: auto;
  }
</style>
