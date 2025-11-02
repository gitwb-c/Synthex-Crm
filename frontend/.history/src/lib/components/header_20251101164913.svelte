<script lang="ts">
  import logo from "$lib/assets/Component 1 2.svg";
  import { Kanban, MessageCircle, Building } from "lucide-svelte";
  export let width: string;
  export let pathname: string;

  $: pages = {
    kanban: pathname.includes("etapa"),
    chat: pathname.includes("chat"),
    departments: pathname.includes("department"),
  };
</script>

<header style={`width: ${width}`}>
  <h1>Synthex</h1>
  <ul>
    <li class:active={pages.kanban} class="hovered">
      <Kanban size="1.3rem"></Kanban>
      Kanban
    </li>
    <li class:active={pages.chat} class="hovered">
      <MessageCircle size="1.3rem"></MessageCircle> Chats
    </li>
    <li class:active={pages.departments} class="hovered">
      <Building size="1.3rem"></Building> Departamentos
    </li>
  </ul>
</header>

<style>
  header {
    height: 100vh;
    background-color: var(--azul-primario);
    display: flex;
    flex-direction: column;
    padding-inline: 0.7rem;
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
    justify-content: flex-start;
    align-items: center;
    gap: 0.7rem;
    cursor: pointer;
    font-size: 0.75rem;
  }

  li:hover {
    color: white;
    transform: scale(0.98);
  }

  li.active {
    &:first-child {
      color: white;
    }
  }
</style>
