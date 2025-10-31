<script lang="ts">
  import type { ComponentType, SvelteComponent } from "svelte";
  import type { HTMLInputTypeAttribute } from "svelte/elements";
  import { Eye, EyeClosed } from "lucide-svelte";

  export let type: HTMLInputTypeAttribute;
  export let backgroundColor: string = "white";
  export let textColor: string = "black";
  export let width: string = "2rem";

  export let leftIcon: ComponentType | null = null;
  export let onLeftIconPressed: () => void = () => {};
  export let border: string = "";

  export let isPassword: boolean = false;
  export let iconWidth: string = "3%";
  export let inputWidth: string = "97%";

  export let value: string = "";

  $: searchBackground = "transparent";

  $: visible = false;
</script>

<div
  id="wrapper"
  style:border
  style:background-color={backgroundColor}
  style:width
>
  {#if isPassword}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <span
      style:background={searchBackground}
      style:height="100%"
      style:width={iconWidth}
      on:mouseenter={() => (searchBackground = "rgba(0,0,0,.05)")}
      on:mouseleave={() => (searchBackground = "transparent")}
      on:click={() => (visible = !visible)}
    >
      <svelte:component this={visible ? EyeClosed : Eye}></svelte:component>
    </span>
    <input
      style:width={inputWidth}
      type={visible ? "text" : "password"}
      style:color={textColor}
      bind:value
    />
  {:else if leftIcon}
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <span
      style:background={searchBackground}
      style:height="100%"
      style:width={iconWidth}
      on:mouseenter={() => (searchBackground = "rgba(0,0,0,.05)")}
      on:mouseleave={() => (searchBackground = "transparent")}
      on:click={onLeftIconPressed}
    >
      <svelte:component this={leftIcon}></svelte:component>
    </span>
    <input style:width={inputWidth} {type} style:color={textColor} bind:value />
  {:else}
    <input {type} style:color={textColor} bind:value />
  {/if}
</div>

<style>
  #wrapper {
    display: flex;
    border-radius: 0.5rem;
    box-shadow: 0.005rem 0.2rem 0.2rem rgba(0, 0, 0, 0.2);
    align-items: center;
    padding-inline: 0.6rem;

    & > span {
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;
      border-radius: 0.4rem;

      &:hover {
        transform: scale(1.01);
      }
    }
  }
  input {
    padding-block: 0.5rem;
    padding-inline: 0.6rem;
    border: none;
    outline: none;
  }
</style>
