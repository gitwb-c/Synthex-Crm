<script lang="ts">
  import Header from "$lib/components/header.svelte";
  import type { LayoutProps } from "../../[companyId]/$types";
  import { page } from "$app/state";

  const path = page.url.pathname;

  let { data, children }: LayoutProps = $props();
</script>

<div id="wrapper">
  <Header pathname={path} width="6%"></Header>
  <main>
    {#if !data.exists}
      <h1>Empresa não encontrada!</h1>
    {:else}
      {@render children()}
    {/if}
  </main>
</div>

<style>
  #wrapper {
    width: 100vw;
    height: auto;
    display: flex;

    & > main {
      width: 50%;
      padding: 1.3rem;
    }
  }
</style>
