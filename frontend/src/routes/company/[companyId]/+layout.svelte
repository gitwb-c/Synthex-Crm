<script lang="ts">
  import Header from "$lib/components/header.svelte";
  import type { LayoutProps } from "../../[companyId]/$types";
  import { page } from "$app/state";

  const path = page.url.pathname;

  let { data, children }: LayoutProps = $props();
</script>

{#if path.includes("/login")}
  {@render children()}
{:else}
  <div id="wrapper">
    <Header pathname={path} width="12%"></Header>
    <main>
      {#if !data.exists}
        <h1>Empresa não encontrada!</h1>
      {:else}
        {@render children()}
      {/if}
    </main>
  </div>
{/if}

<style>
  #wrapper {
    width: 100vw;
    height: auto;
    display: flex;

    & > main {
      width: 88%;
      padding: 1.3rem;
      overflow-x: hidden;
      overflow-y: auto;
      background-color: var(--azul-palido);
    }
  }
</style>
