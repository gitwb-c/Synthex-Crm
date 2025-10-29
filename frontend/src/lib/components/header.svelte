<script lang="ts">
  import logo from "$lib/assets/Component 1 2.svg";
  import { Kanban, MessageCircle, Building } from "lucide-svelte";
  export let width: string;
  export let pathname: string;

  $: reactiveWidth = width;

  $: pages = {
    kanban: pathname.includes("etapa"),
    chat: pathname.includes("chat"),
    departments: pathname.includes("department"),
  };

  $: notHovered = reactiveWidth == width;
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<header
  style={`width: ${reactiveWidth}`}
  onmouseenter={() =>
    (reactiveWidth = (parseInt(width.replace("%", "")) * 2.6).toString() + "%")}
  onmouseleave={() => (reactiveWidth = width)}
>
  <img src={logo} alt="logo" width="45px" />
  <ul>
    {#if notHovered}
      <li class:active={pages.kanban}><Kanban width="100%"></Kanban></li>
      <li class:active={pages.chat}>
        <MessageCircle width="100%"></MessageCircle>
      </li>
      <li class:active={pages.departments}>
        <Building width="100%"></Building>
      </li>
    {:else}
      <li class:active={pages.kanban} class="hovered">
        <Kanban width="100%"></Kanban> Kanban
      </li>
      <li class:active={pages.chat} class="hovered">
        <MessageCircle width="100%"></MessageCircle> Chats
      </li>
      <li class:active={pages.departments} class="hovered">
        <Building width="100%"></Building> Departamentos
      </li>
    {/if}
  </ul>
</header>

<style>
  header {
    height: 100vh;
    background-color: var(--azul-primario);
    display: flex;
    flex-direction: column;
    padding-inline: 1rem;
    padding-block: 1rem;
    align-items: center;
    gap: 2rem;
  }

  ul {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  li {
    width: 100%;
    display: flex;
    justify-content: center;
    gap: 1rem;
    cursor: pointer;
  }

  li:hover {
    color: white;
    transform: scale(0.98);
  }

  li.hovered {
    justify-content: flex-start;
  }

  li.active {
    &:first-child {
      color: white;
    }
  }
</style>
