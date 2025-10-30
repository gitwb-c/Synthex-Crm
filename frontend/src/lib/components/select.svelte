<script lang="ts">
  export let options: { title: string; value: string }[];
  export let fontSize: string | undefined = undefined;
  export let width: string | undefined = undefined;
  export let value: string;
</script>

<select bind:value on:change style:font-size={fontSize} style:width>
  {#each options as opt}
    <option value={opt.value}>{opt.title}</option>
  {/each}
</select>

<style>
  select {
    padding-block: 0.5rem;
    padding-inline: 0.6rem;
    border: none;
    outline: none;
    box-shadow: 0.005rem 0.2rem 0.2rem rgba(0, 0, 0, 0.2);
    border-radius: 0.5rem;
  }
</style>
