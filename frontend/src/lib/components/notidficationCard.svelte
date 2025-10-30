<script lang="ts">
  import type { Notification } from "$lib/models/types/notification";
  import { notifications } from "$lib/stores/notifications";
  import { X } from "lucide-svelte";
  import { onMount } from "svelte";

  let {
    id,
    message,
    backgroundColor,
    fading,
    textColor,
    onClick,
    timeToSkip,
  }: Notification = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  style:background={backgroundColor}
  style:color={textColor}
  style:translate={fading ? "100%" : ""}
>
  <h4>
    {message}
  </h4>
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore event_directive_deprecated -->
  <span on:click={() => notifications.deleteNotification(id)}>
    <X cursor="pointer" color={textColor}></X>
  </span>
</div>

<style>
  div {
    width: 100%;
    border-top-left-radius: 1rem;
    border-bottom-left-radius: 1rem;
    min-height: 3rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding-inline: 1rem;
  }
</style>
