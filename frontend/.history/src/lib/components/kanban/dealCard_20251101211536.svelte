<script lang="ts">
  import type Deal from "$lib/models/entity/deal";

  export let dealParams: Deal;

  const { costumer, createdAt, chat } = dealParams;

  export let themeColor: string;

  let lastMessage = chat!.messages[chat!.messages.length - 1].message;

  if (lastMessage.length > 43) {
    lastMessage = lastMessage.slice(0, 43) + "...";
  } else if (!lastMessage) {
    lastMessage = "...";
  }
</script>

<div id="deal" style:border-left={themeColor}>
  <div id="first">
    <span>
      <img src={costumer.photo} alt="foto" />
    </span>
    <h2>{costumer.name}</h2>
  </div>
  <i>"{lastMessage}"</i>
  <div id="bottom">
    <h5>Negócio criado em {createdAt}</h5>
  </div>
</div>

<style>
  #deal {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background-color: white;
    box-shadow: 0 0 1rem rgba(0, 0, 0, 0.4);
    padding-inline: 1.5rem;
    padding-block: 1.4rem;
    height: 13rem;
    border-top-right-radius: 2rem;
    border-bottom-right-radius: 2rem;
  }

  i {
    color: var(--grey);
  }

  #first {
    width: 100%;
    height: 2.5rem;
    display: flex;
    align-items: center;
    gap: 1rem;

    & > h2 {
      font-weight: 550;
      font-size: 1.4rem;
    }

    & > span {
      width: 2.5rem;
      height: 100%;
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

  #bottom > h5 {
    font-weight: 500;
    color: rgba(0, 0, 0, 0.9);
  }
</style>
