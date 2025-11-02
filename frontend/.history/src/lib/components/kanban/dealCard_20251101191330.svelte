<script lang="ts">
  import type Deal from "$lib/models/entity/deal";

  export let dealParams: Deal;

  const { costumer, createdAt, chat } = dealParams;
</script>

<div id="deal">
  <div id="first">
    <img src={costumer.photo} alt="" />
    <h1>{costumer.name}</h1>
  </div>
  <i>{dealParams.chat!.messages[chat!.messages.length - 1]}</i>
  <div id="bottom">
    <h2>Negócio criado em {createdAt}</h2>
  </div>
</div>

<style>
  #deal {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background-color: white;
    box-shadow: 0 0 1rem rgba(0, 0, 0, 0.4);
  }

  #first {
    width: 100%;
    height: 10%;
    display: flex;
    align-items: center;
  }
</style>
