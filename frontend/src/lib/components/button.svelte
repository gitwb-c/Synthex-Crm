<script lang="ts">
  import type { ComponentType, SvelteComponent } from "svelte";

  export let title: string;
  export let type: "button" | "submit" | "reset" | undefined = undefined;
  export let backgroundColor: string;
  export let textColor: string = "black";
  export let fontSize: string = "";
  export let fontWeight: string = "";
</script>

<button
  {title}
  {type}
  style:background-color={backgroundColor}
  style:color={textColor}
  style:border-color={backgroundColor}
  ><span style:font-size={fontSize} style:font-weight={fontWeight}>
    {title}
  </span>
</button>

<style>
  button {
    padding-inline: 0.8rem;
    padding-block: 0.5rem;
    border-radius: 0.5rem;
    box-shadow: 0.005rem 0.2rem 0.2rem rgba(0, 0, 0, 0.2);
    cursor: pointer;

    &:hover {
      filter: brightness(1.1);
    }
  }
</style>
