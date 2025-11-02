<script lang="ts">
  import type Deal from "$lib/models/entity/deal";

  export let width: string = "20%";

  export let dealParams: Deal;
</script>

<div id="deal" style:width>
  <div id="first">
    <img src={dealParams.costumer.photo} alt="" />
    <h1>{dealParams.costumer.name}</h1>
  </div>
  <i>{dealParams.chat.messages[dealParams.chat.messages.length - 1]}</i>
  <div id="bottom">
    <h2>Negócio criado em {dealParams.createdAt}</h2>
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
