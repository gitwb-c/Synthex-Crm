<script lang="ts">
  import type Deal from "$lib/models/entity/deal";

  export let dealParams: Deal;

  const { costumer, createdAt, chat } = dealParams;
</script>

<div id="deal">
  <div id="first">
    <span>
      <img src={costumer.photo} alt="foto" />
    </span>
    <h1>{costumer.name}</h1>
  </div>
  <i>{chat!.messages[chat!.messages.length - 1].message}</i>
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
    padding-inline: 1rem;
    padding-block: 2rem;
  }

  #first {
    width: 100%;
    height: 10%;
    display: flex;
    align-items: center;

    & > span {
      width: 1px;
    }

    & > img {
      width: 100%;
    }
  }
</style>
