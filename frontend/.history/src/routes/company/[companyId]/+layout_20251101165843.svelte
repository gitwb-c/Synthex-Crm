<script lang="ts">
  import Header from "$lib/components/header.svelte";
  import type { LayoutProps } from "../../[companyId]/$types";
  import { page } from "$app/stores";
  import { notifications } from "$lib/stores/notifications";
  import NotidficationCard from "$lib/components/notidficationCard.svelte";

  const path = $derived($page.url.pathname);

  let { data, children }: LayoutProps = $props();

  const notifications_ = notifications;
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

<div id="notifications-tab">
  {#each $notifications_.toReversed() as not}
    <NotidficationCard
      id={not.id}
      backgroundColor={not.backgroundColor}
      fading={not.fading}
      message={not.message}
      onClick={not.onClick}
      textColor={not.textColor}
      timeToSkip={not.timeToSkip}
    ></NotidficationCard>
  {/each}
</div>

<style>
  #wrapper {
    width: 100vw;
    height: auto;
    display: flex;

    & > main {
      width: 88%;
      border: solid 1px;
      /* height: 100%; */
      padding: 1.3rem;
      overflow-x: hidden;
      overflow-y: auto;
      background-color: var(--azul-palido);
    }
  }

  #notifications-tab {
    position: fixed;
    right: 3rem;
    top: 3rem;
    width: 16rem;
    display: flex;
    flex-direction: column;
    height: 60vh;
    overflow-y: auto;
    gap: 1rem;
    overflow-x: hidden;
  }
</style>
