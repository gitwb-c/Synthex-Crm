<script lang="ts">
  import favicon from "$lib/assets/Component 1 2.svg";

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
  <title>Blendine CRM</title>
</svelte:head>

{@render children?.()}
