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
    <h2>{costumer.name}</h2>
  </div>
  <i>{chat!.messages[chat!.messages.length - 1].message}</i>
  <div id="bottom">
    <h4>Negócio criado em {createdAt}</h4>
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
    padding-block: 3rem;
    min-height: 4rem;
    max-height: 20rem;
  }

  #first {
    width: 100%;
    height: 10%;
    display: flex;
    align-items: center;
    gap: 1rem;

    & > span {
      width: 3rem;
      height: 3rem;
      border-radius: 50%;
      box-shadow: 0 0 0.4rem rgba(0, 0, 0, 0.4);
      overflow: hidden;

      & > img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }
</style>
